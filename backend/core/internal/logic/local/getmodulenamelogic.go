package local

import (
	"context"

	"github.com/pjimming/baize/core/internal/svc"
	"github.com/pjimming/baize/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetModuleNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetModuleNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetModuleNameLogic {
	return &GetModuleNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetModuleNameLogic) GetModuleName(req *types.CommonModulePathReq) (resp *types.GetModuleNameResp, err error) {
	// todo: add your logic here and delete this line

	return
}
