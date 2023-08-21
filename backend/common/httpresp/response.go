package httpresp

import (
	"net/http"

	"github.com/pjimming/baize/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	codeServerError     = 1
	internalServerError = "内部错误，请查看日志信息或联系开发者"
)

type (
	errResp struct {
		Code int    `json:"code"`
		Desc string `json:"desc,omitempty"`
	}

	successResp struct {
		Success bool        `json:"success"`
		Result  interface{} `json:"result"`
	}
)

func HTTP(w http.ResponseWriter, r *http.Request, resp interface{}, err error) {
	if err != nil {
		HttpError(w, r, err)
	} else {
		HttpOkJson(w, r, resp)
	}
}

func HttpError(w http.ResponseWriter, r *http.Request, err error) {
	codeErr, ok := errorx.FromError(err)
	if ok {
		httpx.WriteJson(w, codeErr.Status(), errResp{
			Code: codeErr.Code(),
			Desc: codeErr.Error(),
		})
	} else {
		httpx.WriteJson(w, http.StatusInternalServerError, errResp{
			Code: codeServerError,
			Desc: internalServerError,
		})
		logx.WithContext(r.Context()).Error(err)
	}
}

func HttpOkJson(w http.ResponseWriter, r *http.Request, resp interface{}) {
	httpx.OkJsonCtx(r.Context(), w, successResp{Success: true, Result: resp})
}
