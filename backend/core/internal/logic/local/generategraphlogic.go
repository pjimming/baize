package local

import (
	"context"
	"github.com/pjimming/baize/core/internal/svc"
	"github.com/pjimming/baize/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateGraphLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateGraphLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateGraphLogic {
	return &GenerateGraphLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateGraphLogic) GenerateGraph(req *types.CommonDirReq) (resp *types.GenerateGraphResp, err error) {
	// todo: add your logic here and delete this line

	return
}