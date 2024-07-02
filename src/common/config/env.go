package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type appEnvType string

// IsDev returns true if we are in dev mode.
func (a *appEnvType) IsDev() bool {
	return *a == Development
}

const (
	Production  appEnvType = "PROD"
	Development appEnvType = "DEV"
)

var (
	AppEnv                  appEnvType
	PostgresUser            string
	PostgresPassword        string
	PostgresHost            string
	PostgresPort            string
	FileStorageAccessKey    string
	FileStorageSecretKey    string
	FileStorageBucketName   string
	FileStorageEndpoint     string
	SystemExternalURL       string
	JwtSecret               string
	InternalLoginAuthSecret string
	NginxUrl                string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warnf("didn't load .env file: %v", err)
	}

	AppEnv = requireAppEnv("APP_ENV")
	PostgresUser = RequireEnv("POSTGRES_USER")
	PostgresPassword = RequireSensitiveEnv("POSTGRES_PASSWORD", AppEnv)
	PostgresHost = RequireEnv("POSTGRES_HOST")
	PostgresPort = RequireEnv("POSTGRES_PORT")
	FileStorageAccessKey = RequireSensitiveEnv("FILESTORAGE_ACCESS_KEY", AppEnv)
	FileStorageSecretKey = RequireSensitiveEnv("FILESTORAGE_SECRET_KEY", AppEnv)
	FileStorageBucketName = RequireEnv("FILESTORAGE_BUCKET_NAME")
	FileStorageEndpoint = RequireEnv("FILESTORAGE_ENDPOINT")
	SystemExternalURL = RequireEnv("SYSTEM_EXTERNAL_URL")
	JwtSecret = RequireSensitiveEnv("JWT_SECRET", AppEnv)
	InternalLoginAuthSecret = RequireSensitiveEnv("INTERNAL_AUTH_SECRET", AppEnv)
	NginxUrl = RequireEnv("NGINX_URL")
}

func RequireEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		logrus.Panicf("Required environment variable '%s' is not set. Please create a .env file at the root of the project with the appropriate configuration set.", key)
	}

	return val
}

func RequireSensitiveEnv(key string, appEnv appEnvType) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		logrus.Panicf("Required environment variable '%s' is not set. Please create a .env file at the root of the project with the appropriate configuration set.", key)
	}

	if (val == fmt.Sprintf("@{{%s}}", key) || val == "") && !appEnv.IsDev() {
		logrus.Panicf("Environment variable '%s' is set to '%s' in a production deployment. Please make sure your environment variables are set to safe secret UUIDs.", val, key)
	}

	return val
}

func requireAppEnv(key string) appEnvType {
	val := RequireEnv(key)
	switch appEnvType(val) {
	case Development:
		return Development
	case Production:
		return Production
	default:
		logrus.Errorf("Invalid %s provided, falling back to %s", key, Production)
		return Production
	}
}
