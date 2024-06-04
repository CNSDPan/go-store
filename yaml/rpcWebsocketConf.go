package yaml

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

type WebSocketCon struct {
	zrpc.RpcClientConf
	ServiceId string `json:",optional"`
}

var WebSocketConf *WebSocketCon

func init() {
	// 获取配置文件的路径
	realPath := getCurrentDir()
	websocketFilePath := realPath + "/rpc-websocket.yaml"
	websocketFile := flag.String("websocket-f", websocketFilePath, "the websocket config file")
	var c WebSocketCon
	conf.MustLoad(*websocketFile, &c)
	WebSocketConf = &c
}
