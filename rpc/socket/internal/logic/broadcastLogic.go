package logic

import (
	"context"
	"store/common"
	"store/rpc/socket/internal/svc"
	"store/rpc/socket/pb/socket"

	"github.com/zeromicro/go-zero/core/logx"
)

var module = "socket-broadcast"

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

// Broadcast
// @Auth：
// @Desc：处理普通消息
// @Date：2024-06-04 17:56:59
// @receiver：l
// @param：in
// @return：*socket.ResSuccess
// @return：error
func (l *BroadcastLogic) Broadcast(in *socket.ReqBroadcastNormal) (*socket.ResSuccess, error) {
	var (
		err error
		res = &socket.ResSuccess{
			Module: "socket",
			Code:   common.RESPONSE_SUCCESS,
			Msg:    "",
		}
	)
	defer func() {
		res.Code, res.Msg, err = common.GetCodeMessage(res.Code)
		if err != nil {
			l.Logger.Errorf("%s fail:%s", module, err.Error())
		}
	}()

	if err = AloneRedisClient.LPushX(common.Redis_Socket_Message_Normal_Key, in.Event.Params).Err(); err != nil {
		res.Code = common.SOCKET_BROADCAST_NORMAL
	}
	return &socket.ResSuccess{}, nil
}
