package logic

import (
	"context"

	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserVipListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserVipListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserVipListLogic {
	return &UserVipListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserVipListLogic) UserVipList(in *user.UserVipListRequest) (resp *user.UserVipListResponse, err error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.UserVipModel.FindAllByUid(in.UserId)
	if err != nil {
		return nil, err
	}

	userVips := make([]*user.UserVip, 0)
	for _, item := range list {
		userVips = append(userVips, &user.UserVip{
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

	resp = &user.UserVipListResponse{
		Data: userVips,
	}
	return
}
