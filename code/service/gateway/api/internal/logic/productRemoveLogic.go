package logic

import (
	"context"

	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"
	"mall/service/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductRemoveLogic {
	return ProductRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductRemoveLogic) ProductRemove(req types.ProductRemoveRequest) (resp *types.ProductRemoveResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.ProductRpc.Remove(l.ctx, &product.RemoveRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.ProductRemoveResponse{}
	return
}
