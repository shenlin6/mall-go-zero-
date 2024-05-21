package logic

import (
	"context"
	"errors"

	"goctl-api/mall/service/user/api/internal/svc"
	"goctl-api/mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DetailHanderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailHanderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailHanderLogic {
	return &DetailHanderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DetailHander 传入userID，显示UserName和Gender
func (l *DetailHanderLogic) DetailHander(req *types.Detailrequest) (resp *types.Detailresponse, err error) {
	// todo: add your logic here and delete this line
	// 1.拿到用户传来的userID
	user, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, int64(req.UserID))
	if err != nil { // 1.数据库查询失败 2.没查到userID
		if err != sqlx.ErrNotFound {
			logx.Errorf("user_detail_UserModel.FindOneByUserId failed,err:%#v\n", err)

			return nil, errors.New("内部错误")
		}
		//没查到userID
		return &types.Detailresponse{
			Message: "用户ID不存在",
		}, nil
	}

	// 2.成功查询到了userID,根据userID展示对应用户的信息
	return &types.Detailresponse{
		Message:  "用户信息如下:",
		UserName: user.Username,
		Gender:   int(user.Gender),
	}, nil

}
