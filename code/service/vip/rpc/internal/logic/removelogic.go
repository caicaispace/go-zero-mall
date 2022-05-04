package logic

import (
	"context"

	"mall/service/vip/rpc/internal/svc"
	"mall/service/vip/rpc/vip"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *vip.RemoveRequest) (*vip.RemoveResponse, error) {
	// todo: add your logic here and delete this line

	return &vip.RemoveResponse{}, nil
}
