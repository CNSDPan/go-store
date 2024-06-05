// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: socket.proto

package socket

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

const (
	Socket_Ping_FullMethodName      = "/socket.Socket/Ping"
	Socket_Broadcast_FullMethodName = "/socket.Socket/Broadcast"
)

// SocketClient is the client API for Socket service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SocketClient interface {
	Ping(ctx context.Context, in *ReqPing, opts ...grpc.CallOption) (*ResPong, error)
	Broadcast(ctx context.Context, in *ReqBroadcastNormal, opts ...grpc.CallOption) (*ResSuccess, error)
}

type socketClient struct {
	cc grpc.ClientConnInterface
}

func NewSocketClient(cc grpc.ClientConnInterface) SocketClient {
	return &socketClient{cc}
}

func (c *socketClient) Ping(ctx context.Context, in *ReqPing, opts ...grpc.CallOption) (*ResPong, error) {
	out := new(ResPong)
	err := c.cc.Invoke(ctx, Socket_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socketClient) Broadcast(ctx context.Context, in *ReqBroadcastNormal, opts ...grpc.CallOption) (*ResSuccess, error) {
	out := new(ResSuccess)
	err := c.cc.Invoke(ctx, Socket_Broadcast_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SocketServer is the server API for Socket service.
// All implementations must embed UnimplementedSocketServer
// for forward compatibility
type SocketServer interface {
	Ping(context.Context, *ReqPing) (*ResPong, error)
	Broadcast(context.Context, *ReqBroadcastNormal) (*ResSuccess, error)
	mustEmbedUnimplementedSocketServer()
}

// UnimplementedSocketServer must be embedded to have forward compatible implementations.
type UnimplementedSocketServer struct {
}

func (UnimplementedSocketServer) Ping(context.Context, *ReqPing) (*ResPong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSocketServer) Broadcast(context.Context, *ReqBroadcastNormal) (*ResSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedSocketServer) mustEmbedUnimplementedSocketServer() {}

// UnsafeSocketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SocketServer will
// result in compilation errors.
type UnsafeSocketServer interface {
	mustEmbedUnimplementedSocketServer()
}

func RegisterSocketServer(s grpc.ServiceRegistrar, srv SocketServer) {
	s.RegisterService(&Socket_ServiceDesc, srv)
}

func _Socket_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqPing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocketServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Socket_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocketServer).Ping(ctx, req.(*ReqPing))
	}
	return interceptor(ctx, in, info, handler)
}

func _Socket_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqBroadcastNormal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocketServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Socket_Broadcast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocketServer).Broadcast(ctx, req.(*ReqBroadcastNormal))
	}
	return interceptor(ctx, in, info, handler)
}

// Socket_ServiceDesc is the grpc.ServiceDesc for Socket service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Socket_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "socket.Socket",
	HandlerType: (*SocketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Socket_Ping_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _Socket_Broadcast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "socket.proto",
}
