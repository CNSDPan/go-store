// Code generated by goctl. DO NOT EDIT.
package types

type ReqUserIid struct {
	Iid int64 `path:"iid"`
}

type ReqUserInfo struct {
	Iid    int64  `json:"iid,string"`
	Name   string `json:"name"`
	CnName string `json:"cnName"`
}

type Request struct {
}

type Response struct {
	Code         string      `json:"code"`
	Message      string      `json:"message"`
	ErrMessage   string 	`json:"errMessage"`
	ResponseTime string      `json:"responseTime"`
	Data         interface{} `json:"data"`
}
