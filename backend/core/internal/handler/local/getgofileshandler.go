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

func GetGoFilesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommonDirReq
		if err := httpx.Parse(r, &req); err != nil {
			httpresp.HttpError(w, r, errorx.NewCodeError(2, http.StatusBadRequest, err.Error()))
			return
		}

		l := local.NewGetGoFilesLogic(r.Context(), svcCtx)
		resp, err := l.GetGoFiles(&req)

		httpresp.HTTP(w, r, resp, err)

	}
}
