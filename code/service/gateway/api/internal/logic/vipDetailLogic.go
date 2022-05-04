package logic

import (
	"context"

	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VipDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVipDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) VipDetailLogic {
	return VipDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VipDetailLogic) VipDetail(req types.VipDetailRequest) (resp *types.VipDetailResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
