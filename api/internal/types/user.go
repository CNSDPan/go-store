package types

type UserInfo struct {
	Iid    int64  `json:"iid,string"`
	Name   string `json:"name"`
	CnName string `json:"cnName"`
	Age    int32  `json:"age,string"`
}
