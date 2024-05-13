package main

import (
	"flag"
	"fmt"
	"store/websocket/internal/config"
	"store/websocket/internal/handler"
	"store/websocket/internal/server"
	"store/websocket/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
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

	// 启动websocket服务
	server.StartWebsocket()

	fmt.Printf("Starting websocket server at %s:%d...\n", c.Host, c.Port)
	s.Start()
}
