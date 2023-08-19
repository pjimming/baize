// Code generated by goctl. DO NOT EDIT.
package handler

import (
	local2 "github.com/pjimming/baize/core/internal/handler/local"
	"github.com/pjimming/baize/core/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/baize/v1/local/module",
				Handler: local2.GetModulePathHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/baize/v1/local/info",
				Handler: local2.GetProjectInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/baize/v1/local/graph",
				Handler: local2.GenerateGraphHandler(serverCtx),
			},
		},
	)
}
