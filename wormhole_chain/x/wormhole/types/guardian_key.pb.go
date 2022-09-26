// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wormhole/guardian_key.proto

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

type GuardianKey struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *GuardianKey) Reset()         { *m = GuardianKey{} }
func (m *GuardianKey) String() string { return proto.CompactTextString(m) }
func (*GuardianKey) ProtoMessage()    {}
func (*GuardianKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_87576a0b45454f44, []int{0}
}
func (m *GuardianKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GuardianKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GuardianKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GuardianKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuardianKey.Merge(m, src)
}
func (m *GuardianKey) XXX_Size() int {
	return m.Size()
}
func (m *GuardianKey) XXX_DiscardUnknown() {
	xxx_messageInfo_GuardianKey.DiscardUnknown(m)
}

var xxx_messageInfo_GuardianKey proto.InternalMessageInfo

func (m *GuardianKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*GuardianKey)(nil), "wormhole_foundation.wormholechain.wormhole.GuardianKey")
}

func init() { proto.RegisterFile("wormhole/guardian_key.proto", fileDescriptor_87576a0b45454f44) }

var fileDescriptor_87576a0b45454f44 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0xcf, 0x2f, 0xca,
	0xcd, 0xc8, 0xcf, 0x49, 0xd5, 0x4f, 0x2f, 0x4d, 0x2c, 0x4a, 0xc9, 0x4c, 0xcc, 0x8b, 0xcf, 0x4e,
	0xad, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x82, 0x49, 0xc6, 0xa7, 0xe5, 0x97, 0xe6,
	0xa5, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0xe9, 0xc1, 0xc4, 0x92, 0x33, 0x12, 0x33, 0x11, 0x3c, 0x25,
	0x79, 0x2e, 0x6e, 0x77, 0xa8, 0x09, 0xde, 0xa9, 0x95, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x20, 0xa6, 0x53, 0xc4, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c,
	0x37, 0x1e, 0xcb, 0x31, 0x44, 0xd9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7,
	0xea, 0xc3, 0xcc, 0xd3, 0x45, 0xd8, 0x88, 0x10, 0x03, 0x5b, 0xa9, 0x5f, 0x01, 0x17, 0xd0, 0x2f,
	0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xbb, 0xd6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x69,
	0xc8, 0x1b, 0xf9, 0xcc, 0x00, 0x00, 0x00,
}

func (m *GuardianKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GuardianKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GuardianKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintGuardianKey(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGuardianKey(dAtA []byte, offset int, v uint64) int {
	offset -= sovGuardianKey(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GuardianKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovGuardianKey(uint64(l))
	}
	return n
}

func sovGuardianKey(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGuardianKey(x uint64) (n int) {
	return sovGuardianKey(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GuardianKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuardianKey
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
			return fmt.Errorf("proto: GuardianKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GuardianKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardianKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGuardianKey
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGuardianKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGuardianKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGuardianKey
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
func skipGuardianKey(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGuardianKey
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
					return 0, ErrIntOverflowGuardianKey
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
					return 0, ErrIntOverflowGuardianKey
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
				return 0, ErrInvalidLengthGuardianKey
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGuardianKey
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGuardianKey
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGuardianKey        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGuardianKey          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGuardianKey = fmt.Errorf("proto: unexpected end of group")
)
