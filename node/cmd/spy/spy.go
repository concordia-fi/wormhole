package spy

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/certusone/wormhole/node/pkg/common"
	"github.com/certusone/wormhole/node/pkg/p2p"
	gossipv1 "github.com/certusone/wormhole/node/pkg/proto/gossip/v1"
	spyv1 "github.com/certusone/wormhole/node/pkg/proto/spy/v1"
	"github.com/certusone/wormhole/node/pkg/supervisor"
	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	ipfslog "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/wormhole-foundation/wormhole/sdk/vaa"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	rootCtx       context.Context
	rootCtxCancel context.CancelFunc
)

var (
	p2pNetworkID *string
	p2pPort      *uint
	p2pBootstrap *string

	statusAddr *string

	nodeKeyPath *string

	logLevel *string

	spyRPC *string
)

func init() {
	p2pNetworkID = SpyCmd.Flags().String("network", "/wormhole/dev", "P2P network identifier")
	p2pPort = SpyCmd.Flags().Uint("port", 8999, "P2P UDP listener port")
	p2pBootstrap = SpyCmd.Flags().String("bootstrap", "", "P2P bootstrap peers (comma-separated)")

	statusAddr = SpyCmd.Flags().String("statusAddr", "[::]:6060", "Listen address for status server (disabled if blank)")

	nodeKeyPath = SpyCmd.Flags().String("nodeKey", "", "Path to node key (will be generated if it doesn't exist)")

	logLevel = SpyCmd.Flags().String("logLevel", "info", "Logging level (debug, info, warn, error, dpanic, panic, fatal)")

	spyRPC = SpyCmd.Flags().String("spyRPC", "", "Listen address for gRPC interface")
}

// SpyCmd represents the node command
var SpyCmd = &cobra.Command{
	Use:   "spy",
	Short: "Run gossip spy client",
	Run:   runSpy,
}

type spyServer struct {
	spyv1.UnimplementedSpyRPCServiceServer
	logger          *zap.Logger
	subsSignedVaa   map[string]*subscriptionSignedVaa
	subsSignedVaaMu sync.Mutex
	subsAllVaa      map[string]*subscriptionAllVaa
	subsAllVaaMu    sync.Mutex
}

type message struct {
	vaaBytes []byte
}

type filterSignedVaa struct {
	chainId     vaa.ChainID
	emitterAddr vaa.Address
}
type subscriptionSignedVaa struct {
	filters []filterSignedVaa
	ch      chan message
}
type subscriptionAllVaa struct {
	filters []spyv1.FilterEntry
	ch      chan *spyv1.SubscribeSignedVAAByTypeResponse
}

func subscriptionId() string {
	return uuid.New().String()
}

func decodeEmitterAddr(hexAddr string) (vaa.Address, error) {
	address, err := hex.DecodeString(hexAddr)
	if err != nil {
		return vaa.Address{}, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to decode address: %v", err))
	}
	if len(address) != 32 {
		return vaa.Address{}, status.Error(codes.InvalidArgument, "address must be 32 bytes")
	}

	addr := vaa.Address{}
	copy(addr[:], address)

	return addr, nil
}

// TEMP: this can be imported from /wormhole-foundation/wormhole/sdk/vaa once it lands.
// StringToHash converts a hex-encoded string into a common.Hash
func StringToHash(value string) (eth_common.Hash, error) {
	var tx eth_common.Hash

	// Make sure we have enough to decode
	if len(value) < 2 {
		return tx, fmt.Errorf("value must be at least 1 byte")
	}

	// Trim any preceding "0x" to the address
	value = strings.TrimPrefix(value, "0x")

	res, err := hex.DecodeString(value)
	if err != nil {
		return tx, err
	}

	tx = eth_common.BytesToHash(res)

	return tx, nil
}

func (s *spyServer) PublishSignedVAA(vaaBytes []byte) error {
	s.subsSignedVaaMu.Lock()
	defer s.subsSignedVaaMu.Unlock()

	var v *vaa.VAA

	for _, sub := range s.subsSignedVaa {
		if len(sub.filters) == 0 {
			sub.ch <- message{vaaBytes: vaaBytes}
		} else {
			if v == nil {
				var err error
				v, err = vaa.Unmarshal(vaaBytes)
				if err != nil {
					return err
				}
			}

			for _, fi := range sub.filters {
				if fi.chainId == v.EmitterChain && fi.emitterAddr == v.EmitterAddress {
					sub.ch <- message{vaaBytes: vaaBytes}
				}
			}
		}
	}

	return nil
}

// TEMP - commented out until the BatchVAA structs land on dev.v2.
// // TransactionIdMatches decodes both transactionIDs and checks if they are the same.
// func TransactionIdMatches(batch vaa.BatchVAA, t *spyv1.BatchFilter) (matches bool, err error) {

// 	// first check if the transaction IDs match
// 	filterHash, err := StringToHash(t.TransactionId)
// 	if err != nil {
// 		return matches, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to decode filter's txId: %v", err))
// 	}
// 	obsHash, err := StringToHash(batch.TransactionID)
// 	if err != nil {
// 		return matches, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to decode BatchVAA's txId: %v", err))
// 	}

// 	matches = filterHash == obsHash
// 	return matches, nil
// }

// // BatchMatchFilter asserts that the obervation matches the values of the filter.
// func BatchMatchesFilter(batch vaa.BatchVAA, f *spyv1.BatchFilter) (matches bool, err error) {

// 	txMatch, err := TransactionIdMatches(batch, f)
// 	if err != nil {
// 		return matches, err
// 	}

// 	if txMatch {
// 		// the BatchVAA's transaction ID matches the transaction ID of this filter.
// 		// now check if the other properties of the filter match.

// 		if obs, ok := batch.Observations[0]; ok {
// 			obsVAA := obs.Observation

// 			if obsVAA.EmitterChain == vaa.ChainID(f.ChainId) {
// 				// the emitter chain of the observation matches the filter

// 				if f.Nonce >= 1 {
// 					// filter has a nonce, so make sure it matches
// 					if obsVAA.Nonce == f.Nonce {
// 						// filter's nonce matches the nonce of the obervations. send it.
// 						matches = true
// 						return matches, err
// 					}

// 				} else {
// 					// filter does not have a nonce, everything else matched, send it.
// 					matches = true
// 					return matches, err
// 				}
// 			}
// 		}
// 	}
// 	return matches, err
// }

func (s *spyServer) PublishSignedVAAByType(vaaBytes []byte) error {
	s.subsAllVaaMu.Lock()
	defer s.subsAllVaaMu.Unlock()

	// TEMP - commented out until the BatchVAA structs land on dev.v2.
	// var b *vaa.BatchVaa

	v, _ := vaa.Unmarshal(vaaBytes)
	// do nothing with the error, until we can try to unmarshal the bytes as a batch.
	// if err != nil {
	// 	// check if it is a batch

	// 	// TEMP - commented out until the BatchVAA structs land on dev.v2.
	// 	// it is not a VAA, try unmarshaling to a BatchVAA
	// 	b, err = vaa.UnmarshalBatch(vaaBytes)
	// 	if err != nil {
	// 		// it is not either type of VAA we know, nothing to do.
	// 		return err
	// 	}
	// }

	// create the response(s) that will get sent out if this VAA satisfies a subscription.

	// create the top-level response struct that is agnostic to the VAA type
	var topRes *spyv1.SubscribeSignedVAAByTypeResponse

	// if v has values, the unmarshal was successful.
	if v.Payload != nil {
		// resData is the lowest level proto struct, it holds the byte data for whatever
		// type of response it is (VAA in this case).
		resData := &spyv1.SubscribeSignedVAAResponse{
			VaaBytes: vaaBytes,
		}

		// resType defines what struct vaa will be retuned below, res of type SignedVaa
		resType := &spyv1.SubscribeSignedVAAByTypeResponse_SignedVaa{
			SignedVaa: resData,
		}

		// topRes is the highest level proto struct, the response to the subscription
		topRes = &spyv1.SubscribeSignedVAAByTypeResponse{
			// VaaType: &spyv1.SubscribeSignedVAAByTypeResponse_SignedVaa{
			VaaType: resType,
		}

		// the proto is fully constructed ready to send.
	}

	// TEMP - commented out until the BatchVAA structs land on dev.v2.
	// // if b has vaules, the unmarshal was successful.
	// if len(b) > 0 {
	// 	// resData is the lowest level proto struct, it holds the byte data for whatever
	// 	// type of response it is (BatchVAA in this case).
	// 	resData := &spyv1.SubscribeSignedBatchVAAResponse{
	// 		BatchVaa: vaaBytes,
	// 	}

	// 	// resType defines what struct vaa will be retuned below, res of type SignedBatchVaa.
	// 	resType := &spyv1.SubscribeSignedVAAByTypeResponse_SignedBatchVaa{
	// 		SignedBatchVaa: resData,
	// 	}

	// 	// topRes is the highest level proto struct, the response to the subscription
	// 	topRes = &spyv1.SubscribeSignedVAAByTypeResponse{
	// 		// VaaType: &spyv1.SubscribeSignedVAAByTypeResponse_SignedVaa{
	// 		VaaType: resType,
	// 	}
	// 	// proto is fully constructed ready to send.
	// }

	for _, sub := range s.subsAllVaa {
		if len(sub.filters) == 0 {
			// this subscription has no filters, send it.
			sub.ch <- topRes

		} else {
			// this subscription has filters.
			if v.Payload != nil {

				for i := range sub.filters {
					// get the filter by index/GetFilter rather than from range,
					// so as to not copy the mutex within the Filter.
					filter := sub.filters[i].GetFilter()
					switch t := filter.(type) {
					case *spyv1.FilterEntry_EmitterFilter:

						// TEMP - commented out until the BatchVAA structs land on dev.v2.
						// // take an observation from the batch's list, set it to VAA
						// // so it's data can be considered for a filter match.
						// if obs, ok := b.Observation[1]; ok {
						// 	v = obs.Observation
						// }

						if len(v.EmitterAddress) > 0 && v.EmitterChain > 0 {
							// emitter chain of the vaa and filter match

							addr, err := decodeEmitterAddr(t.EmitterFilter.EmitterAddress)
							if err != nil {
								return status.Error(codes.InvalidArgument, fmt.Sprintf("failed to decode emitter address: %v", err))
							}
							if v.EmitterChain == vaa.ChainID(t.EmitterFilter.ChainId) && addr == v.EmitterAddress {
								// it is a match, send the response
								sub.ch <- topRes
							}
						}

					// TEMP - commented out until the BatchVAA structs land on dev.v2.
					// case *spyv1.FilterEntry_BatchFilter:
					// 	match, err := BatchMatchesFilter(b, t.BatchFilter)
					// 	if err != nil {
					// 		return err
					// 	}
					// 	if match {
					// 		sub.ch <- topRes
					// 	}

					// case *spyv1.FilterEntry_BatchTransactionFilter:
					// 	// make a BatchFilter struct from the BatchTransactionFilter since the latter is
					// 	// a subset of the former's properties, so we can use TransactionIdMatches.
					// 	batchFilter := &spyv1.BatchFilter{
					// 		ChainId: t.BatchTransactionFilter.ChainId,
					// 		TransactionId: t.BatchTransactionFilter.TransactionId,
					// 	}

					// 	match, err := BatchMatchesFilter(b, batchFilter)
					// 	if err != nil {
					// 		return err
					// 	}
					// 	if match {
					// 		sub.ch <- topRes
					// 	}
					default:
						return status.Error(codes.InvalidArgument, "unsupported filter type")
					}
				}
			}
		}
	}

	return nil
}

func (s *spyServer) SubscribeSignedVAA(req *spyv1.SubscribeSignedVAARequest, resp spyv1.SpyRPCService_SubscribeSignedVAAServer) error {
	var fi []filterSignedVaa
	if req.Filters != nil {
		for _, f := range req.Filters {
			switch t := f.Filter.(type) {
			case *spyv1.FilterEntry_EmitterFilter:
				addr, err := decodeEmitterAddr(t.EmitterFilter.EmitterAddress)
				if err != nil {
					return status.Error(codes.InvalidArgument, fmt.Sprintf("failed to decode emitter address: %v", err))
				}
				fi = append(fi, filterSignedVaa{
					chainId:     vaa.ChainID(t.EmitterFilter.ChainId),
					emitterAddr: addr,
				})
			default:
				return status.Error(codes.InvalidArgument, "unsupported filter type")
			}
		}
	}

	s.subsSignedVaaMu.Lock()
	id := subscriptionId()
	sub := &subscriptionSignedVaa{
		ch:      make(chan message, 1),
		filters: fi,
	}
	s.subsSignedVaa[id] = sub
	s.subsSignedVaaMu.Unlock()

	defer func() {
		s.subsSignedVaaMu.Lock()
		defer s.subsSignedVaaMu.Unlock()
		delete(s.subsSignedVaa, id)
	}()

	for {
		select {
		case <-resp.Context().Done():
			return resp.Context().Err()
		case msg := <-sub.ch:
			if err := resp.Send(&spyv1.SubscribeSignedVAAResponse{
				VaaBytes: msg.vaaBytes,
			}); err != nil {
				return err
			}
		}
	}
}

// SubscribeSignedVAAByType fields requests for subscriptions. Each new subscription adds a channel and request params (filters)
// to the map of active subscriptions.
func (s *spyServer) SubscribeSignedVAAByType(req *spyv1.SubscribeSignedVAAByTypeRequest, resp spyv1.SpyRPCService_SubscribeSignedVAAByTypeServer) error {
	var fi []spyv1.FilterEntry
	if req.Filters != nil {
		for _, f := range req.Filters {
			switch t := f.Filter.(type) {

			case *spyv1.FilterEntry_EmitterFilter:
				// validate the emitter address is valid by decoding it
				_, err := decodeEmitterAddr(t.EmitterFilter.EmitterAddress)
				if err != nil {
					return status.Error(codes.InvalidArgument, fmt.Sprintf("failed to decode emitter address: %v", err))
				}
				fi = append(fi, spyv1.FilterEntry{Filter: t})

			// TEMP - commented out until the BatchVAA structs land on dev.v2.
			// case *spyv1.FilterEntry_BatchFilter:
			// 	// validate the TransactionId is valid by decoding it.
			// 	_, err := StringToHash(t.BatchFilter.TransactionId)
			// 	if err != nil {
			// 		return status.Error(codes.InvalidArgument, fmt.Sprintf("failed to decode filter's txId: %v", err))
			// 	}
			// 	fi = append(fi, spyv1.FilterEntry{Filter: t})

			// case *spyv1.FilterEntry_BatchTransactionFilter:
			// 	// validate the TransactionId is valid by decoding it.
			// 	_, err := StringToHash(t.BatchTransactionFilter.TransactionId)
			// 	if err != nil {
			// 		return status.Error(codes.InvalidArgument, fmt.Sprintf("failed to decode filter's txId: %v", err))
			// 	}
			// 	fi = append(fi, spyv1.FilterEntry{Filter: t})
			default:
				return status.Error(codes.InvalidArgument, "unsupported filter type")
			}
		}
	}

	s.subsAllVaaMu.Lock()
	id := subscriptionId()
	sub := &subscriptionAllVaa{
		ch:      make(chan *spyv1.SubscribeSignedVAAByTypeResponse, 1),
		filters: fi,
	}
	s.subsAllVaa[id] = sub
	s.subsAllVaaMu.Unlock()

	defer func() {
		s.subsAllVaaMu.Lock()
		defer s.subsAllVaaMu.Unlock()
		delete(s.subsAllVaa, id)
	}()

	for {
		select {
		case <-resp.Context().Done():
			return resp.Context().Err()
		case msg := <-sub.ch:
			if err := resp.Send(msg); err != nil {
				return err
			}
		}
	}
}

func newSpyServer(logger *zap.Logger) *spyServer {
	return &spyServer{
		logger:        logger.Named("spyserver"),
		subsSignedVaa: make(map[string]*subscriptionSignedVaa),
		subsAllVaa:    make(map[string]*subscriptionAllVaa),
	}
}

func spyServerRunnable(s *spyServer, logger *zap.Logger, listenAddr string) (supervisor.Runnable, *grpc.Server, error) {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen: %w", err)
	}

	logger.Info("publicrpc server listening", zap.String("addr", l.Addr().String()))

	grpcServer := common.NewInstrumentedGRPCServer(logger)
	spyv1.RegisterSpyRPCServiceServer(grpcServer, s)

	return supervisor.GRPCServer(grpcServer, l, false), grpcServer, nil
}

func runSpy(cmd *cobra.Command, args []string) {
	common.SetRestrictiveUmask()

	lvl, err := ipfslog.LevelFromString(*logLevel)
	if err != nil {
		fmt.Println("Invalid log level")
		os.Exit(1)
	}

	logger := ipfslog.Logger("wormhole-spy").Desugar()

	ipfslog.SetAllLoggers(lvl)

	// Status server
	if *statusAddr != "" {
		router := mux.NewRouter()

		router.Handle("/metrics", promhttp.Handler())

		go func() {
			logger.Info("status server listening on [::]:6060")
			logger.Error("status server crashed", zap.Error(http.ListenAndServe(*statusAddr, router)))
		}()
	}

	// Verify flags

	if *nodeKeyPath == "" {
		logger.Fatal("Please specify --nodeKey")
	}
	if *p2pBootstrap == "" {
		logger.Fatal("Please specify --bootstrap")
	}

	// Node's main lifecycle context.
	rootCtx, rootCtxCancel = context.WithCancel(context.Background())
	defer rootCtxCancel()

	// Outbound gossip message queue
	sendC := make(chan []byte)

	// Inbound observations
	obsvC := make(chan *gossipv1.SignedObservation, 50)

	// Inbound observation requests
	obsvReqC := make(chan *gossipv1.ObservationRequest, 50)

	// Inbound signed VAAs
	signedInC := make(chan *gossipv1.SignedVAAWithQuorum, 50)

	// Guardian set state managed by processor
	gst := common.NewGuardianSetState()

	// RPC server
	s := newSpyServer(logger)
	rpcSvc, _, err := spyServerRunnable(s, logger, *spyRPC)
	if err != nil {
		logger.Fatal("failed to start RPC server", zap.Error(err))
	}

	// Ignore observations
	go func() {
		for {
			select {
			case <-rootCtx.Done():
				return
			case <-obsvC:
			}
		}
	}()

	// Ignore observation requests
	// Note: without this, the whole program hangs on observation requests
	go func() {
		for {
			select {
			case <-rootCtx.Done():
				return
			case <-obsvReqC:
			}
		}
	}()

	// Log signed VAAs
	go func() {
		for {
			select {
			case <-rootCtx.Done():
				return
			case v := <-signedInC:
				logger.Info("Received signed VAA",
					zap.Any("vaa", v.Vaa))
				if err := s.PublishSignedVAA(v.Vaa); err != nil {
					logger.Error("failed to publish signed VAA", zap.Error(err))
				}
				if err := s.PublishSignedVAAByType(v.Vaa); err != nil {
					logger.Error("failed to publish signed VAA by type", zap.Error(err))
				}
			}
		}
	}()

	// Load p2p private key
	var priv crypto.PrivKey
	priv, err = common.GetOrCreateNodeKey(logger, *nodeKeyPath)
	if err != nil {
		logger.Fatal("Failed to load node key", zap.Error(err))
	}

	// Run supervisor.
	supervisor.New(rootCtx, logger, func(ctx context.Context) error {
		if err := supervisor.Run(ctx, "p2p", p2p.Run(obsvC, obsvReqC, nil, sendC, signedInC, priv, nil, gst, *p2pPort, *p2pNetworkID, *p2pBootstrap, "", false, rootCtxCancel, nil)); err != nil {
			return err
		}

		if err := supervisor.Run(ctx, "spyrpc", rpcSvc); err != nil {
			return err
		}

		logger.Info("Started internal services")

		<-ctx.Done()
		return nil
	},
		// It's safer to crash and restart the process in case we encounter a panic,
		// rather than attempting to reschedule the runnable.
		supervisor.WithPropagatePanic)

	<-rootCtx.Done()
	logger.Info("root context cancelled, exiting...")
	// TODO: wait for things to shut down gracefully
}
