package logic

import (
	"context"

	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"
	"mall/service/order/rpc/orderclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderListLogic {
	return OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req types.OrderListRequest) (resp []*types.OrderListResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.OrderRpc.List(l.ctx, &orderclient.ListRequest{
		Uid: req.Uid,
	})
	if err != nil {
		return nil, err
	}

	for _, item := range res.Data {
		resp = append(resp, &types.OrderListResponse{
			Id:     item.Id,
			Uid:    item.Uid,
			Pid:    item.Pid,
			Amount: item.Amount,
			Status: item.Status,
		})
	}
	return
}
