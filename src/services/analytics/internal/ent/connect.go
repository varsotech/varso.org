package ent

import (
	"context"
	"embed"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
	common_config "github.com/luminancetech/varso/src/common/config"
	"github.com/luminancetech/varso/src/common/database"
	"github.com/luminancetech/varso/src/services/analytics/config"
	"github.com/luminancetech/varso/src/services/analytics/internal/ent/build"
	"github.com/pkg/errors"
)

//go:embed migrations/*.sql
var fs embed.FS

var Database *build.Client

func Connect(ctx context.Context) error {
	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return errors.Wrapf(err, "failed creating new iofs for migrations dir")
	}

	err = database.ApplyMigrations(ctx, d, config.AnalyticsDatabaseName)
	if err != nil {
		return errors.Wrapf(err, "failed applying migrations")
	}

	connectionString := common_config.PostgresConnectionString(config.AnalyticsDatabaseName)

	client, err := build.Open("postgres", connectionString)
	if err != nil {
		return errors.Wrapf(err, "failed opening connection to postgres")
	}

	Database = client
	return nil
}
