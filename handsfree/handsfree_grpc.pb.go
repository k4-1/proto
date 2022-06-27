// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: handsfree/handsfree.proto

package handsfree

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HandsfreeServiceClient is the client API for HandsfreeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HandsfreeServiceClient interface {
	GetToken(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (HandsfreeService_GetTokenClient, error)
}

type handsfreeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHandsfreeServiceClient(cc grpc.ClientConnInterface) HandsfreeServiceClient {
	return &handsfreeServiceClient{cc}
}

func (c *handsfreeServiceClient) GetToken(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (HandsfreeService_GetTokenClient, error) {
	stream, err := c.cc.NewStream(ctx, &HandsfreeService_ServiceDesc.Streams[0], "/infinimesh.handsfree.HandsfreeService/GetToken", opts...)
	if err != nil {
		return nil, err
	}
	x := &handsfreeServiceGetTokenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HandsfreeService_GetTokenClient interface {
	Recv() (*ControlPacket, error)
	grpc.ClientStream
}

type handsfreeServiceGetTokenClient struct {
	grpc.ClientStream
}

func (x *handsfreeServiceGetTokenClient) Recv() (*ControlPacket, error) {
	m := new(ControlPacket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HandsfreeServiceServer is the server API for HandsfreeService service.
// All implementations must embed UnimplementedHandsfreeServiceServer
// for forward compatibility
type HandsfreeServiceServer interface {
	GetToken(*ConnectionRequest, HandsfreeService_GetTokenServer) error
	mustEmbedUnimplementedHandsfreeServiceServer()
}

// UnimplementedHandsfreeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHandsfreeServiceServer struct {
}

func (UnimplementedHandsfreeServiceServer) GetToken(*ConnectionRequest, HandsfreeService_GetTokenServer) error {
	return status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedHandsfreeServiceServer) mustEmbedUnimplementedHandsfreeServiceServer() {}

// UnsafeHandsfreeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HandsfreeServiceServer will
// result in compilation errors.
type UnsafeHandsfreeServiceServer interface {
	mustEmbedUnimplementedHandsfreeServiceServer()
}

func RegisterHandsfreeServiceServer(s grpc.ServiceRegistrar, srv HandsfreeServiceServer) {
	s.RegisterService(&HandsfreeService_ServiceDesc, srv)
}

func _HandsfreeService_GetToken_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HandsfreeServiceServer).GetToken(m, &handsfreeServiceGetTokenServer{stream})
}

type HandsfreeService_GetTokenServer interface {
	Send(*ControlPacket) error
	grpc.ServerStream
}

type handsfreeServiceGetTokenServer struct {
	grpc.ServerStream
}

func (x *handsfreeServiceGetTokenServer) Send(m *ControlPacket) error {
	return x.ServerStream.SendMsg(m)
}

// HandsfreeService_ServiceDesc is the grpc.ServiceDesc for HandsfreeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HandsfreeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infinimesh.handsfree.HandsfreeService",
	HandlerType: (*HandsfreeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetToken",
			Handler:       _HandsfreeService_GetToken_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "handsfree/handsfree.proto",
}
