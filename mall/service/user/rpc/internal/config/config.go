package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
    
	//MySQL
	Mysql struct { //数据库配置
		DataSource string
	}

	//Redis
	CacheRedis cache.CacheConf
}
