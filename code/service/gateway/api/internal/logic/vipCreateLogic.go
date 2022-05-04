package logic

import (
	"context"

	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VipCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVipCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) VipCreateLogic {
	return VipCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VipCreateLogic) VipCreate(req types.VipCreateRequest) (resp *types.VipCreateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
