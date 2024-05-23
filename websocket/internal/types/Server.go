package types

import (
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

// Server
// @Auth：parker
// @Desc：websocket结构
// @Date：2024-05-23 10:43:44
type Server struct {
	Buckets   []*Bucket
	BucketIdx uint32
	ServerId  string
	Log       logx.Logger
	Node      *snowflake.Node
}

// Bucket
// @Auth：parker
// @Desc：连接池
// @Date：2024-05-23 10:43:39
type Bucket struct {
	CLock   sync.RWMutex
	Clients map[int64]*Client
	Stores  map[int64]*Store
	Rooms   map[int64]*Room
}
