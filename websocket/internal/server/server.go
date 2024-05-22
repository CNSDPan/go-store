package server

import (
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"store/websocket/internal/types"
	"sync"
	"time"
)

type Server struct {
	Buckets   []*Bucket
	BucketIdx uint32
	ServerId  string
	Log       logx.Logger
	Node      *snowflake.Node
	Conf      ServerConf
}
type ServerConf struct {
	PingPeriod time.Duration
}

type Bucket struct {
	cLock   sync.RWMutex
	clients map[int64]*Client
	stores  map[int64]*types.Store
	rooms   map[int64]*types.Room
}

// NewServer
// @Auth：parker
// @Desc：初始化Server
// @Date：2024-05-14 14:32:44
// @return：*Server
func NewServer() *Server {
	return &Server{}
}

// StartWebsocket
// @Auth：parker
// @Desc：启动websocket服务
// @Date：2024-05-14 14:34:52
// @receiver：s
func (s *Server) StartWebsocket() {
	websocket := NewConnect(s.ServerId, s.Node)
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		websocket.Run(writer, request, s)
	})
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
			clients: make(map[int64]*Client),
			stores:  make(map[int64]*types.Store),
			rooms:   make(map[int64]*types.Room),
		}
	}
	return buckets
}

// writeMessage
// @Auth：parker
// @Desc：写消息的
// @Date：2024-05-22 15:47:30
// @receiver：s
func (s *Server) writeChannel(client *Client) {
	// ping前端的时隔
	ticker := time.NewTicker(s.Conf.PingPeriod)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case <-ticker.C:

		}
	}
}

// writeMessage
// @Auth：parker
// @Desc：读消息的
// @Date：2024-05-22 15:47:30
// @receiver：s
func (s *Server) readChannel(client *Client) {

}
