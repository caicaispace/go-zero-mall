package logic

import (
	"context"

	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCartListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCartListLogic {
	return &UserCartListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCartListLogic) UserCartList(in *user.UserCartListRequest) (*user.UserCartListResponse, error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.UserCartModel.FindAllByUidWithUserInfo(in.UserId)
	if err != nil {
		return nil, err
	}

	userCartList := make([]*user.UserCart, 0)
	for _, item := range list {
		userCartList = append(userCartList, &user.UserCart{
			Id:          item.Id,
			Username:    item.Username,
			UserId:      item.UserId,
			SourceId:    item.SourceId,
			SourceType:  item.SourceType,
			LicenseType: item.LicenseType,
			VideoRate:   item.VideoRate,
			SourceNum:   item.SourceNum,
			CreatedTime: item.CreatedTime,
			UpdatedTime: item.UpdatedTime,
		})
	}

	return &user.UserCartListResponse{
		Data: userCartList,
	}, nil
}
