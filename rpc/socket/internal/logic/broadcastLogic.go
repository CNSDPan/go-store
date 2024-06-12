package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"store/common"
	"store/rpc/socket/internal/svc"
	"store/rpc/socket/internal/types"
	"store/rpc/socket/pb/socket"

	"github.com/zeromicro/go-zero/core/logx"
)

var module = "socket-broadcast-logic"

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
func (l *BroadcastLogic) Broadcast(in *socket.ReqBroadcastNormal) (res *socket.ResSuccess, rpcErr error) {
	var (
		body     []byte
		err      error
		writeMsg types.WriteMsg
	)
	res = &socket.ResSuccess{
		Module: "socket",
		Code:   common.RESPONSE_SUCCESS,
		Msg:    "",
	}
	defer func() {
		res.Code, res.Msg = common.GetCodeMessage(res.Code)
		if err != nil {
			res.ErrMsg = err.Error()
			l.Logger.Errorf("%s Broadcast fail:%s", module, err.Error())
		}
	}()
	writeMsg = types.WriteMsg{
		Version:      int(in.Version),
		Operate:      int(in.Operate),
		Method:       in.Method,
		SendClientId: in.ToClientId,
		Extend:       in.Extend,
		Body:         []byte(in.Event.Params),
	}
	if body, err = jsonx.Marshal(writeMsg); err != nil {
		res.Code = common.SOCKET_BROADCAST_NORMAL
		goto result
	}
	// 发布消息，将消息都分发给订阅了的消费者
	if err = AloneRedisClient.Publish(common.PubSubSocketMessageNormalChannelKey, string(body)).Err(); err != nil {
		res.Code = common.SOCKET_BROADCAST_NORMAL
		goto result
	}
result:
	return res, rpcErr
}
