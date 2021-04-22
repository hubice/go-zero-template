package logic

import (
	"context"

	"dark-work/service/userm/crm/api/internal/svc"
	"dark-work/service/userm/crm/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserIndexIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserIndexIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserIndexIndexLogic {
	return UserIndexIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserIndexIndexLogic) UserIndexIndex() (*types.UserIndexIndexResp, error) {
	return &types.UserIndexIndexResp{
		Msg: "user ...",
	}, nil
}
