package svc

import (
	"store/api/internal/config"
	"store/rpc/api/pb/api"
)

type ServiceContext struct {
	Config    config.Config
	RpcClient api.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
