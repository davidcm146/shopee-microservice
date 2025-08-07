package config

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv(key, fallback string) string {
	_ = godotenv.Load(".env")
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
