package handler

import (
	"net/http"

	"mall/service/gateway/api/internal/logic"
	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewOrderCreateLogic(r.Context(), svcCtx)
		resp, err := l.OrderCreate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
