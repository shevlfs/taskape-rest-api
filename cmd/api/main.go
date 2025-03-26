package main

import (
	"log"

	"taskape-rest-api/internal/api"
	"taskape-rest-api/internal/config"
	"taskape-rest-api/internal/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	client, err := grpc.NewClient(cfg.BackendHost)
	if err != nil {
		log.Fatalf("Failed to create GRPC client: %v", err)
	}
	defer client.Close()

	server := api.NewServer(cfg, client.BackendClient)
	log.Fatal(server.Start())
}
