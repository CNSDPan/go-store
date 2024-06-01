package logic

import (
	"context"
	"store/rpc/socket/internal/svc"
	"store/rpc/socket/pb/socket"

	"github.com/zeromicro/go-zero/core/logx"
)

type BroadcastLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBroadcastLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BroadcastLogic {
	return &BroadcastLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BroadcastLogic) Broadcast(in *socket.ReqBroadcastNormal) (*socket.ResSuccess, error) {
	// todo: add your logic here and delete this line
	//server.AloneRedisClient.LPushX(common.Redis_Socket_Message_Normal_Key,in.)
	return &socket.ResSuccess{}, nil
}
