package server

import "github.com/zeromicro/go-zero/zrpc"

func InitGrpcSocket(c zrpc.RpcClientConf) {
	zrpc.MustNewClient(c)
}
