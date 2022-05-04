package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	OrderRpc   zrpc.RpcClientConf
	PayRpc     zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf
	SearchRpc  zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
}
