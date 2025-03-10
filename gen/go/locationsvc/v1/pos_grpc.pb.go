// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: locationsvc/v1/pos.proto

package locationsvc

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
	LocationService_Position_FullMethodName       = "/locationsvc.v1.LocationService/Position"
	LocationService_StreamPosition_FullMethodName = "/locationsvc.v1.LocationService/StreamPosition"
)

// LocationServiceClient is the client API for LocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationServiceClient interface {
	Position(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (*PositionResponse, error)
	StreamPosition(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PositionResponse], error)
}

type locationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationServiceClient(cc grpc.ClientConnInterface) LocationServiceClient {
	return &locationServiceClient{cc}
}

func (c *locationServiceClient) Position(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (*PositionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PositionResponse)
	err := c.cc.Invoke(ctx, LocationService_Position_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) StreamPosition(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PositionResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &LocationService_ServiceDesc.Streams[0], LocationService_StreamPosition_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PositionRequest, PositionResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LocationService_StreamPositionClient = grpc.ServerStreamingClient[PositionResponse]

// LocationServiceServer is the server API for LocationService service.
// All implementations should embed UnimplementedLocationServiceServer
// for forward compatibility.
type LocationServiceServer interface {
	Position(context.Context, *PositionRequest) (*PositionResponse, error)
	StreamPosition(*PositionRequest, grpc.ServerStreamingServer[PositionResponse]) error
}

// UnimplementedLocationServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLocationServiceServer struct{}

func (UnimplementedLocationServiceServer) Position(context.Context, *PositionRequest) (*PositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Position not implemented")
}
func (UnimplementedLocationServiceServer) StreamPosition(*PositionRequest, grpc.ServerStreamingServer[PositionResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamPosition not implemented")
}
func (UnimplementedLocationServiceServer) testEmbeddedByValue() {}

// UnsafeLocationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationServiceServer will
// result in compilation errors.
type UnsafeLocationServiceServer interface {
	mustEmbedUnimplementedLocationServiceServer()
}

func RegisterLocationServiceServer(s grpc.ServiceRegistrar, srv LocationServiceServer) {
	// If the following call pancis, it indicates UnimplementedLocationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LocationService_ServiceDesc, srv)
}

func _LocationService_Position_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).Position(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationService_Position_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).Position(ctx, req.(*PositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_StreamPosition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LocationServiceServer).StreamPosition(m, &grpc.GenericServerStream[PositionRequest, PositionResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LocationService_StreamPositionServer = grpc.ServerStreamingServer[PositionResponse]

// LocationService_ServiceDesc is the grpc.ServiceDesc for LocationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "locationsvc.v1.LocationService",
	HandlerType: (*LocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Position",
			Handler:    _LocationService_Position_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamPosition",
			Handler:       _LocationService_StreamPosition_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "locationsvc/v1/pos.proto",
}
