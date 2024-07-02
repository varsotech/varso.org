package config

import (
	_ "embed"

	"github.com/luminancetech/varso/src/common/config"
)

var (
	AlertsDiscordWebhook = config.RequireSensitiveEnv("ALERTS_DISCORD_WEBHOOK", config.AppEnv)
)
