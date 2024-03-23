package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"mbplayer/config"
	"mbplayer/internal/store"
	"net"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize the database connection
	dbStore, err := store.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer func() {
		err := dbStore.DB.Close()
		if err != nil {
			log.Fatalf("Could not close database: %v", err)
		}
	}()

	fmt.Println("Successfully connected to the database")

	// Create a new gRPC server instance
	grpcServer := grpc.NewServer()

	port := "50051"
	if envPort := os.Getenv("GRPC_PORT"); envPort != "" {
		port = envPort
	}

	// Start listening on the specified port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	fmt.Printf("gRPC server listening on port %s\n", port)

	// Start serving gRPC requests
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}

}
