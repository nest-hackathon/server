package config

import (
	"log"

	"github.com/nest-hackathon/server/pkg"
)

type Config struct {
	Port    string
	DbUrl   string
	DbName  string
}

func LoadConfig() *Config {
	port, err := pkg.LoadEnv("PORT", "8080")
	if err != nil {
		log.Fatalf("Failed to load PORT environment variable: %v", err)
	}

	dbUrl, err := pkg.LoadEnv("DB_URL", "")
	if err != nil {
		log.Fatalf("Failed to load DB_URL environment variable: %v", err)
	}

	dbName, err := pkg.LoadEnv("DB_NAME", "")
	if err != nil {
		log.Fatalf("Failed to load DB_NAME environment variable: %v", err)
	}

	return &Config{
		Port:    port,
		DbUrl:   dbUrl,
		DbName:  dbName,
	}
}
