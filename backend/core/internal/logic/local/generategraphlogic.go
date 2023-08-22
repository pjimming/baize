package local

import (
	"context"
	"fmt"
	"strings"

	"github.com/pjimming/baize/common/errorx"
	"github.com/pjimming/baize/core/helper"
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

	// init response
	resp = &types.GenerateGraphResp{
		Nodes: make([]*types.GraphNode, 0),
		Edges: make([]*types.GraphEdge, 0),
	}

	// get module path
	modulePath, err := helper.GetModulePath(req.Dir)
	if err != nil {
		err = errorx.Error400(err.Error())
		return nil, err
	}

	// get nodes(local pkg)
	nodes, err := helper.GetLocalPackages(req.Dir)
	if err != nil {
		err = errorx.Error400(err.Error())
		return nil, err
	}

	nodeSet := make(map[string]struct{})
	for _, node := range nodes {
		nodeSet[node] = struct{}{}
		resp.Nodes = append(resp.Nodes, &types.GraphNode{
			Id:    node,
			Label: node,
		})
	}

	// get local go files
	goFiles, err := helper.GetAllGoFiles(req.Dir)
	if err != nil {
		err = errorx.Error400(err.Error())
		return nil, err
	}

	edgeSet := make(map[string]struct{})
	for _, goFile := range goFiles {
		pkgName, err := helper.GetFullPackageName(modulePath, goFile.Name)
		if err != nil {
			err = errorx.Error400(err.Error())
			return nil, err
		}

		imports, err := helper.GetFileImports(goFile.Name)
		if err != nil {
			err = errorx.Error400(err.Error())
			return nil, err
		}

		for _, pkgImport := range imports {
			if !checkSetExist(nodeSet, pkgImport) {
				continue
			}

			// unique edge
			edgeSet[fmt.Sprintf("%s$%s", pkgName, pkgImport)] = struct{}{}
		}
	}

	// add edge
	for edgeStr := range edgeSet {
		edge := strings.Split(edgeStr, "$")
		resp.Edges = append(resp.Edges, &types.GraphEdge{
			Source: edge[0],
			Target: edge[1],
		})
	}

	return
}

func checkSetExist(set map[string]struct{}, key string) bool {
	_, ok := set[key]
	return ok
}
