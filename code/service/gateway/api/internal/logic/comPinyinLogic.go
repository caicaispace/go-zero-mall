package logic

import (
	"context"

	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"
	"mall/service/search/rpc/searchclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ComPinyinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewComPinyinLogic(ctx context.Context, svcCtx *svc.ServiceContext) ComPinyinLogic {
	return ComPinyinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ComPinyinLogic) ComPinyin(req types.ComPinyinRequest) (resp *types.ComPinyinResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SearchRpc.Pinyin(l.ctx, &searchclient.PinyinRequest{
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
	resp = &types.ComPinyinResponse{
		List: list,
	}
	return
}
