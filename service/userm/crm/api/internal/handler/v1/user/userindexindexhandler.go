package handler

import (
	"net/http"

	"dark-work/service/userm/crm/api/internal/logic/v1/user"
	"dark-work/service/userm/crm/api/internal/svc"
)

func UserIndexIndexHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserIndexIndexLogic(r.Context(), ctx)
		resp, err := l.UserIndexIndex()
		respz.Json(w, resp, err)
	}
}
