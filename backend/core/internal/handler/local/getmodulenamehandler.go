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

func GetModuleNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommonModulePathReq
		if err := httpx.Parse(r, &req); err != nil {
			httpresp.HttpError(w, r,
				errorx.NewCodeError(http.StatusBadRequest, 2, err.Error()))
			return
		}

		l := local.NewGetModuleNameLogic(r.Context(), svcCtx)
		resp, err := l.GetModuleName(&req)

		httpresp.HTTP(w, r, resp, err)
	}
}
