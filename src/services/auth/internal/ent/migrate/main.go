//go:build ignore

package main

import (
	"os"

	_ "github.com/lib/pq"
	"github.com/luminancetech/varso/src/common/database"
	"github.com/luminancetech/varso/src/services/auth/internal/config"
	"github.com/luminancetech/varso/src/services/auth/internal/ent"
	"github.com/luminancetech/varso/src/services/auth/internal/ent/build/migrate"
	"github.com/sirupsen/logrus"
)

// Usage: go generate ./... && go run -mod=mod ./src/services/auth/ent/migrate/main.go <migration_name>

// To generate an initial migration, run:
// go generate ./... && atlas migrate new --dir file://src/services/auth/ent/migrations --dir-format golang-migrate

// To fix checksum mismatch, make sure you don't have conflicts within your migrations, and then run:
// atlas migrate hash --dir file://src/services/auth/ent/migrations
func main() {
	err := database.CreateMigration(config.AuthDirectoryName, config.AuthDatabaseName, os.Args[1], ent.Connect, migrate.NamedDiff)
	if err != nil {
		logrus.WithError(err).Fatalf("failed running migration")
	}
}
