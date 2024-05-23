package logic

import (
	"context"
	"errors"

	"goctl-api/mall/service/user/rpc/internal/svc"
	"goctl-api/mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserRequest) (*user.GetUserResponse, error) {
	// todo: add your logic here and delete this line
	// 1.根据user_id查询数据库，并返回信息
	one, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, in.UserID)
	if errors.Is(err, sqlx.ErrNotFound) {
		logx.Errorf("FindOneByUserId failed,err:%v\n", err)

		return nil, errors.New("无效的userID")
	}
	if err != nil {
		logx.Errorf("FindOneByUserId failed,err:%v\n", err)
		return nil, errors.New("查询失败")
	}

	//返回响应信息
	return &user.GetUserResponse{
		Message:  "该用户的信息如下:",
		UserID:   one.UserId,
		Username: one.Username,
		Gender:   one.Gender,
	}, nil
}
