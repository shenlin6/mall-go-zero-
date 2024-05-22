package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	//注意对应关系！
	rest.RestConf

	Auth struct {// JWT 认证需要的密钥和过期时间配置
        AccessSecret string
        AccessExpire int64
    }

	Mysql struct { //数据库配置
		DataSource string
	}

	CacheRedis cache.CacheConf
}
