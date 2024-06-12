package server

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"store/common"
	"store/rpc/socket/pb/socket"
	"store/tools"
	"store/websocket/internal/types"
	"strconv"
)

type ClientM interface {
	Connect()
	DisConnect()
}

type ClientManager struct {
}

var userMap = userTest()

func userTest() map[string]map[string]interface{} {
	userMaps := make(map[string]map[string]interface{})
	userMaps["2gDGQwDxsrX0UG8yRbophdHxHqD"] = map[string]interface{}{
		"userId":   int64(1788408218839183360),
		"name":     "用户1",
		"fund":     strconv.FormatFloat(tools.OutExchange(int64(10000000)), 'f', 2, 64),
		"status":   "1",
		"clientId": int64(1788408218839183361),
	}
	userMaps["2gDGQugkyFF4MI10hK7WfT3W3Pe"] = map[string]interface{}{
		"userId":   int64(1788408218897903616),
		"name":     "用户2",
		"fund":     strconv.FormatFloat(tools.OutExchange(int64(10000000)), 'f', 2, 64),
		"status":   "1",
		"clientId": int64(1788408218897903612),
	}
	userMaps["2gDGQvEugR6Y5riFp2kVLdc7J0O"] = map[string]interface{}{
		"userId":   int64(1788408218960818176),
		"name":     "用户3",
		"fund":     strconv.FormatFloat(tools.OutExchange(int64(10000000)), 'f', 2, 64),
		"status":   "1",
		"clientId": int64(1788408218960818173),
	}
	userMaps["2gDGQwhqJQczjkCikEvg3StOKSR"] = map[string]interface{}{
		"userId":   int64(1788408219027927040),
		"name":     "用户4",
		"fund":     strconv.FormatFloat(tools.OutExchange(int64(10000000)), 'f', 2, 64),
		"status":   "1",
		"clientId": int64(1788408219027927044),
	}
	return userMaps
}

// Connect
// @Auth：
// @Desc：连接事件
// @Date：2024-05-28 11:52:51
// @receiver：clientManager
// @param：autoToken
// @return：userId 用户ID
// @return：ClientId 连接唯一ID
func (clientM *ClientManager) Connect(autoToken string) (userId int64, clientId int64) {
	//请求grpc处理业务，校验用户是否正常的，正常则返回clientId和用户信息等

	// 例子
	var user = make(map[string]interface{})
	var ok = false
	if user, ok = userMap[autoToken]; !ok {
		userId = 0
		clientId = 0
	} else {
		userId, _ = user["userId"].(int64)
		clientId, _ = user["clientId"].(int64)
	}

	return userId, clientId
}

// DisConnect
// @Auth：
// @Desc：断链事件
// @Date：2024-05-27 17:44:32
// @receiver：client
func (clientM *ClientManager) DisConnect() {
	// 请求Grpc处理业务
	return
}

// NewBuckets
// @Auth：
// @Desc：初始化websocket连接池
// @Date：2024-05-14 14:11:27
// @param：bucketNumber 池子数量
// @return：[]*Bucket
func NewBuckets(bucketNumber uint) []*Bucket {
	buckets := make([]*Bucket, bucketNumber)
	for i := uint(0); i < bucketNumber; i++ {
		buckets[i] = &Bucket{
			Clients: make(map[int64]*Client),
			Rooms:   make(map[int64][]int64),
			Idx:     uint32(i),
		}
	}
	return buckets
}

// putBucket
// @Auth：
// @Desc：加入某个池子的
// @Date：2024-05-28 17:56:51
// @receiver：b
// @param：client
// @param：roomId
func (b *Bucket) putBucket(client *Client, roomId int64) {
	b.CLock.Lock()
	defer b.CLock.Unlock()
	b.Clients[client.ClientId] = client
	b.Rooms[roomId] = append(b.Rooms[roomId], client.ClientId)
	return
}

// DelBucket
// @Auth：
// @Desc：移出池子
// @Date：2024-05-29 10:36:35
// @receiver：b
// @param：client
func (b *Bucket) DelBucket(client *Client) {
	b.CLock.Lock()
	defer b.CLock.Unlock()
	delete(b.Rooms, client.RoomId)
	delete(b.Clients, client.ClientId)
	return
}

// MethodHandle
// @Auth：
// @Desc：事件处理
// @Date：2024-05-30 15:30:41
// @receiver：clientM
// @param：msg
func (clientM *ClientManager) MethodHandle(msg types.ReceiveMsg, l logx.Logger) (code string, message string, err error) {
	var (
		fromClientId int64
		toClientId   int64
		res          *socket.ResSuccess
	)
	defer func() {
		if err != nil {
			code = common.RESPONSE_FAIL
			message = common.ReturnCodeMessage()[code]
		} else {
			code = "200"
			message = ""
		}
	}()
	if fromClientId, err = strconv.ParseInt(msg.FromClientId, 10, 64); err != nil {
		goto EndHandle
	}
	if toClientId, err = strconv.ParseInt(msg.ToClientId, 10, 64); err != nil {
		goto EndHandle
	}
	switch msg.Method {
	case "Enter":
	case "Out":
	case "Normal":
		params, ok := msg.Event.Params.(string)
		if !ok {
			err = errors.New("msg.Event.Params interface to string not ok")
			goto EndHandle
		}
		res, err = GrpcSocketClient.Broadcast(context.Background(), &socket.ReqBroadcastNormal{
			Version:      int32(msg.Version),
			Operate:      int32(msg.Operate),
			Method:       msg.Method,
			Event:        &socket.EventNoraml{Params: params},
			RoomId:       msg.RoomId,
			FromClientId: fromClientId,
			ToClientId:   toClientId,
			Msg:          "",
			Extend:       msg.Extend,
			AutoToken:    msg.AuthToken,
		})
	case "Server":
	default:
		res.Code = common.RESPONSE_FAIL
		res.Msg = "无效操作"
	}
EndHandle:
	return code, message, err
}
