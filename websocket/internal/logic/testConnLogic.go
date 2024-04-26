package logic

import (
	"context"

	"store/websocket/internal/svc"
	"store/websocket/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestConnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestConnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestConnLogic {
	return &TestConnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestConnLogic) TestConn() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
