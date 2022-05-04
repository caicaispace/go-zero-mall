package logic

import (
	"context"

	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"
	"mall/service/pay/rpc/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) PayDetailLogic {
	return PayDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayDetailLogic) PayDetail(req types.PayDetailRequest) (resp *types.PayDetailResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.PayRpc.Detail(l.ctx, &pay.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.PayDetailResponse{
		Id:     req.Id,
		Uid:    res.Uid,
		Oid:    res.Oid,
		Amount: res.Amount,
		Source: res.Source,
		Status: res.Status,
	}

	return
}
