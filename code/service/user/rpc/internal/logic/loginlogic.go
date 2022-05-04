package logic

import (
	"context"
	"fmt"

	"mall/common/cryptx"
	"mall/service/user/model"
	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// return &user.LoginResponse{
	// 	Id:     1,
	// 	Name:   "test",
	// 	Gender: 111,
	// 	Mobile: "13333333333",
	// }, nil
	// 查询用户是否存在
	fmt.Println(">>>>>>>>>>>>>>>")
	res, err := l.svcCtx.UserModel.FindOneByMobile(in.Mobile)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	fmt.Println("pass", in.Password)
	fmt.Println("pass", password)
	fmt.Println("pass", res.Password)
	if password != res.Password {
		return nil, status.Error(100, "密码错误")
	}
	resp = &user.LoginResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}
	return
}
