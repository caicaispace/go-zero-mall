package logic

import (
	"context"

	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VipRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVipRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) VipRemoveLogic {
	return VipRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VipRemoveLogic) VipRemove(req types.VipRemoveRequest) (resp *types.VipRemoveResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
