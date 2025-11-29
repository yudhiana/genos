package main

import (
	"genos/internal/bootstrap"
	"log"
)

func main() {
	// Initialize bootstrap container
	container := bootstrap.NewContainer()

	// sentry.InitSentry()

	// Initialize and start server
	err := container.StartServer()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
