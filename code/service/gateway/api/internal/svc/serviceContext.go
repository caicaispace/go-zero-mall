package svc

import (
	"mall/service/gateway/api/internal/config"
	"mall/service/order/rpc/orderclient"
	"mall/service/pay/rpc/payclient"
	"mall/service/product/rpc/productclient"
	"mall/service/search/rpc/searchclient"
	"mall/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc   orderclient.Order
	PayRpc     payclient.Pay
	ProductRpc productclient.Product
	SearchRpc  searchclient.Search
	UserRpc    userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// OrderRpc:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		// ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		// PayRpc: payclient.NewPay(zrpc.MustNewClient(c.PayRpc)),
		// SearchRpc: searchclient.NewSearch(zrpc.MustNewClient(c.SearchRpc)),
	}
}
