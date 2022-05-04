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

type UserVipListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserVipListLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserVipListLogic {
	return UserVipListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserVipListLogic) UserVipList() (resp []*types.UserVipResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	uid = 20201023
	log.Println("uid", uid)
	res, err := l.svcCtx.UserRpc.UserVipList(l.ctx, &userclient.UserVipListRequest{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}

	for _, item := range res.Data {
		resp = append(resp, &types.UserVipResponse{
			Id:            item.Id,
			UserId:        item.UserId,
			VipId:         item.VipId,
			VipType:       item.VipType,
			VideoId:       item.VideoId,
			OrderId:       item.OrderId,
			LicenseId:     item.LicenseId,
			StartTime:     item.StartTime,
			EndTime:       item.EndTime,
			DayLimit:      item.DayLimit,
			TotalLimit:    item.TotalLimit,
			LastAdminUser: item.LastAdminUser,
			Remark:        item.Remark,
			CreatedTime:   item.CreatedTime,
		})
	}
	return
}
