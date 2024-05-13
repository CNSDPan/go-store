package yaml

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"runtime"
	"strings"
)

func getCurrentDir() string {
	_, fileName, _, _ := runtime.Caller(1)
	aPath := strings.Split(fileName, "/")
	dir := strings.Join(aPath[0:len(aPath)-1], "/")
	return dir
}

// GetRpcApiConf
// @Auth：parker
// @Desc：获取rpc-api的配置文件信息
// @Date：2024-04-15 14:30:45
// @return：zrpc.RpcClientConf
func GetRpcApiConf() zrpc.RpcClientConf {
	// 获取配置文件的路径
	realPath := getCurrentDir()
	filePath := realPath + "/rpc-api.yaml"
	file := flag.String("r-api-f", filePath, "the rpc api config file")
	var rpcCon zrpc.RpcClientConf
	conf.MustLoad(*file, &rpcCon)
	return rpcCon
}

// GetRpcWebsocketConf
// @Auth：parker
// @Desc：获取rpc-websocket的配置文件信息
// @Date：2024-05-10 11:29:23
// @return：zrpc.RpcClientConf
func GetRpcWebsocketConf() zrpc.RpcClientConf {
	// 获取配置文件的路径
	realPath := getCurrentDir()
	filePath := realPath + "/rpc-websocket.yaml"
	file := flag.String("r-websocket-f", filePath, "the rpc websocket config file")
	var rpcCon zrpc.RpcClientConf
	conf.MustLoad(*file, &rpcCon)
	return rpcCon
}
