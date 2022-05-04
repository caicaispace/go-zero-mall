package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mall/service/gateway/api/internal/logic"
	"mall/service/gateway/api/internal/svc"
	"mall/service/gateway/api/internal/types"
)

func ComBannedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ComBannedRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewComBannedLogic(r.Context(), svcCtx)
		resp, err := l.ComBanned(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
