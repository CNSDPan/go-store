package test

import (
	"errors"
	"fmt"
	"gopkg.in/square/go-jose.v2/json"
	"k/api/internal/types"
	"k/rpc/api/pb/api"
	"testing"
)

func TestGetUser(t *testing.T) {
	u := types.UserInfo{}
	uByte := []byte{}
	err := errors.New("")
	user := &api.ResUser{
		Iid:    1,
		Name:   "parker",
		CnName: "潘",
		Age:    26,
	}
	if uByte, err = json.Marshal(user); err != nil {
		fmt.Printf("报错 json.Marshal :%v %s", err, "\n")
	}

	if err = json.Unmarshal(uByte, &u); err != nil {
		fmt.Printf("报错 json.Unmarshal :%v %s", err, "\n")
	}
	u.Iid = user.Iid
	u.Age = user.Age
	fmt.Printf("打印 u：%v %s", u, "\n")
}
