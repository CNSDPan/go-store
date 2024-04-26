package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"store/websocket/internal/logic"
	"store/websocket/internal/svc"
)

func TestConnHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewTestConnLogic(r.Context(), svcCtx)
		resp, err := l.TestConn()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
