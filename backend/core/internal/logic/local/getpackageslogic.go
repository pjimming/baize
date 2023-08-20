package local

import (
	"context"
	"github.com/pjimming/baize/core/internal/helper"

	"github.com/pjimming/baize/core/internal/svc"
	"github.com/pjimming/baize/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPackagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPackagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPackagesLogic {
	return &GetPackagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPackagesLogic) GetPackages(req *types.CommonDirReq) (resp *types.GetPackagesResp, err error) {

	packageList, err := helper.GetLocalPackages(req.Dir)
	if err != nil {
		return nil, err
	}
	resp = &types.GetPackagesResp{
		PackageList:  packageList,
		PackageCount: uint(len(packageList)),
	}

	return
}
