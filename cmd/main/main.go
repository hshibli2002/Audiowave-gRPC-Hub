package main

import (
	"database/sql"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"mbplayer/config"
	"mbplayer/internal/Handlers"
	"mbplayer/internal/Services"
	"mbplayer/internal/store"
	queries "mbplayer/internal/store/Queries"
	pb "mbplayer/pkg/grpcapi"
	"net"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	dbStore, err := store.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatalf("Could not close database: %v", err)
		}
	}(dbStore.DB)

	// Initialize services
	songQueries := queries.NewSongQueries(dbStore.DB)
	songService := Services.NewSongService(songQueries)

	// Create a new gRPC server instance
	grpcServer := grpc.NewServer()

	// Register your services with the gRPC server
	pb.RegisterSongsServiceServer(grpcServer, Handlers.NewSongServer(songService))

	// Determine the port to listen on
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
