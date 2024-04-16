package types

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"store/common"
	"time"
)

func ResponseWithCode(res *Response, logger logx.Logger) {
	defer func() {
		res.ResponseTime = time.Now().Format("2006-01-02 15:04:05")
		if res.Code != common.RESPONSE_SUCCESS {
			res.Data = make(map[string]interface{})
		}
	}()

	defaultCodeMsg := common.ReturnOverCodeMessage()
	if codeMsg, err := common.ReturnCodeMessage(); err != nil {
		logger.Errorf("get common.ReturnCodeMessage fail:", err.Error())
		res.Code = common.RESPONSE_FAIL
		res.Message = defaultCodeMsg[res.Code]
	} else {
		msg := ""
		ok := false
		if msg, ok = codeMsg[res.Code]; !ok {
			code := common.RESPONSE_NOT_CODE
			msg = fmt.Sprintf("%s code:%s", codeMsg[code], res.Code)
			res.Code = code
		}
		res.Message = msg
	}
	return
}
