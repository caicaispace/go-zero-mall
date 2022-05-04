package logic

import (
	"context"
	"encoding/json"

	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"
	"mall/service/order/rpc/orderclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderCreateLogic {
	return OrderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderCreateLogic) OrderCreate(req types.OrderCreateRequest) (resp *types.OrderCreateResponse, err error) {
	// todo: add your logic here and delete this line
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	res, err := l.svcCtx.OrderRpc.Create(l.ctx, &orderclient.CreateRequest{
		Uid:    uid,
		Pid:    req.Pid,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.OrderCreateResponse{
		Id: res.Id,
	}
	return
}
