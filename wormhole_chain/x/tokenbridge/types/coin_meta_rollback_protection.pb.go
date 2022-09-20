// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenbridge/coin_meta_rollback_protection.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CoinMetaRollbackProtection struct {
	Index              string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	LastUpdateSequence uint64 `protobuf:"varint,2,opt,name=lastUpdateSequence,proto3" json:"lastUpdateSequence,omitempty"`
}

func (m *CoinMetaRollbackProtection) Reset()         { *m = CoinMetaRollbackProtection{} }
func (m *CoinMetaRollbackProtection) String() string { return proto.CompactTextString(m) }
func (*CoinMetaRollbackProtection) ProtoMessage()    {}
func (*CoinMetaRollbackProtection) Descriptor() ([]byte, []int) {
	return fileDescriptor_23ec5ccab8f2b4ca, []int{0}
}
func (m *CoinMetaRollbackProtection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoinMetaRollbackProtection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoinMetaRollbackProtection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoinMetaRollbackProtection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinMetaRollbackProtection.Merge(m, src)
}
func (m *CoinMetaRollbackProtection) XXX_Size() int {
	return m.Size()
}
func (m *CoinMetaRollbackProtection) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinMetaRollbackProtection.DiscardUnknown(m)
}

var xxx_messageInfo_CoinMetaRollbackProtection proto.InternalMessageInfo

func (m *CoinMetaRollbackProtection) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *CoinMetaRollbackProtection) GetLastUpdateSequence() uint64 {
	if m != nil {
		return m.LastUpdateSequence
	}
	return 0
}

func init() {
	proto.RegisterType((*CoinMetaRollbackProtection)(nil), "certusone.wormholechain.tokenbridge.CoinMetaRollbackProtection")
}

func init() {
	proto.RegisterFile("tokenbridge/coin_meta_rollback_protection.proto", fileDescriptor_23ec5ccab8f2b4ca)
}

var fileDescriptor_23ec5ccab8f2b4ca = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x4a, 0xc4, 0x40,
	0x10, 0xc7, 0xb3, 0xa2, 0x82, 0x29, 0x83, 0xc5, 0x61, 0xb1, 0x1c, 0xda, 0x5c, 0xe3, 0xa6, 0xf0,
	0x09, 0xd4, 0x5a, 0x90, 0x88, 0x8d, 0x16, 0x61, 0x3f, 0xc6, 0xcb, 0x72, 0x9b, 0x99, 0xb8, 0x99,
	0xc5, 0xf3, 0x2d, 0x7c, 0x2c, 0xcb, 0x2b, 0x2d, 0x25, 0x79, 0x11, 0xb9, 0xf3, 0x1b, 0xae, 0x9b,
	0x0f, 0xf8, 0xcd, 0x6f, 0xfe, 0x79, 0xc9, 0xb4, 0x00, 0x34, 0xd1, 0xbb, 0x39, 0x94, 0x96, 0x3c,
	0xd6, 0x2d, 0xb0, 0xae, 0x23, 0x85, 0x60, 0xb4, 0x5d, 0xd4, 0x5d, 0x24, 0x06, 0xcb, 0x9e, 0x50,
	0xad, 0x4b, 0x2a, 0x4e, 0x2c, 0x44, 0x4e, 0x3d, 0x21, 0xa8, 0x27, 0x8a, 0x6d, 0x43, 0x01, 0x6c,
	0xa3, 0x3d, 0xaa, 0x3f, 0xa0, 0x63, 0x93, 0x1f, 0x5d, 0x92, 0xc7, 0x2b, 0x60, 0x5d, 0x7d, 0x91,
	0xae, 0x7f, 0x40, 0xc5, 0x61, 0xbe, 0xe7, 0xd1, 0xc1, 0x72, 0x22, 0xa6, 0x62, 0x76, 0x50, 0x7d,
	0x36, 0x85, 0xca, 0x8b, 0xa0, 0x7b, 0xbe, 0xed, 0x9c, 0x66, 0xb8, 0x81, 0xc7, 0x04, 0x68, 0x61,
	0xb2, 0x33, 0x15, 0xb3, 0xdd, 0x6a, 0xcb, 0xe6, 0xe2, 0xfe, 0x75, 0x90, 0x62, 0x35, 0x48, 0xf1,
	0x3e, 0x48, 0xf1, 0x32, 0xca, 0x6c, 0x35, 0xca, 0xec, 0x6d, 0x94, 0xd9, 0xdd, 0xf9, 0xdc, 0x73,
	0x93, 0x8c, 0xb2, 0xd4, 0x96, 0xdf, 0x8e, 0xa7, 0x0f, 0x94, 0xd0, 0xe9, 0xf5, 0xfd, 0xdf, 0xd9,
	0x46, 0xbc, 0x5c, 0xfe, 0xcb, 0x80, 0x9f, 0x3b, 0xe8, 0xcd, 0xfe, 0xe6, 0xd9, 0xb3, 0x8f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x52, 0x65, 0x76, 0x1f, 0x1f, 0x01, 0x00, 0x00,
}

func (m *CoinMetaRollbackProtection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoinMetaRollbackProtection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoinMetaRollbackProtection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastUpdateSequence != 0 {
		i = encodeVarintCoinMetaRollbackProtection(dAtA, i, uint64(m.LastUpdateSequence))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintCoinMetaRollbackProtection(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCoinMetaRollbackProtection(dAtA []byte, offset int, v uint64) int {
	offset -= sovCoinMetaRollbackProtection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CoinMetaRollbackProtection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovCoinMetaRollbackProtection(uint64(l))
	}
	if m.LastUpdateSequence != 0 {
		n += 1 + sovCoinMetaRollbackProtection(uint64(m.LastUpdateSequence))
	}
	return n
}

func sovCoinMetaRollbackProtection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCoinMetaRollbackProtection(x uint64) (n int) {
	return sovCoinMetaRollbackProtection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CoinMetaRollbackProtection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoinMetaRollbackProtection
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CoinMetaRollbackProtection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoinMetaRollbackProtection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinMetaRollbackProtection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCoinMetaRollbackProtection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinMetaRollbackProtection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdateSequence", wireType)
			}
			m.LastUpdateSequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinMetaRollbackProtection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastUpdateSequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCoinMetaRollbackProtection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoinMetaRollbackProtection
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCoinMetaRollbackProtection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoinMetaRollbackProtection
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCoinMetaRollbackProtection
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCoinMetaRollbackProtection
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCoinMetaRollbackProtection
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCoinMetaRollbackProtection
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCoinMetaRollbackProtection
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCoinMetaRollbackProtection        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoinMetaRollbackProtection          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCoinMetaRollbackProtection = fmt.Errorf("proto: unexpected end of group")
)
