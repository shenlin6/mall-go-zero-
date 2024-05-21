package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"goctl-api/mall/service/user/api/internal/svc"
	"goctl-api/mall/service/user/api/internal/types"
	"goctl-api/mall/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var secret = []byte("阿巴阿巴阿巴")

type SignupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignupLogic {
	return &SignupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignupLogic) Signup(req *types.SignupRequest) (resp *types.SignupResponse, err error) {
	// 参数校验
	if req.RePassword != req.Password {
		return nil, errors.New("两次输入的密码不一致")
	}

	//注意:可以为唯一标识一个人的信息，不能落盘（使用info级别）
	logx.Debugv(req) //使用json/Marshall(req)
	logx.Debugf("req:%#v\n", req)

	// todo: add your logic here and delete this line
	//填充业务逻辑

	//// 0. 查询username是否已经被注册
	u, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.UserName)
	// 0.1 查询数据库失败
	if err != nil && err != sqlx.ErrNotFound { // err中包含 sqlx.ErrNotFound 因此要分开写
		logx.Errorf("user_signup_UserModel.FindOneByUsername failed,err:%v\n", err)
		return nil, errors.New("内部错误")
	}
	// 0.2 查到记录(该用户名已经被注册)
	if u != nil {
		return nil, errors.New("用户名已经存在")
	}
	// 0.3 没查到记录

	//// 1. 生成 userID(雪花算法)

	//// 2. 加密密码(加盐 md5)
	// var encryptPassword string
	// encryptPassword, err = hashPassword(req.Password)
	// if err != nil {
	// 	log.Printf("hashPassword failed,err:%v\n", err)
	// 	return nil, err
	// }

	//或者这样写
	h := md5.New()
	h.Write([]byte(req.Password)) //用户密码计算md5
	h.Write(secret)               //加盐
	encryptPassword := hex.EncodeToString(h.Sum(nil))

	//保存进MySQL数据库
	user := &model.User{
		UserId:   time.Now().Unix(), //后面再用雪花算法
		Username: req.UserName,
		Password: encryptPassword, //存入加密后的密码
		Gender:   int64(req.Gender),
	}
	_, err = l.svcCtx.UserModel.Insert(context.Background(), user)
	if err != nil {
		logx.Errorf(
			"user_signup_UserModel.Insert failed,err:%v\n",
			logx.Field("err", err),
		)

		fmt.Printf("UserModel.Insert failed,err:%v\n", err)
	}

	return &types.SignupResponse{Message: "success"}, nil
}

// hashPassword 对密码进行加盐和哈希处理
//func hashPassword(password string) (string, error) {
// 生成随机盐
// salt := make([]byte, 32)
// _, err := rand.Read(salt)
// if err != nil {
// 	return "", err
// }

// // 使用 scrypt 算法对密码和盐进行哈希处理
// hashedPassword, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)
// if err != nil {
// 	return "", err
// }

// // 将密码和盐拼接后的结果以十六进制编码返回
// return hex.EncodeToString(append(hashedPassword, salt...)), nil
//}
