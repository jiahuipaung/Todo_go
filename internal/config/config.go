package config

import (
	"os"
)

var DatabaseDSN string

func LoadConfig() {
	DatabaseDSN = getEnv("DATABASE_DSN", "appuser:123456@tcp(127.0.0.1:3306)/todo_app")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
