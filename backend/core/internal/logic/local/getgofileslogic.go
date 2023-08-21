package local

import (
	"context"

	"github.com/pjimming/baize/common/errorx"
	"github.com/pjimming/baize/core/helper"
	"github.com/pjimming/baize/core/internal/svc"
	"github.com/pjimming/baize/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoFilesLogic {
	return &GetGoFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoFilesLogic) GetGoFiles(req *types.CommonDirReq) (resp *types.GetGoFilesResp, err error) {

	goFiles, err := helper.GetAllGoFiles(req.Dir)
	if err != nil {
		err = errorx.Error400(err.Error())
		return nil, err
	}

	resp = &types.GetGoFilesResp{
		GoFileList:  goFiles,
		GoFileCount: uint(len(goFiles)),
	}

	return
}
