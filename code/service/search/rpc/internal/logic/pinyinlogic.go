package logic

import (
	"context"
	"strings"

	"mall/service/search/rpc/internal/svc"
	"mall/service/search/rpc/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type PinyinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPinyinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PinyinLogic {
	return &PinyinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PinyinLogic) Pinyin(in *search.PinyinRequest) (*search.PinyinResponse, error) {
	// todo: add your logic here and delete this line
	var resp []*search.PinyinItem
	resp = append(resp, &search.PinyinItem{
		Title: strings.Split(in.Keyword, " ")[0],
		Url:   strings.Split(in.Keyword, " ")[1],
	})
	return &search.PinyinResponse{List: resp}, nil
}
