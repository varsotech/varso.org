package config

import (
	_ "embed"

	"github.com/luminancetech/varso/src/common/config"
)

var (
	AuthDatabaseName  = config.RequireEnv("AUTH_DATABASE_NAME")
	AuthAdminEmail    = config.RequireEnv("AUTH_ADMIN_EMAIL")
	AuthAdminPassword = config.RequireEnv("AUTH_ADMIN_PASSWORD")
)
