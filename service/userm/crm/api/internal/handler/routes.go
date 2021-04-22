// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	v1user "dark-work/service/userm/crm/api/internal/handler/v1/user"
	"dark-work/service/userm/crm/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/index/index",
				Handler: v1user.UserIndexIndexHandler(serverCtx),
			},
		},
	)
}
