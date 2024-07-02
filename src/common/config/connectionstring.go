package config

import "fmt"

// PostgresConnectionString creates a connection string from the environment variables
func PostgresConnectionString(databaseName string) string {
	// TODO: #6 - Secure postgres communication with SSL
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", PostgresUser, PostgresPassword, PostgresHost, PostgresPort, databaseName)
}

// PostgresConnectionString creates a connection string from the environment variables
func PostgresConnectionStringWithoutDatabase() string {
	// TODO: #6 - Secure postgres communication with SSL
	return fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=disable", PostgresUser, PostgresPassword, PostgresHost, PostgresPort)
}
