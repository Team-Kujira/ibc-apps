// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc_hooks/v7/ibc_hooks.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the set of on-chain interchain query parameters.
type Params struct {
	// optionally set axelar account & channel data to pass origin source_address to info.sender
	Axelar *Axelar `protobuf:"bytes,1,opt,name=axelar,proto3" json:"axelar,omitempty" yaml:"axelar"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ab0203f53c12df8, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAxelar() *Axelar {
	if m != nil {
		return m.Axelar
	}
	return nil
}

type Axelar struct {
	// the canonical GMP account on the Axelar chain
	GmpAccount string `protobuf:"bytes,1,opt,name=gmp_account,json=gmpAccount,proto3" json:"gmp_account,omitempty" yaml:"gmp_account"`
	// the channel ID that the chain is connected via
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty" yaml:"channel_id"`
}

func (m *Axelar) Reset()         { *m = Axelar{} }
func (m *Axelar) String() string { return proto.CompactTextString(m) }
func (*Axelar) ProtoMessage()    {}
func (*Axelar) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ab0203f53c12df8, []int{1}
}
func (m *Axelar) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Axelar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Axelar.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Axelar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Axelar.Merge(m, src)
}
func (m *Axelar) XXX_Size() int {
	return m.Size()
}
func (m *Axelar) XXX_DiscardUnknown() {
	xxx_messageInfo_Axelar.DiscardUnknown(m)
}

var xxx_messageInfo_Axelar proto.InternalMessageInfo

func (m *Axelar) GetGmpAccount() string {
	if m != nil {
		return m.GmpAccount
	}
	return ""
}

func (m *Axelar) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "ibc_hooks.v7.Params")
	proto.RegisterType((*Axelar)(nil), "ibc_hooks.v7.Axelar")
}

func init() { proto.RegisterFile("ibc_hooks/v7/ibc_hooks.proto", fileDescriptor_9ab0203f53c12df8) }

var fileDescriptor_9ab0203f53c12df8 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0x4c, 0x4a, 0x8e,
	0xcf, 0xc8, 0xcf, 0xcf, 0x2e, 0xd6, 0x2f, 0x33, 0xd7, 0x87, 0x73, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x78, 0x10, 0x02, 0x65, 0xe6, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x09, 0x7d,
	0x10, 0x0b, 0xa2, 0x46, 0xc9, 0x93, 0x8b, 0x2d, 0x20, 0xb1, 0x28, 0x31, 0xb7, 0x58, 0xc8, 0x9e,
	0x8b, 0x2d, 0xb1, 0x22, 0x35, 0x27, 0xb1, 0x48, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x44,
	0x0f, 0x59, 0xbb, 0x9e, 0x23, 0x58, 0xce, 0x49, 0xf0, 0xd3, 0x3d, 0x79, 0xde, 0xca, 0xc4, 0xdc,
	0x1c, 0x2b, 0x25, 0x88, 0x6a, 0xa5, 0x20, 0xa8, 0x36, 0xa5, 0x72, 0x2e, 0x36, 0x88, 0x22, 0x21,
	0x73, 0x2e, 0xee, 0xf4, 0xdc, 0x82, 0xf8, 0xc4, 0xe4, 0xe4, 0xfc, 0xd2, 0xbc, 0x12, 0xb0, 0x79,
	0x9c, 0x4e, 0x62, 0x9f, 0xee, 0xc9, 0x0b, 0x41, 0x74, 0x22, 0x49, 0x2a, 0x05, 0x71, 0xa5, 0xe7,
	0x16, 0x38, 0x42, 0x38, 0x42, 0x26, 0x5c, 0x5c, 0xc9, 0x19, 0x89, 0x79, 0x79, 0xa9, 0x39, 0xf1,
	0x99, 0x29, 0x12, 0x4c, 0x60, 0x7d, 0xa2, 0x9f, 0xee, 0xc9, 0x0b, 0x42, 0xf4, 0x21, 0xe4, 0x94,
	0x82, 0x38, 0xa1, 0x1c, 0xcf, 0x14, 0x27, 0xff, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96,
	0x63, 0x88, 0x32, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce,
	0x2f, 0xce, 0xcd, 0x2f, 0x06, 0x05, 0x92, 0x6e, 0x62, 0x41, 0x41, 0xb1, 0x7e, 0x6e, 0x7e, 0x4a,
	0x69, 0x4e, 0x6a, 0xb1, 0x3e, 0x4a, 0x10, 0x96, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xc3,
	0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x83, 0x4c, 0xe3, 0x5f, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Axelar != nil {
		{
			size, err := m.Axelar.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIbcHooks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Axelar) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Axelar) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Axelar) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintIbcHooks(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.GmpAccount) > 0 {
		i -= len(m.GmpAccount)
		copy(dAtA[i:], m.GmpAccount)
		i = encodeVarintIbcHooks(dAtA, i, uint64(len(m.GmpAccount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIbcHooks(dAtA []byte, offset int, v uint64) int {
	offset -= sovIbcHooks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Axelar != nil {
		l = m.Axelar.Size()
		n += 1 + l + sovIbcHooks(uint64(l))
	}
	return n
}

func (m *Axelar) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GmpAccount)
	if l > 0 {
		n += 1 + l + sovIbcHooks(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovIbcHooks(uint64(l))
	}
	return n
}

func sovIbcHooks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIbcHooks(x uint64) (n int) {
	return sovIbcHooks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbcHooks
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Axelar", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcHooks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIbcHooks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIbcHooks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Axelar == nil {
				m.Axelar = &Axelar{}
			}
			if err := m.Axelar.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIbcHooks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbcHooks
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
func (m *Axelar) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbcHooks
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
			return fmt.Errorf("proto: Axelar: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Axelar: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GmpAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcHooks
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
				return ErrInvalidLengthIbcHooks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbcHooks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GmpAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcHooks
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
				return ErrInvalidLengthIbcHooks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbcHooks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIbcHooks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbcHooks
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
func skipIbcHooks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIbcHooks
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
					return 0, ErrIntOverflowIbcHooks
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
					return 0, ErrIntOverflowIbcHooks
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
				return 0, ErrInvalidLengthIbcHooks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIbcHooks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIbcHooks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIbcHooks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIbcHooks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIbcHooks = fmt.Errorf("proto: unexpected end of group")
)
