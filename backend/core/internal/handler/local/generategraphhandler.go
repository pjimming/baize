package local

import (
	"net/http"

	"github.com/pjimming/baize/core/internal/logic/local"
	"github.com/pjimming/baize/core/internal/svc"
	"github.com/pjimming/baize/core/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GenerateGraphHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommonDirReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := local.NewGenerateGraphLogic(r.Context(), svcCtx)
		resp, err := l.GenerateGraph(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
