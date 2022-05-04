package logic

import (
	"context"

	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"
	"mall/service/order/rpc/orderclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderRemoveLogic {
	return OrderRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderRemoveLogic) OrderRemove(req types.OrderRemoveRequest) (resp *types.OrderRemoveResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.OrderRpc.Remove(l.ctx, &orderclient.RemoveRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return
}
