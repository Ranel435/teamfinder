package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
}

func LoadConfig() Config {
	return Config{
		DatabaseURL: getEnvOrDefault("DATABASE_URL", "postgres://teamfinder_user:teamfinder_password@postgres:5432/teamfinder_db?sslmode=disable"),
		JWTSecret:   getEnvOrDefault("JWT_SECRET", "your-secret-key"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
