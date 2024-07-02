package config

import (
	_ "embed"

	"github.com/luminancetech/varso/src/common/config"
)

var (
	AnalyticsDatabaseName = config.RequireEnv("ANALYTICS_DATABASE_NAME")
)
