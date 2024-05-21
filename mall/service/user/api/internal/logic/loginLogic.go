package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"

	"goctl-api/mall/service/user/api/internal/svc"
	"goctl-api/mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func PasswordMd5(password []byte) string {
	h := md5.New()
	h.Write([]byte(password)) //用户密码计算md5
	h.Write(secret)           //加盐
	encryptPassword := hex.EncodeToString(h.Sum(nil))
	return encryptPassword
}

// Login 实现登录功能
func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	// 1. 处理用户传来的请求，拿到用户名和密码

	// 2。看用户名和密码和数据库中是不是一致的
	// 用用户名查，再判断密码
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.UserName)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Errorf("user_login_UserModel.FindOneByUsername failed,err:%v\n", err)

		return &types.LoginResponse{
			Message: "用户名不存在",
		}, errors.New("内部错误")
	}

	if err == sqlx.ErrNotFound {
		return &types.LoginResponse{
			Message: "用户名不存在",
		}, nil
	}

	//需要先把用户登陆时输入的密码加盐再与数据库里的做对比
	if user.Password != PasswordMd5([]byte(req.Password)) {
		// 2.1 如果结果不一致--登陆失败
		return &types.LoginResponse{
			Message: "用户名或密码错误",
		}, nil
	}

	// 2.2 如果一致--登陆成功
	return &types.LoginResponse{
		Message: "登陆成功",
	}, nil

}
