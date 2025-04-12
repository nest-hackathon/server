package main

import (
	"log"

	"github.com/nest-hackathon/server/internal/server"
)

func main() {
	server := server.NewServer()

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
