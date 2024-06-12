package logic

import (
	"context"
	"time"

	"store/rpc/socket/internal/svc"
	"store/rpc/socket/pb/socket"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *socket.ReqPing) (*socket.ResPong, error) {
	// todo: add your logic here and delete this line

	return &socket.ResPong{Pong: "pong" + time.Now().Format("2006-01-02 15:04:05")}, nil
}
