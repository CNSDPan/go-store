package server

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"store/rpc/socket/socketClient"
)

var GrpcSocketClient *socketClient.Socket

func InitGrpcSocket(c zrpc.RpcClientConf) (string, error) {
	var (
		res *socketClient.ResPong
		err error
	)
	conn := zrpc.MustNewClient(c)
	client := socketClient.NewSocket(conn)
	res, err = client.Ping(context.Background(), &socketClient.ReqPing{})
	if err == nil {
		GrpcSocketClient = &client
	}
	return res.Pong, err
}
