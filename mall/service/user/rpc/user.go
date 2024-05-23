package main

import (
	"context"
	"flag"
	"fmt"

	"goctl-api/mall/service/user/rpc/internal/config"
	"goctl-api/mall/service/user/rpc/internal/server"
	"goctl-api/mall/service/user/rpc/internal/svc"
	"goctl-api/mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 注册服务端拦截器(勿忘！)
	s.AddUnaryInterceptors(shoneIntercepterfunc)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

// shoneIntercepterfunc 自定义拦截器，打印传入的metadata
func shoneIntercepterfunc(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// 调用之前
	fmt.Println("服务器拦截器 in")
	// 拦截器逻辑（写在调用前或者调用后或者两边，看实际需求）
	// 取出元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	fmt.Printf("metadata:%#v\n", md)
	// 根据metadata中的数据进行一些校验处理
	if md["token"][0] != "mall-order-shone" {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	m, err := handler(ctx, req) //实际的RPC方法调用
	// 调用之后
	fmt.Println("服务器拦截器 out")
	return m, err

}
