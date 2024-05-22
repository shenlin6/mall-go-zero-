package main

import (
	"flag"
	"fmt"

	"goctl-api/mall/service/user/api/internal/config"
	"goctl-api/mall/service/user/api/internal/handler"
	"goctl-api/mall/service/user/api/internal/middleware"
	"goctl-api/mall/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	fmt.Printf("-->conf:%#v\n", c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//应用全局中间件
	server.Use(middleware.CopyResp)
	//server.Use(middleware.IsOpen(true))


	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
