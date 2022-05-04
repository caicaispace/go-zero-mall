package logic

import (
	"context"
	"encoding/json"
	"log"

	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"
	"mall/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCartListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserCartListLogic {
	return UserCartListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCartListLogic) UserCartList() (resp []*types.UserCartResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	uid = 20266849
	log.Println("uid", uid)
	res, err := l.svcCtx.UserRpc.UserCartList(l.ctx, &userclient.UserCartListRequest{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}

	for _, item := range res.Data {
		resp = append(resp, &types.UserCartResponse{
			Id:          item.Id,
			UserId:      item.UserId,
			Username:    item.Username,
			SourceId:    item.SourceId,
			SourceType:  item.SourceType,
			LicenseType: item.LicenseType,
			VideoRate:   item.VideoRate,
			SourceNum:   item.SourceNum,
			CreatedTime: item.CreatedTime,
			UpdatedTime: item.UpdatedTime,
		})
	}
	return
}
