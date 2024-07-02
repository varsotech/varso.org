package router

import (
	"context"

	"github.com/luminancetech/varso/src/common/api"
	"github.com/luminancetech/varso/src/services/app/internal/router/routes"
)

func ListenAndServe(ctx context.Context) error {
	router := api.NewRouter()

	// Register routes
	api.RouteGET(router, api.Public, "/api/v1/app", api.SuccessEndpoint)
	api.RouteGET(router, api.Public, "/api/v1/app/organization", routes.GetOrganizations)
	api.RouteGET(router, api.Public, "/api/v1/app/news/:type", routes.GetNews)
	api.RoutePOST(router, api.User, "/api/v1/app/news/item/blur/toggle", routes.BlurToggle)
	api.RoutePOST(router, api.User, "/api/v1/app/news/item/delete", routes.Delete)

	return router.ListenAndServe(":5001")
}
