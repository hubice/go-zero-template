package handler

import (
	"dark-work/common/respz"
	"net/http"

	"dark-work/service/workm/crm/api/internal/logic/v1/work"
	"dark-work/service/workm/crm/api/internal/svc"
)

func WorkIndexIndexHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewWorkIndexIndexLogic(r.Context(), ctx)
		resp, err := l.WorkIndexIndex()
		respz.Json(w, resp, err)
	}
}
