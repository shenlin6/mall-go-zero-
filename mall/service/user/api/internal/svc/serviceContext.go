package svc

import (
	"goctl-api/mall/service/user/api/internal/config"
	"goctl-api/mall/service/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UserModel //加入User表增删改查操作Model
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn:=sqlx.NewMysql(c.Mysql.DataSource)
    
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(sqlxConn,c.CacheRedis),
	}
}
