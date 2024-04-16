package roleslogic

import (
	"context"

	"store/rpc/api/internal/svc"
	"store/rpc/api/pb/api"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolesListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRolesListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolesListLogic {
	return &GetRolesListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRolesListLogic) GetRolesList(in *api.ReqRolesReq) (*api.ReqRolesRes, error) {
	// todo: add your logic here and delete this line

	return &api.ReqRolesRes{}, nil
}
