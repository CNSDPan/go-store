package server

import (
	"github.com/bwmarrin/snowflake"
	"k8s.io/apimachinery/pkg/util/rand"
	"net/http"
	"strconv"
)

type Bucket struct {
}

// StartWebsocket
// @Auth：parker
// @Desc：启动websocket服务
// @Date：2024-05-10 13:54:17
func StartWebsocket(serverId string) {
	var nodeId int64
	var err error
	// 设置节点ID，用于获取唯一ID
	var node *snowflake.Node
	nodeId, err = strconv.ParseInt(serverId+rand.String(10), 10, 64)
	if err != nil {
		panic("start websocket ParseInt fail:" + err.Error())
	}
	node, err = snowflake.NewNode(nodeId)
	if err != nil {
		panic("start websocket snowflake.NewNode fail:" + err.Error())
	}

	websocket := NewConnect(serverId, node)
	http.HandleFunc("/ws", websocket.Run)
}
