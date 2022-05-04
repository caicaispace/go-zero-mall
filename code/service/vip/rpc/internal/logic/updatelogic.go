package logic

import (
	"context"

	"mall/service/vip/rpc/internal/svc"
	"mall/service/vip/rpc/vip"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *vip.UpdateRequest) (*vip.UpdateResponse, error) {
	// todo: add your logic here and delete this line

	return &vip.UpdateResponse{}, nil
}
