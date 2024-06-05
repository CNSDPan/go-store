// Code generated by goctl. DO NOT EDIT.
// Source: socket.proto

package socketClient

import (
	"context"

	"store/rpc/socket/pb/socket"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	EventNoraml        = socket.EventNoraml
	ReqBroadcastNormal = socket.ReqBroadcastNormal
	ReqPing            = socket.ReqPing
	ResPong            = socket.ResPong
	ResSuccess         = socket.ResSuccess

	Socket interface {
		Ping(ctx context.Context, in *ReqPing, opts ...grpc.CallOption) (*ResPong, error)
		Broadcast(ctx context.Context, in *ReqBroadcastNormal, opts ...grpc.CallOption) (*ResSuccess, error)
	}

	defaultSocket struct {
		cli zrpc.Client
	}
)

func NewSocket(cli zrpc.Client) Socket {
	return &defaultSocket{
		cli: cli,
	}
}

func (m *defaultSocket) Ping(ctx context.Context, in *ReqPing, opts ...grpc.CallOption) (*ResPong, error) {
	client := socket.NewSocketClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultSocket) Broadcast(ctx context.Context, in *ReqBroadcastNormal, opts ...grpc.CallOption) (*ResSuccess, error) {
	client := socket.NewSocketClient(m.cli.Conn())
	return client.Broadcast(ctx, in, opts...)
}
