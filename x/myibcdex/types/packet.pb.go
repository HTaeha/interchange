// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: myibcdex/packet.proto

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

type MyibcdexPacketData struct {
	// Types that are valid to be assigned to Packet:
	//	*MyibcdexPacketData_NoData
	//	*MyibcdexPacketData_CreatePairPacket
	Packet isMyibcdexPacketData_Packet `protobuf_oneof:"packet"`
}

func (m *MyibcdexPacketData) Reset()         { *m = MyibcdexPacketData{} }
func (m *MyibcdexPacketData) String() string { return proto.CompactTextString(m) }
func (*MyibcdexPacketData) ProtoMessage()    {}
func (*MyibcdexPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_63622f42d7523591, []int{0}
}
func (m *MyibcdexPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MyibcdexPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MyibcdexPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MyibcdexPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyibcdexPacketData.Merge(m, src)
}
func (m *MyibcdexPacketData) XXX_Size() int {
	return m.Size()
}
func (m *MyibcdexPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_MyibcdexPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_MyibcdexPacketData proto.InternalMessageInfo

type isMyibcdexPacketData_Packet interface {
	isMyibcdexPacketData_Packet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type MyibcdexPacketData_NoData struct {
	NoData *NoData `protobuf:"bytes,1,opt,name=noData,proto3,oneof" json:"noData,omitempty"`
}
type MyibcdexPacketData_CreatePairPacket struct {
	CreatePairPacket *CreatePairPacketData `protobuf:"bytes,2,opt,name=createPairPacket,proto3,oneof" json:"createPairPacket,omitempty"`
}

func (*MyibcdexPacketData_NoData) isMyibcdexPacketData_Packet()           {}
func (*MyibcdexPacketData_CreatePairPacket) isMyibcdexPacketData_Packet() {}

func (m *MyibcdexPacketData) GetPacket() isMyibcdexPacketData_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *MyibcdexPacketData) GetNoData() *NoData {
	if x, ok := m.GetPacket().(*MyibcdexPacketData_NoData); ok {
		return x.NoData
	}
	return nil
}

func (m *MyibcdexPacketData) GetCreatePairPacket() *CreatePairPacketData {
	if x, ok := m.GetPacket().(*MyibcdexPacketData_CreatePairPacket); ok {
		return x.CreatePairPacket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MyibcdexPacketData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MyibcdexPacketData_NoData)(nil),
		(*MyibcdexPacketData_CreatePairPacket)(nil),
	}
}

type NoData struct {
}

func (m *NoData) Reset()         { *m = NoData{} }
func (m *NoData) String() string { return proto.CompactTextString(m) }
func (*NoData) ProtoMessage()    {}
func (*NoData) Descriptor() ([]byte, []int) {
	return fileDescriptor_63622f42d7523591, []int{1}
}
func (m *NoData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NoData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NoData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoData.Merge(m, src)
}
func (m *NoData) XXX_Size() int {
	return m.Size()
}
func (m *NoData) XXX_DiscardUnknown() {
	xxx_messageInfo_NoData.DiscardUnknown(m)
}

var xxx_messageInfo_NoData proto.InternalMessageInfo

// CreatePairPacketData defines a struct for the packet payload
type CreatePairPacketData struct {
	SourceDenom string `protobuf:"bytes,1,opt,name=sourceDenom,proto3" json:"sourceDenom,omitempty"`
	TargetDenom string `protobuf:"bytes,2,opt,name=targetDenom,proto3" json:"targetDenom,omitempty"`
}

func (m *CreatePairPacketData) Reset()         { *m = CreatePairPacketData{} }
func (m *CreatePairPacketData) String() string { return proto.CompactTextString(m) }
func (*CreatePairPacketData) ProtoMessage()    {}
func (*CreatePairPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_63622f42d7523591, []int{2}
}
func (m *CreatePairPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePairPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePairPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreatePairPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePairPacketData.Merge(m, src)
}
func (m *CreatePairPacketData) XXX_Size() int {
	return m.Size()
}
func (m *CreatePairPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePairPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePairPacketData proto.InternalMessageInfo

func (m *CreatePairPacketData) GetSourceDenom() string {
	if m != nil {
		return m.SourceDenom
	}
	return ""
}

func (m *CreatePairPacketData) GetTargetDenom() string {
	if m != nil {
		return m.TargetDenom
	}
	return ""
}

// CreatePairPacketAck defines a struct for the packet acknowledgment
type CreatePairPacketAck struct {
}

func (m *CreatePairPacketAck) Reset()         { *m = CreatePairPacketAck{} }
func (m *CreatePairPacketAck) String() string { return proto.CompactTextString(m) }
func (*CreatePairPacketAck) ProtoMessage()    {}
func (*CreatePairPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_63622f42d7523591, []int{3}
}
func (m *CreatePairPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePairPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePairPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreatePairPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePairPacketAck.Merge(m, src)
}
func (m *CreatePairPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *CreatePairPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePairPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePairPacketAck proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MyibcdexPacketData)(nil), "coreators.interchange.myibcdex.MyibcdexPacketData")
	proto.RegisterType((*NoData)(nil), "coreators.interchange.myibcdex.NoData")
	proto.RegisterType((*CreatePairPacketData)(nil), "coreators.interchange.myibcdex.CreatePairPacketData")
	proto.RegisterType((*CreatePairPacketAck)(nil), "coreators.interchange.myibcdex.CreatePairPacketAck")
}

func init() { proto.RegisterFile("myibcdex/packet.proto", fileDescriptor_63622f42d7523591) }

var fileDescriptor_63622f42d7523591 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xad, 0xcc, 0x4c,
	0x4a, 0x4e, 0x49, 0xad, 0xd0, 0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x4b, 0xce, 0x2f, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x2a, 0xd6, 0xcb, 0xcc, 0x2b,
	0x49, 0x2d, 0x4a, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0xd5, 0x83, 0x29, 0x56, 0x3a, 0xc2, 0xc8, 0x25,
	0xe4, 0x0b, 0xe5, 0x04, 0x80, 0x35, 0xba, 0x24, 0x96, 0x24, 0x0a, 0x39, 0x70, 0xb1, 0xe5, 0xe5,
	0x83, 0x58, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x6a, 0x7a, 0xf8, 0xcd, 0xd1, 0xf3, 0x03,
	0xab, 0xf6, 0x60, 0x08, 0x82, 0xea, 0x13, 0x4a, 0xe2, 0x12, 0x48, 0x06, 0xe9, 0x48, 0x0d, 0x48,
	0xcc, 0x2c, 0x82, 0x98, 0x2c, 0xc1, 0x04, 0x36, 0xcb, 0x84, 0x90, 0x59, 0xce, 0x68, 0xfa, 0xa0,
	0x26, 0x63, 0x98, 0xe7, 0xc4, 0xc1, 0xc5, 0x06, 0xf1, 0xac, 0x12, 0x07, 0x17, 0x1b, 0xc4, 0x05,
	0x4a, 0x51, 0x5c, 0x22, 0xd8, 0xf4, 0x0b, 0x29, 0x70, 0x71, 0x17, 0xe7, 0x97, 0x16, 0x25, 0xa7,
	0xba, 0xa4, 0xe6, 0xe5, 0xe7, 0x82, 0xbd, 0xc5, 0x19, 0x84, 0x2c, 0x04, 0x52, 0x51, 0x92, 0x58,
	0x94, 0x9e, 0x5a, 0x02, 0x51, 0xc1, 0x04, 0x51, 0x81, 0x24, 0xa4, 0x24, 0xca, 0x25, 0x8c, 0x6e,
	0xb6, 0x63, 0x72, 0xb6, 0x93, 0xf7, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78,
	0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44,
	0x19, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xc3, 0x3d, 0xad, 0x8f,
	0xe4, 0x69, 0xfd, 0x0a, 0x7d, 0x78, 0xbc, 0x95, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xe3,
	0xcd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xa9, 0xe6, 0x65, 0xd0, 0x01, 0x00, 0x00,
}

func (m *MyibcdexPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MyibcdexPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MyibcdexPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Packet != nil {
		{
			size := m.Packet.Size()
			i -= size
			if _, err := m.Packet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *MyibcdexPacketData_NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MyibcdexPacketData_NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NoData != nil {
		{
			size, err := m.NoData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *MyibcdexPacketData_CreatePairPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MyibcdexPacketData_CreatePairPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CreatePairPacket != nil {
		{
			size, err := m.CreatePairPacket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *NoData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *CreatePairPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePairPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreatePairPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TargetDenom) > 0 {
		i -= len(m.TargetDenom)
		copy(dAtA[i:], m.TargetDenom)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.TargetDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourceDenom) > 0 {
		i -= len(m.SourceDenom)
		copy(dAtA[i:], m.SourceDenom)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.SourceDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CreatePairPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePairPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreatePairPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MyibcdexPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Packet != nil {
		n += m.Packet.Size()
	}
	return n
}

func (m *MyibcdexPacketData_NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoData != nil {
		l = m.NoData.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *MyibcdexPacketData_CreatePairPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreatePairPacket != nil {
		l = m.CreatePairPacket.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *CreatePairPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourceDenom)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.TargetDenom)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *CreatePairPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MyibcdexPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: MyibcdexPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MyibcdexPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NoData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &MyibcdexPacketData_NoData{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatePairPacket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &CreatePairPacketData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &MyibcdexPacketData_CreatePairPacket{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *NoData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: NoData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NoData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *CreatePairPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: CreatePairPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePairPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *CreatePairPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: CreatePairPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePairPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
