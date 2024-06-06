package types

import (
	"fmt"
	"store/common"
	"time"
)

func ResponseWithCode(res *Response) {
	defer func() {
		res.ResponseTime = time.Now().Format("2006-01-02 15:04:05")
		if res.Code != common.RESPONSE_SUCCESS {
			res.Data = make(map[string]interface{})
		}
	}()

	defaultCodeMsg := common.ReturnOverCodeMessage()
	if res.Code == common.RESPONSE_SUCCESS {
		res.Message = defaultCodeMsg[res.Code]
		return
	}
	codeMsg := common.ReturnCodeMessage()
	msg := ""
	ok := false
	if msg, ok = codeMsg[res.Code]; !ok {
		code := common.RESPONSE_NOT_CODE
		msg = fmt.Sprintf("%s code:%s", codeMsg[code], res.Code)
		res.Code = code
	}
	res.Message = msg
	return
}
