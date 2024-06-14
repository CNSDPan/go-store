package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"net/http"
	"store/websocket/internal/config"
	"store/websocket/internal/handler"
	"store/websocket/internal/server"
	"store/websocket/internal/svc"
	"store/yaml"
	"strconv"
)

var configFile = flag.String("f", "etc/websocket-api.yaml", "the config file")

func main() {
	fmt.Println(fmt.Sprintln("Start websocket server ...\n "))
	var (
		c      config.Config
		nodeId int64
		node   *snowflake.Node
		err    error
		pong   string
		l      = logx.WithContext(context.Background())
	)
	flag.Parse()
	conf.MustLoad(*configFile, &c)
	s := rest.MustNewServer(c.RestConf)
	defer s.Stop()
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(s, ctx)
	// 服务uuid节点池
	if nodeId, err = strconv.ParseInt(c.ServiceId, 10, 64); err != nil {
		panic("服务-websocket serverId server string to int64 fail:" + err.Error())
	}
	if node, err = snowflake.NewNode(nodeId); err != nil {
		panic("服务-websocket start server newNode func fail:" + err.Error())
	}
	// 初始化redis
	if err = server.InitAloneRedis(); err != nil {
		panic("服务-websocket InitAloneRedis fail:" + err.Error())
	}
	// 初始grpc
	wbc := yaml.WebSocketConf
	pong, err = server.InitGrpcSocket(zrpc.RpcClientConf{
		Etcd: wbc.Etcd,
	})
	if err != nil {
		panic("服务-websocket InitGrpcSocket fail:" + err.Error())
	}
	fmt.Println(fmt.Sprintf("服务-websocket InitGrpcSocket ok pong:%s", pong))

	// 初始化连接池
	buckets := server.NewBuckets(c.BucketNumber)
	// 初始化websocket服务
	webServer := server.NewServer()
	webServer.Buckets = buckets
	webServer.BucketIdx = uint32(len(buckets))
	webServer.ServerId = c.ServiceId
	webServer.Log = l
	webServer.Node = node
	// 启动websocket服务
	//webServer.StartWebsocket()
	connect := server.NewConnect(webServer.ServerId, webServer.Node)
	s.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/ws",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			connect.Run(w, r, webServer)
		},
	})

	// 初始化订阅redis
	var subscribeServer *server.Subscribe
	subscribeServer, err = server.NewSubscribe()
	if err != nil {
		panic("服务-websocket NewSubscribe fail:" + err.Error())
	}
	subscribeServer.Log = l
	subscribeServer.SubReceive()

	fmt.Println(fmt.Sprintf("Starting websocket server at %s:%d...\n", c.Host, c.Port))
	s.Start()
}
