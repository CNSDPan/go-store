package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"k/api/internal/logic"
	"k/api/internal/svc"
	"k/api/internal/types"
)

func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqUserIid
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			types.ResponseWithCode(resp, l.Logger)
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
