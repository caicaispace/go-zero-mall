package logic

import (
	"context"

	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"
	"mall/service/pay/rpc/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) PayCreateLogic {
	return PayCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayCreateLogic) PayCreate(req types.PayCreateRequest) (resp *types.PayCreateResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.PayRpc.Create(l.ctx, &pay.CreateRequest{
		Uid:    req.Uid,
		Oid:    req.Oid,
		Amount: req.Amount,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.PayCreateResponse{
		Id: res.Id,
	}
	return
}
