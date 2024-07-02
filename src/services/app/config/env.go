package config

import (
	_ "embed"

	"github.com/luminancetech/varso/src/common/config"
)

var (
	AppDatabaseName = config.RequireEnv("APP_DATABASE_NAME")
)
