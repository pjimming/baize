package local

import (
	"github.com/pjimming/baize/core/internal/logic/local"
	"github.com/pjimming/baize/core/internal/svc"
	"github.com/pjimming/baize/core/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetProjectInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProjectInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := local.NewGetProjectInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetProjectInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
