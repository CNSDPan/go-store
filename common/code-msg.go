package common

import (
	"errors"
	"github.com/imdario/mergo"
)

var codeMessage = map[string]string{
	RESPONSE_SUCCESS:           "success",
	RESPONSE_FAIL:              "服务器内部错误",
	RESPONSE_NOT_FOUND:         "请求资源不存在",
	RESPONSE_UNAUTHORIZED:      "缺少身份认证",
	RESPONSE_NOT_CODE:          "无定义code码",
	RESPONSE_REQUEST_TIME_FAIL: "缺少请求日期",
	RESPONSE_TOKEN_FAIL:        "无效token",
	RESPONSE_APPID_FAIL:        "无效APPID",
	RESPONSE_SECRET_FAIL:       "无效secret",
	RESPONSE_SIGN_FAIL:         "无效sign",
}

var codeMessageByUser = map[string]string{
	USER_INFO_FAIL: "用户信息不存在",
	USER_ID_FAIL:   "用户ID不存在|错误",
}

// ReturnOverCodeMessage
// @Auth：parker
// @Desc：全局的codeMsg
// @Date：2024-04-10 16:29:07
// @return：map[string]string
func ReturnOverCodeMessage() map[string]string {
	return codeMessage
}

// ReturnCodeMessage
// @Auth：parker
// @Desc：集合所有code对应得msg，并返回完整map
// @Date：2024-04-10 16:26:26
// @return：map[string]string
// @return：error
func ReturnCodeMessage() (map[string]string, error) {
	codeMsg := make(map[string]string)
	err := errors.New("")
	if err = mergo.Merge(&codeMsg, codeMessage); err != nil {
		goto res
	}
	if err = mergo.Merge(&codeMsg, codeMessageByUser); err != nil {
		goto res
	}
res:
	return codeMsg, err
}
