package config

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func Get(key, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	return value
}
