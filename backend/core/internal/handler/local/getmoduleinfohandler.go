package local

import (
	"net/http"

	"github.com/pjimming/baize/common/errorx"
	"github.com/pjimming/baize/common/httpresp"
	"github.com/pjimming/baize/core/internal/logic/local"
	"github.com/pjimming/baize/core/internal/svc"
	"github.com/pjimming/baize/core/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetModuleInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommonDirReq
		if err := httpx.Parse(r, &req); err != nil {
			httpresp.HttpError(w, r, errorx.NewCodeError(2, http.StatusBadRequest, err.Error()))
			return
		}

		l := local.NewGetModuleInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetModuleInfo(&req)

		httpresp.HTTP(w, r, resp, err)

	}
}
