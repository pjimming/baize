package local

import (
	"context"
	"github.com/pjimming/baize/core/internal/helper"
	"github.com/pjimming/baize/core/internal/svc"
	"github.com/pjimming/baize/core/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProjectInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectInfoLogic {
	return &GetProjectInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProjectInfoLogic) GetProjectInfo(req *types.GetProjectInfoReq) (resp *types.GetProjectInfoResp, err error) {

	moduleName, err := helper.GetModuleName(req.ModulePath)
	if err != nil {
		return nil, err
	}

	packageList, err := helper.GetAllPackages(req.Dir)
	if err != nil {
		return nil, err
	}

	resp = &types.GetProjectInfoResp{
		ModuleName:   moduleName,
		PackageList:  packageList,
		PackageCount: uint(len(packageList)),
	}

	return
}
