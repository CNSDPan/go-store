// Package common
// Author：parker
// Date：  2024-04-07 18:16:50
package common

// Author：parker
// Desc：公共code
// Date：2024-04-07 18:16:50
const (
	RESPONSE_SUCCESS      = "200"
	RESPONSE_FAIL         = "500"
	RESPONSE_NOT_FOUND    = "404"
	RESPONSE_UNAUTHORIZED = "401"
	RESPONSE_NOT_CODE     = "1000"

	RESPONSE_REQUEST_TIME_FAIL = "10001"
	RESPONSE_TOKEN_FAIL        = "10002"

	RESPONSE_APPID_FAIL  = "10011"
	RESPONSE_SECRET_FAIL = "10012"
	RESPONSE_SIGN_FAIL   = "10013"
)

// Author：parker
// Desc：user模块code
// Date：2024-04-07 18:16:50
const (
	USER_INFO_FAIL = "10101"
	USER_ID_FAIL   = "10102"
)

// Author：parker
// Desc：rpc-socket模块code
// Date：
const (
	SOCKET_BROADCAST_LOGIN  = "20101"
	SOCKET_BROADCAST_LOGOUT = "20102"
	SOCKET_BROADCAST_NORMAL = "20103"
)
