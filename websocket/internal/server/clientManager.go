package server

import (
	"store/tools"
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
// @Auth：parker
// @Desc：连接事件
// @Date：2024-05-28 11:52:51
// @receiver：clientManager
// @param：autoToken
// @return：userId 用户ID
// @return：ClientId 连接唯一ID
func (clientManager *ClientManager) Connect(autoToken string) (userId int64, clientId int64) {
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
// @Auth：parker
// @Desc：断链事件
// @Date：2024-05-27 17:44:32
// @receiver：client
func (clientManager *ClientManager) DisConnect() {
	// 请求Grpc处理业务
	return
}

// NewBuckets
// @Auth：parker
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
// @Auth：parker
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
// @Auth：parker
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
