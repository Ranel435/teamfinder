package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
	// Добавьте другие конфигурационные параметры здесь
}

func LoadConfig() Config {
	return Config{
		DatabaseURL: getEnvOrDefault("DATABASE_URL", "postgres://teamfinder:teamfinder@postgres:5432/postgres?sslmode=disable"),
		JWTSecret:   getEnvOrDefault("JWT_SECRET", "your-secret-key"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
