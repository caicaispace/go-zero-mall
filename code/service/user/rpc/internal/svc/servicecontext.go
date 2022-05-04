package svc

import (
	"mall/service/user/model"
	"mall/service/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel     model.UserModel
	UserVipModel  model.UserVipModel
	UserCartModel model.UserCartModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(conn, c.CacheRedis),
		UserVipModel:  model.NewUserVipModel(conn, c.CacheRedis),
		UserCartModel: model.NewUserCartModel(conn, c.CacheRedis),
	}
}
