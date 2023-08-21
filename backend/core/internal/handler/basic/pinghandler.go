package basic

import (
	"net/http"

	"github.com/pjimming/baize/core/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
		httpx.Ok(w)
	}
}
