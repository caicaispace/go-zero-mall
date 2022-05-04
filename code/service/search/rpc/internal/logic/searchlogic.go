package logic

import (
	"context"
	"strings"

	"mall/service/search/rpc/internal/svc"
	"mall/service/search/rpc/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLogic) Search(in *search.SearchRequest) (*search.SearchResponse, error) {
	// todo: add your logic here and delete this line
	var resp []*search.SearchItem
	resp = append(resp, &search.SearchItem{
		Title: strings.Split(in.Keyword, " ")[0],
		Url:   strings.Split(in.Keyword, " ")[1],
	})
	return &search.SearchResponse{List: resp}, nil
}
