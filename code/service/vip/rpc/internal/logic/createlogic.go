package logic

import (
	"context"

	"mall/service/vip/rpc/internal/svc"
	"mall/service/vip/rpc/vip"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *vip.CreateRequest) (*vip.CreateResponse, error) {
	// todo: add your logic here and delete this line

	return &vip.CreateResponse{}, nil
}
