package router

import (
	"context"

	"github.com/luminancetech/varso/src/common/api"
	"github.com/luminancetech/varso/src/services/auth/internal/router/routes/login"
)

func ListenAndServe(ctx context.Context) error {
	router := api.NewRouter()

	api.RouteGET(router, api.Public, "/api/v1/auth", api.SuccessEndpoint)
	api.RoutePOST(router, api.Public, "/api/v1/auth/login", login.Login)
	// api.RoutePOST(router, api.Public, "/api/v1/auth/register", register.Register) // Registration closed
	// api.RoutePOST(router, api.Public, "/api/v1/auth/internal_login", internal_login.InternalLogin) // TODO: Separate to different port

	return router.ListenAndServe(":5004")
}
