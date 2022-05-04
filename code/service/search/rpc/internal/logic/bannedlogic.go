package logic

import (
	"context"
	"strings"

	"mall/service/search/rpc/internal/svc"
	"mall/service/search/rpc/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type BannedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBannedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BannedLogic {
	return &BannedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BannedLogic) Banned(in *search.BannedRequest) (*search.BannedResponse, error) {
	// todo: add your logic here and delete this line

	var resp []*search.BannedItem
	resp = append(resp, &search.BannedItem{
		Title: strings.Split(in.Keyword, " ")[0],
		Url:   strings.Split(in.Keyword, " ")[1],
	})
	return &search.BannedResponse{List: resp}, nil
}
