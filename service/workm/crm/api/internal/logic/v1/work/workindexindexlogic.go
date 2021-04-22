package logic

import (
	"context"

	"dark-work/service/workm/crm/api/internal/svc"
	"dark-work/service/workm/crm/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type WorkIndexIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWorkIndexIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) WorkIndexIndexLogic {
	return WorkIndexIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WorkIndexIndexLogic) WorkIndexIndex() (*types.WorkIndexIndexResp, error) {
	return &types.WorkIndexIndexResp{}, nil
}
