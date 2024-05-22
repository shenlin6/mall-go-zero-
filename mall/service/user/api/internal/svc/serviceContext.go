package svc

import (
	"goctl-api/mall/service/user/api/internal/config"
	"goctl-api/mall/service/user/api/internal/middleware"
	"goctl-api/mall/service/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	Cost      rest.Middleware //自定义路由中间件(字段名要与.api中的声明一致对齐)
	UserModel model.UserModel //加入User表增删改查操作Model
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlxConn, c.CacheRedis),
		Cost:      middleware.NewCostMiddleware().Handle,
	}
}
