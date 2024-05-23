package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

////context存值的经验！！！

// 为了避免其他人协作写项目的时候定义 "adminID"（any类型）导致冲突，需要重新自己的类型
type CtxKey string

const (
	CtxKeyAdminID CtxKey = "adminID"
)

// unaryInterceptor 客户端一元拦截器
func ShoneunaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("客户拦截器 in")
	//编写拦截器逻辑
	//RPC调用前
	adminID := ctx.Value(CtxKeyAdminID).(string)

	md := metadata.Pairs(
		//写死的数据
		"key1", "val1",
		"key1", "val1-2", // "key1"的值将会是 []string{"val1", "val1-2"}
		"requestID", "123456",
		"token", "mall-order-shone",
		//userID想要从外部获取
		"adminID", adminID,
	)
	ctx = metadata.NewOutgoingContext(ctx, md)           // 把metadata随RPC发送出去
	err := invoker(ctx, method, req, reply, cc, opts...) //实际的RPC调用
	fmt.Println("客户拦截器 out")
	return err
}
