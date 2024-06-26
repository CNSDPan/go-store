package main

import (
	"flag"
	"fmt"

	"store/rpc/api/internal/config"
	rolesServer "store/rpc/api/internal/server/roles"
	userServer "store/rpc/api/internal/server/user"
	"store/rpc/api/internal/svc"
	"store/rpc/api/pb/api"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("r-a-f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		api.RegisterUserServer(grpcServer, userServer.NewUserServer(ctx))
		api.RegisterRolesServer(grpcServer, rolesServer.NewRolesServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
