package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goctl-api/mall/service/user/api/internal/logic"
	"goctl-api/mall/service/user/api/internal/svc"
	"goctl-api/mall/service/user/api/internal/types"
)

func SignupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignupRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSignupLogic(r.Context(), svcCtx)
		resp, err := l.Signup(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
