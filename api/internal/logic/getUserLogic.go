package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"k/api/internal/svc"
	"k/api/internal/types"
	"k/common"
	"k/rpc/api/pb/api"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.ReqUserIid) (*types.Response, error) {
	l.Logger.Infof("req :%v", req)
	user := &types.UserInfo{}
	userInfo := &api.ResUser{}
	err := errors.New("")
	resp := &types.Response{
		Code:    common.RESPONSE_FAIL,
		Message: "",
	}
	defer func() {
		if err != nil {
			resp.Data = &types.UserInfo{}
		} else {
			resp.Code = common.RESPONSE_SUCCESS
			resp.Data = user
		}
	}()
	if userInfo, err = l.svcCtx.RpcClient.GetUser(l.ctx, &api.ReqUser{Iid: req.Iid}); err != nil {
		l.Logger.Infof("rpc client GetUserInfo fail:%s", err.Error())
		goto Res
	}
	user.Iid = userInfo.Iid
	user.Name = userInfo.Name
	user.CnName = userInfo.CnName
	user.Age = userInfo.Age
Res:
	return resp, err
}
