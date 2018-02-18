// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1.proto

/*
Package instance is a generated protocol buffer package.

It is generated from these files:
	v1.proto

It has these top-level messages:
	CreateRequest
	CreateReply
*/
package instance

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateRequest struct {
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateReply struct {
	InstanceId string `protobuf:"bytes,1,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
}

func (m *CreateReply) Reset()                    { *m = CreateReply{} }
func (m *CreateReply) String() string            { return proto.CompactTextString(m) }
func (*CreateReply) ProtoMessage()               {}
func (*CreateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateReply) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "instance.CreateRequest")
	proto.RegisterType((*CreateReply)(nil), "instance.CreateReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Instance service

type InstanceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error)
}

type instanceClient struct {
	cc *grpc.ClientConn
}

func NewInstanceClient(cc *grpc.ClientConn) InstanceClient {
	return &instanceClient{cc}
}

func (c *instanceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := grpc.Invoke(ctx, "/instance.Instance/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Instance service

type InstanceServer interface {
	Create(context.Context, *CreateRequest) (*CreateReply, error)
}

func RegisterInstanceServer(s *grpc.Server, srv InstanceServer) {
	s.RegisterService(&_Instance_serviceDesc, srv)
}

func _Instance_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instance.Instance/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Instance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "instance.Instance",
	HandlerType: (*InstanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Instance_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1.proto",
}

func init() { proto.RegisterFile("v1.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x33, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0xcc, 0x2b, 0x2e, 0x49, 0xcc, 0x4b, 0x4e, 0x55, 0x12,
	0xe5, 0xe2, 0x75, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xf1,
	0x62, 0xe1, 0x60, 0x14, 0x60, 0x52, 0xd2, 0xe3, 0xe2, 0x86, 0x09, 0x17, 0xe4, 0x54, 0x0a, 0xc9,
	0x73, 0x71, 0xc3, 0x74, 0xc4, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x71, 0xc1,
	0x84, 0x3c, 0x53, 0x8c, 0xdc, 0xb8, 0x38, 0x3c, 0xa1, 0x3c, 0x21, 0x2b, 0x2e, 0x36, 0x88, 0x5e,
	0x21, 0x71, 0x3d, 0x98, 0x12, 0x3d, 0x14, 0x4b, 0xa4, 0x44, 0x31, 0x25, 0x0a, 0x72, 0x2a, 0x95,
	0x18, 0x92, 0xd8, 0xc0, 0xee, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x5b, 0x77, 0xb2,
	0xab, 0x00, 0x00, 0x00,
}