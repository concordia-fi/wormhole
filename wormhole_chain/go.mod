module github.com/wormhole-foundation/wormhole-chain

go 1.16

require (
	github.com/CosmWasm/wasmd v0.28.0
	github.com/cosmos/cosmos-sdk v0.45.8
	github.com/cosmos/ibc-go/v3 v3.3.0
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/ethereum/go-ethereum v1.10.21
	github.com/gogo/protobuf v1.3.3
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/prometheus/client_golang v1.12.2
	github.com/spf13/cast v1.5.0
	github.com/spf13/cobra v1.5.0
	github.com/stretchr/testify v1.8.0
	github.com/tendermint/spm v0.1.9
	github.com/tendermint/tendermint v0.34.21
	github.com/tendermint/tm-db v0.6.7
	github.com/wormhole-foundation/wormhole/sdk v0.0.0-20220921004715-3103e59217da
	google.golang.org/genproto v0.0.0-20220725144611-272f38e5d71b
	google.golang.org/grpc v1.48.0
	nhooyr.io/websocket v1.8.7 // indirect
)

replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/CosmWasm/wasmd v0.28.0 => github.com/wormhole-foundation/wasmd v0.28.0-wormhole
	github.com/cosmos/cosmos-sdk v0.45.7 => github.com/wormhole-foundation/cosmos-sdk v0.45.7-wormhole
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/wormhole-foundation/wormhole/sdk => ../sdk
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
