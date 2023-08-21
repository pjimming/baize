package local

import (
	"context"
	"github.com/pjimming/baize/common/errorx"
	"github.com/pjimming/baize/core/helper"

	"github.com/pjimming/baize/core/internal/svc"
	"github.com/pjimming/baize/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetModuleInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetModuleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetModuleInfoLogic {
	return &GetModuleInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetModuleInfoLogic) GetModuleInfo(req *types.CommonDirReq) (resp *types.GetModuleInfoResp, err error) {

	modulePath, err := helper.GetModulePath(req.Dir)
	if err != nil {
		err = errorx.Error400(err.Error())
		return nil, err
	}

	moduleName, err := helper.GetModuleName(modulePath)
	if err != nil {
		err = errorx.Error400(err.Error())
		return nil, err
	}

	resp = &types.GetModuleInfoResp{
		ModulePath: modulePath,
		ModuleName: moduleName,
	}

	return
}
