package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"k/api/internal/config"
	"k/api/internal/handler"
	"k/api/internal/svc"
	"k/rpc/api/pb/api"
	"k/yaml"
)

var configFile = flag.String("c-f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 初始grpc服务
	rpcCon := yaml.GetRpcApiConf()
	grpcConn := zrpc.MustNewClient(rpcCon)
	grpcClient := api.NewUserClient(grpcConn.Conn())

	ctx := svc.NewServiceContext(c)
	ctx.RpcClient = grpcClient
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	server.Start()
}
