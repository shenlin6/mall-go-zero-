package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	//注意对应关系！
	rest.RestConf

	Mysql struct { //数据库配置
		DataSource string
	}

	CacheRedis cache.CacheConf
}
