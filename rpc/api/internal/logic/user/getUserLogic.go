package userlogic

import (
	"context"

	"k/rpc/api/internal/svc"
	"k/rpc/api/pb/api"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(req *api.ReqUser) (*api.ResUser, error) {
	// todo: add your logic here and delete this line
	var user *api.ResUser
	if req.Iid == 1 {
		user = &api.ResUser{
			Iid:    1,
			Name:   "parker",
			CnName: "潘",
			Age:    26,
		}
	} else {
		user = &api.ResUser{
			Iid:    2,
			Name:   "xiha",
			CnName: "嘻哈",
			Age:    30,
		}
	}
	return user, nil
}
