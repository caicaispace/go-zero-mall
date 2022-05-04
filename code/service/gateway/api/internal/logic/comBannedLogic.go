package logic

import (
	"context"

	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"
	"mall/service/search/rpc/searchclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ComBannedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewComBannedLogic(ctx context.Context, svcCtx *svc.ServiceContext) ComBannedLogic {
	return ComBannedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ComBannedLogic) ComBanned(req types.ComBannedRequest) (resp *types.ComBannedResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SearchRpc.Banned(l.ctx, &searchclient.BannedRequest{
		Keyword: req.Keyword,
	})
	if err != nil {
		return nil, err
	}
	var list []map[string]string
	for _, v := range res.List {
		list = append(list, map[string]string{
			"title": v.Title,
			"url":   v.Url,
		})
	}
	resp = &types.ComBannedResponse{
		List: list,
	}
	return
}
