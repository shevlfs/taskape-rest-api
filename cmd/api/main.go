package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"taskape-rest-api/internal/api"
	"taskape-rest-api/internal/config"
	"taskape-rest-api/internal/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create the connection manager instead of simple client
	connectionManager := grpc.NewConnectionManager(cfg.BackendHost)
	defer connectionManager.Close()

	server := api.NewServer(cfg, connectionManager)

	// Handle graceful shutdown
	go handleShutdown(server, connectionManager)

	// Start the server
	log.Fatal(server.Start())
}

// handleShutdown sets up graceful shutdown on SIGINT/SIGTERM
func handleShutdown(server *api.Server, connectionManager *grpc.ConnectionManager) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh
	log.Println("Received shutdown signal, closing connections...")

	// Close connections
	if err := connectionManager.Close(); err != nil {
		log.Printf("Error closing gRPC connection: %v", err)
	}

	log.Println("Shutdown complete")
	os.Exit(0)
}
