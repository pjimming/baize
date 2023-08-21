package local

import (
	"context"

	"github.com/pjimming/baize/common/errorx"
	"github.com/pjimming/baize/core/helper"
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

	projectPkgs, err := helper.GetLocalPackages(req.Dir)
	if err != nil {
		err = errorx.Error400(err.Error())
		return nil, err
	}

	OtherPkgs, err := helper.GetThirdPackages(req.Dir)
	if err != nil {
		err = errorx.Error400(err.Error())
		return nil, err
	}

	resp = &types.GetPackagesResp{
		OtherPkgList:    OtherPkgs,
		OtherPkgCount:   uint(len(OtherPkgs)),
		ProjectPkgList:  projectPkgs,
		ProjectPkgCount: uint(len(projectPkgs)),
	}

	return
}
