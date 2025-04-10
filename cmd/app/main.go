package main

import (
	"log"
	"os"

	"github.com/nest-hackathon/server/internal/server"
)

func main() {
	server := server.NewServer()

	if err := server.Start(); err != nil {
		os.Exit(1)
		log.Fatalf("Failed to start server: %v", err)
	}
}