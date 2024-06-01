package main

import (
	"flag"
	"fmt"

	"store/rpc/socket/internal/config"
	"store/rpc/socket/internal/server"
	"store/rpc/socket/internal/svc"
	"store/rpc/socket/pb/socket"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/socket.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		socket.RegisterSocketServer(grpcServer, server.NewSocketServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	// 初始化redis
	if err := server.InitAloneRedis(); err != nil {
		panic("rpc.socket init alone redis panic:" + err.Error())
	}
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
