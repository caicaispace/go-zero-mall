package logic

import (
	"context"

	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VipUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVipUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) VipUpdateLogic {
	return VipUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VipUpdateLogic) VipUpdate(req types.VipUpdateRequest) (resp *types.VipUpdateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
