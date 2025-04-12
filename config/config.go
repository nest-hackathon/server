package config

import (
	"log"

	"github.com/nest-hackathon/server/pkg"
)

type Config struct {
	Port     string
	DbUrl    string
	JwtKey   string
	RedisUrl string
}

func LoadConfig() *Config {
	port, err := pkg.LoadEnv("PORT", "8080")
	if err != nil {
		log.Fatalf("Failed to load PORT environment variable: %v", err)
	}

	dbUrl, err := pkg.LoadEnv("DB_PATH", "")
	if err != nil {
		log.Fatalf("Failed to load DB_PATH environment variable: %v", err)
	}

	jwtKey, err := pkg.LoadEnv("JWT_KEY", "")
	if err != nil {
		log.Fatalf("Failed to load JWT_KEY environment variable: %v", err)
	}

	redisUrl, err := pkg.LoadEnv("REDIS_URL", "localhost:6379")
	if err != nil {
		log.Fatalf("Failed to load REDIS_URL environment variable: %v", err)
	}

	return &Config{
		Port:     port,
		DbUrl:    dbUrl,
		JwtKey:   jwtKey,
		RedisUrl: redisUrl,
	}
}
