package config

import "fmt"

const (
	projectServicesDirectory         = "src/services/"
	projectMigrationsDirectoryFormat = projectServicesDirectory + "%s/internal/ent/migrations/"
)

func ProjectMigrationsDirectory(serviceName string) string {
	return fmt.Sprintf(projectMigrationsDirectoryFormat, serviceName)
}
