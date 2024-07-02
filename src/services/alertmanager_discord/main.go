package main

import (
	"log"

	"github.com/luminancetech/varso/src/common/api"
	common_config "github.com/luminancetech/varso/src/common/config"
	"github.com/luminancetech/varso/src/services/alertmanager_discord/internal/routes/alert"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Starting alertmanager_discord in %s mode", common_config.AppEnv)
	defer logrus.Infof("Finished gracefully shutting down")

	router := api.NewRouter()

	api.RouteGET(router, api.Public, "/", api.SuccessEndpoint)
	api.RoutePOST(router, api.Public, "/api/v1/alertmanager_discord/alert", alert.Alert)

	log.Fatal(router.ListenAndServe(":9094"))
}
