package main

import (
	"context"

	common_config "github.com/luminancetech/varso/src/common/config"
	"github.com/luminancetech/varso/src/common/maingroup"
	"github.com/luminancetech/varso/src/services/app/internal/ent"
	"github.com/luminancetech/varso/src/services/app/internal/modules/news"
	"github.com/luminancetech/varso/src/services/app/internal/router"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Starting app service in %s mode", common_config.AppEnv)
	defer logrus.Infof("Finished gracefully shutting down")

	ctx := context.Background()

	err := ent.Connect(ctx)
	if err != nil {
		logrus.WithError(err).Panicf("failed connecting to database")
	}

	err = news.Initialize(ctx)
	if err != nil {
		logrus.WithError(err).Panicf("failed initializing organization module")
	}

	mainGroup := maingroup.New(ctx)

	mainGroup.Go(router.ListenAndServe)
	mainGroup.Go(news.SyncRSSFeeds)

	logrus.Panic(mainGroup.Wait())
}
