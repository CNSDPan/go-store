package common

import "fmt"

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

var codeMessageBySocket = map[string]string{
	SOCKET_BROADCAST_LOGIN:  "socket连接错误",
	SOCKET_BROADCAST_LOGOUT: "socket关闭错误",
	SOCKET_BROADCAST_NORMAL: "socket广播消息错误",
}

// ReturnOverCodeMessage
// @Auth：
// @Desc：全局的codeMsg
// @Date：2024-04-10 16:29:07
// @return：map[string]string
func ReturnOverCodeMessage() map[string]string {
	return codeMessage
}

// ReturnCodeMessage
// @Auth：
// @Desc：集合所有code对应得msg，并返回完整map
// @Date：2024-04-10 16:26:26
// @return：map[string]string
// @return：error
func ReturnCodeMessage() map[string]string {
	codeMsg := make(map[string]string)
	mergeMap := func(codeMsg map[string]string, m map[string]string) map[string]string {
		for key, value := range m {
			codeMsg[key] = value
		}
		return codeMsg
	}
	codeMsg = mergeMap(codeMsg, codeMessage)
	codeMsg = mergeMap(codeMsg, codeMessageByUser)
	codeMsg = mergeMap(codeMsg, codeMessageBySocket)
	return codeMsg
}

// GetCodeMessage
// @Auth：
// @Desc：获取codeMsg
// @Date：2024-06-04 17:36:49
// @param：code
// @return：string code
// @return：string msg
// @return：error
func GetCodeMessage(code string) (string, string) {
	var (
		codeMsg map[string]string
		message string
		ok      bool
	)
	defaultCodeMsg := ReturnOverCodeMessage()
	if code == RESPONSE_SUCCESS {
		return code, defaultCodeMsg[code]
	}
	codeMsg = ReturnCodeMessage()
	if message, ok = codeMsg[code]; !ok {
		message = defaultCodeMsg[RESPONSE_NOT_CODE] + fmt.Sprintf("; code: %s", code)
		code = RESPONSE_NOT_CODE
	}
	return code, message
}
