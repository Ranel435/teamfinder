package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() Config {
	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
