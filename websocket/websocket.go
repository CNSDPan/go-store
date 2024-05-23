package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"store/websocket/internal/config"
	"store/websocket/internal/handler"
	"store/websocket/internal/server"
	"store/websocket/internal/svc"
	"strconv"
)

var configFile = flag.String("f", "etc/websocket-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	s := rest.MustNewServer(c.RestConf)
	defer s.Stop()
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(s, ctx)

	// 服务uuid节点池
	var nodeId int64
	var node *snowflake.Node
	var err error
	if nodeId, err = strconv.ParseInt(c.ServiceId, 10, 64); err != nil {
		panic("服务-websocket serverId server string to int64 fail:" + err.Error())
	}
	if node, err = snowflake.NewNode(nodeId); err != nil {
		panic("服务-websocket start server newNode func fail:" + err.Error())
	}
	// 初始化连接池
	buckets := server.NewBuckets(c.BucketNumber)
	// 初始化websocket服务
	webServer := server.NewServer()
	webServer.Buckets = buckets
	webServer.BucketIdx = uint32(len(buckets))
	webServer.ServerId = c.ServiceId
	webServer.Log = logx.WithContext(context.Background())
	webServer.Node = node
	// 启动websocket服务
	webServer.StartWebsocket()

	fmt.Printf("Starting websocket server at %s:%d...\n", c.Host, c.Port)
	s.Start()
}
