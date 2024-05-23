package svc

import (
	model "goctl-api/mall/service/order"
	"goctl-api/mall/service/order/api/internal/config"
	"goctl-api/mall/service/order/api/internal/interceptor"
	"goctl-api/mall/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserRPC    userclient.User //RPC客户端代码
	OrderModel model.OrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//初始化user服务的RPC客户端(这里可以放拦截器)
		UserRPC: userclient.NewUser(zrpc.MustNewClient(c.UserRPC,
			zrpc.WithUnaryClientInterceptor(interceptor.ShoneunaryInterceptor))),
	}
}
