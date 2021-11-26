// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: myibcdex/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
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

func init() { proto.RegisterFile("myibcdex/query.proto", fileDescriptor_58ca6b6d334c15a5) }

var fileDescriptor_58ca6b6d334c15a5 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0x31, 0x4e, 0xc4, 0x30,
	0x10, 0x85, 0xe1, 0xa4, 0x00, 0xa4, 0x94, 0x88, 0x2a, 0x42, 0x3e, 0x00, 0x85, 0x47, 0x81, 0x1b,
	0xd0, 0x52, 0xd1, 0xd2, 0x8d, 0xcd, 0xc8, 0xb1, 0x44, 0x3c, 0xc6, 0x9e, 0xa0, 0xe4, 0x16, 0x1c,
	0x6b, 0xcb, 0x94, 0x5b, 0xae, 0x92, 0x8b, 0xac, 0x36, 0xd1, 0x46, 0xdb, 0x7f, 0xf3, 0xf4, 0x4f,
	0xf5, 0xd4, 0x8d, 0xde, 0xd8, 0x6f, 0x1a, 0xe0, 0xb7, 0xa7, 0x34, 0xea, 0x98, 0x58, 0xf8, 0x51,
	0x59, 0x4e, 0x84, 0xc2, 0x29, 0x6b, 0x1f, 0x84, 0x92, 0x6d, 0x31, 0x38, 0xd2, 0x57, 0x5b, 0x3f,
	0x3b, 0x66, 0xf7, 0x43, 0x80, 0xd1, 0x03, 0x86, 0xc0, 0x82, 0xe2, 0x39, 0xe4, 0xed, 0xba, 0x7e,
	0xb1, 0x9c, 0x3b, 0xce, 0x60, 0x30, 0xd3, 0x36, 0x0b, 0x7f, 0x8d, 0x21, 0xc1, 0x06, 0x22, 0x3a,
	0x1f, 0x56, 0xbc, 0xd9, 0xd7, 0x87, 0xea, 0xee, 0xf3, 0x22, 0xde, 0x3f, 0x0e, 0xb3, 0x2a, 0xa7,
	0x59, 0x95, 0xa7, 0x59, 0x95, 0xff, 0x8b, 0x2a, 0xa6, 0x45, 0x15, 0xc7, 0x45, 0x15, 0x5f, 0x8d,
	0xf3, 0xd2, 0xf6, 0x46, 0x5b, 0xee, 0x60, 0xef, 0x82, 0x9b, 0x2e, 0x18, 0x60, 0xff, 0x42, 0xc6,
	0x48, 0xd9, 0xdc, 0xaf, 0xe3, 0x6f, 0xe7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0x34, 0xd7, 0x82,
	0xde, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "coreators.interchange.myibcdex.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "myibcdex/query.proto",
}
