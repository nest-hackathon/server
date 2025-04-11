package main

import (
	"log"

	"github.com/nest-hackathon/server/config"
	"github.com/nest-hackathon/server/internal/infra/db"
	"github.com/nest-hackathon/server/internal/server"
)

func main() {
	cfg := config.LoadConfig()

	db, err := db.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	server := server.NewServer()

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	
}