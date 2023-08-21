package basic

import (
	"net/http"

	"github.com/pjimming/baize/common/httpresp"
	"github.com/pjimming/baize/core/internal/logic/basic"
	"github.com/pjimming/baize/core/internal/svc"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := basic.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping()

		httpresp.HTTP(w, r, resp, err)

	}
}
