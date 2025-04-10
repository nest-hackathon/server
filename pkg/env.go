package pkg

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


func LoadEnv(field, defaultValue string) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", fmt.Errorf("failed to load .env file: %w", err)
	}

	envValue := os.Getenv(field)
	if envValue == "" {
		return defaultValue, nil
	}
	return envValue, nil
}
