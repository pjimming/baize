package {{.PkgName}}

import (
	"net/http"

	"github.com/pjimming/baize/common/errorx"
	"github.com/pjimming/baize/common/httpresp"
	{{.ImportPackages}}

	{{if .HasRequest}}"github.com/zeromicro/go-zero/rest/httpx"{{end}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpresp.HttpError(w, r, errorx.NewCodeError(2, http.StatusBadRequest, err.Error()))
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})

		{{if .HasResp}}
		httpresp.HTTP(w, r, resp, err)
		{{else}}
		httpresp.HTTP(w, r, nil, err)
		{{end}}
	}
}
