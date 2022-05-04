package logic

import (
	"context"

	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"
	"mall/service/search/rpc/searchclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchLogic {
	return SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req types.SearchRequest) (resp *types.SearchResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SearchRpc.Search(l.ctx, &searchclient.SearchRequest{
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
	resp = &types.SearchResponse{
		List: list,
	}
	return
}
