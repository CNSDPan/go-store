// Code generated by goctl. DO NOT EDIT.
// Source: api.proto

package server

import (
	"context"

	"store/rpc/api/internal/logic/user"
	"store/rpc/api/internal/svc"
	"store/rpc/api/pb/api"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	api.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *api.ReqUser) (*api.ResUser, error) {
	l := userlogic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}