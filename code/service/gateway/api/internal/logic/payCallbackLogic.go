package logic

import (
	"context"

	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"
	"mall/service/pay/rpc/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) PayCallbackLogic {
	return PayCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayCallbackLogic) PayCallback(req types.PayCallbackRequest) (resp *types.PayCallbackResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.PayRpc.Callback(l.ctx, &pay.CallbackRequest{
		Id:     req.Id,
		Uid:    req.Uid,
		Oid:    req.Oid,
		Amount: req.Amount,
		Source: req.Source,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return
}
