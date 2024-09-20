// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: grpc_dialout_v3.proto

package telemetry

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GRPCDialoutV3_DialoutV3_FullMethodName = "/grpc_dialout_v3.gRPCDialoutV3/DialoutV3"
)

// GRPCDialoutV3Client is the client API for GRPCDialoutV3 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GRPCDialoutV3Client interface {
	DialoutV3(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[DialoutV3Args, DialoutV3Args], error)
}

type gRPCDialoutV3Client struct {
	cc grpc.ClientConnInterface
}

func NewGRPCDialoutV3Client(cc grpc.ClientConnInterface) GRPCDialoutV3Client {
	return &gRPCDialoutV3Client{cc}
}

func (c *gRPCDialoutV3Client) DialoutV3(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[DialoutV3Args, DialoutV3Args], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GRPCDialoutV3_ServiceDesc.Streams[0], GRPCDialoutV3_DialoutV3_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DialoutV3Args, DialoutV3Args]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GRPCDialoutV3_DialoutV3Client = grpc.BidiStreamingClient[DialoutV3Args, DialoutV3Args]

// GRPCDialoutV3Server is the server API for GRPCDialoutV3 service.
// All implementations must embed UnimplementedGRPCDialoutV3Server
// for forward compatibility.
type GRPCDialoutV3Server interface {
	DialoutV3(grpc.BidiStreamingServer[DialoutV3Args, DialoutV3Args]) error
	mustEmbedUnimplementedGRPCDialoutV3Server()
}

// UnimplementedGRPCDialoutV3Server must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGRPCDialoutV3Server struct{}

func (UnimplementedGRPCDialoutV3Server) DialoutV3(grpc.BidiStreamingServer[DialoutV3Args, DialoutV3Args]) error {
	return status.Errorf(codes.Unimplemented, "method DialoutV3 not implemented")
}
func (UnimplementedGRPCDialoutV3Server) mustEmbedUnimplementedGRPCDialoutV3Server() {}
func (UnimplementedGRPCDialoutV3Server) testEmbeddedByValue()                       {}

// UnsafeGRPCDialoutV3Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GRPCDialoutV3Server will
// result in compilation errors.
type UnsafeGRPCDialoutV3Server interface {
	mustEmbedUnimplementedGRPCDialoutV3Server()
}

func RegisterGRPCDialoutV3Server(s grpc.ServiceRegistrar, srv GRPCDialoutV3Server) {
	// If the following call pancis, it indicates UnimplementedGRPCDialoutV3Server was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GRPCDialoutV3_ServiceDesc, srv)
}

func _GRPCDialoutV3_DialoutV3_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GRPCDialoutV3Server).DialoutV3(&grpc.GenericServerStream[DialoutV3Args, DialoutV3Args]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GRPCDialoutV3_DialoutV3Server = grpc.BidiStreamingServer[DialoutV3Args, DialoutV3Args]

// GRPCDialoutV3_ServiceDesc is the grpc.ServiceDesc for GRPCDialoutV3 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GRPCDialoutV3_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_dialout_v3.gRPCDialoutV3",
	HandlerType: (*GRPCDialoutV3Server)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DialoutV3",
			Handler:       _GRPCDialoutV3_DialoutV3_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc_dialout_v3.proto",
}