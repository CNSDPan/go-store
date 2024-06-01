// Code generated by goctl. DO NOT EDIT.
// Source: socket.proto

package server

import (
	"context"

	"store/rpc/socket/internal/logic"
	"store/rpc/socket/internal/svc"
	"store/rpc/socket/pb/socket"
)

type SocketServer struct {
	svcCtx *svc.ServiceContext
	socket.UnimplementedSocketServer
}

func NewSocketServer(svcCtx *svc.ServiceContext) *SocketServer {
	return &SocketServer{
		svcCtx: svcCtx,
	}
}

func (s *SocketServer) Broadcast(ctx context.Context, in *socket.ReqBroadcastNormal) (*socket.ResSuccess, error) {
	l := logic.NewBroadcastLogic(ctx, s.svcCtx)
	return l.Broadcast(in)
}
