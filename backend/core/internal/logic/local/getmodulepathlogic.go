package local

import (
	"context"
	"github.com/pjimming/baize/core/helper"

	"github.com/pjimming/baize/core/internal/svc"
	"github.com/pjimming/baize/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetModulePathLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetModulePathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetModulePathLogic {
	return &GetModulePathLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetModulePathLogic) GetModulePath(req *types.CommonDirReq) (resp *types.GetModulePathResp, err error) {

	modulePath, err := helper.GetModulePath(req.Dir)
	if err != nil {
		return nil, err
	}
	resp = &types.GetModulePathResp{ModulePath: modulePath}

	return
}
