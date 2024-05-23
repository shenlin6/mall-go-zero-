package logic

import (
	"context"

	"goctl-api/mall/service/order/api/internal/svc"
	"goctl-api/mall/service/order/api/internal/types"
	"goctl-api/mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	// 1.通过订单ID找到对应订单记录
	// orderid, err := strconv.ParseUint(req.OrderID, 10, 64)
	// if err != nil {
	// 	logx.Errorf("strconv.ParseUint failed,err:%v\n", err)
	// 	return nil, err
	// }
	// order, _ := l.svcCtx.OrderModel.FindOne(l.ctx,orderid)

	// fmt.Printf("----------------------%d", order.UserId)

	// 2.通过订单记录上的用户ID找到对应的用户信息
	user, err := l.svcCtx.UserRPC.GetUser(l.ctx, &user.GetUserRequest{
		UserID: 1716270472,
	})
	if err != nil {
		logx.Errorf("UserRPC.GetUser failed,err:%v\n", err)
		return nil, err
	}

	// 3.拼接返回结果(这个接口的数据不是只由一个服务组成)
	return &types.SearchResponse{
		Message:  "用户信息如下：",
		OrderID:  "1716270472", //int(user.GetUserID()),
		Username: user.GetUsername(),
		Status:   100,
	}, nil
}
