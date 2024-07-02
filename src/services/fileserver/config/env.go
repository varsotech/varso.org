package config

import (
	_ "embed"

	"github.com/luminancetech/varso/src/common/config"
)

var (
	FileServerDatabaseName = config.RequireEnv("FILESERVER_DATABASE_NAME")
)
