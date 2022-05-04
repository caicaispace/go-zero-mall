package logic

import (
	"context"

	"mall/service/vip/rpc/internal/svc"
	"mall/service/vip/rpc/vip"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *vip.DetailRequest) (*vip.DetailResponse, error) {
	// todo: add your logic here and delete this line

	return &vip.DetailResponse{}, nil
}
