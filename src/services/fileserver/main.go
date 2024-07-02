package main

import (
	"context"
	"log"

	"github.com/luminancetech/varso/src/common/api"
	common_config "github.com/luminancetech/varso/src/common/config"
	"github.com/luminancetech/varso/src/services/fileserver/ent"
	"github.com/luminancetech/varso/src/services/fileserver/routes/download"
	"github.com/luminancetech/varso/src/services/fileserver/routes/stitch"
	"github.com/luminancetech/varso/src/services/fileserver/routes/upload"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Starting fileserver in %s mode", common_config.AppEnv)
	defer logrus.Infof("Finished gracefully shutting down")

	err := ent.Connect(context.Background())
	if err != nil {
		logrus.WithError(err).Panicf("failed connecting to database")
	}

	router := api.NewRouter()

	api.RouteGET(router, api.Public, "/", api.SuccessEndpoint)

	api.RouteGET(router, api.Public, "/api/v1/fileserver/file/public/:filename", download.DownloadPublic)
	api.RoutePOST(router, api.User, "/api/v1/fileserver/file/public", upload.UploadPublic)

	api.RouteGET(router, api.User, "/api/v1/fileserver/file/default/:filename", download.Download)
	api.RoutePOST(router, api.User, "/api/v1/fileserver/file/default", upload.Upload)

	api.RoutePOST(router, api.Internal, "/api/v1/fileserver/stitch", stitch.Stitch)

	log.Fatal(router.ListenAndServe(":5002"))
}
