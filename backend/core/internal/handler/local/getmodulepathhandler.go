package local

import (
	"github.com/pjimming/baize/core/internal/logic/local"
	"github.com/pjimming/baize/core/internal/svc"
	"github.com/pjimming/baize/core/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetModulePathHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommonDirReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := local.NewGetModulePathLogic(r.Context(), svcCtx)
		resp, err := l.GetModulePath(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
