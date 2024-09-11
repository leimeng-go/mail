package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mail/internal/logic"
	"mail/internal/svc"
	"mail/internal/types"
)

func UserOrdersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserOrderRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserOrdersLogic(r.Context(), svcCtx)
		resp, err := l.UserOrders(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
