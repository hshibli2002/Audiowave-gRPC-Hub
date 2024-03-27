package main

import (
	"Audiowave-gRPC-Hub/config"
	"Audiowave-gRPC-Hub/internal/Handlers"
	"Audiowave-gRPC-Hub/internal/Services"
	"Audiowave-gRPC-Hub/internal/store"
	queries "Audiowave-gRPC-Hub/internal/store/Queries"
	pb "Audiowave-gRPC-Hub/pkg/grpcapi/api/protobuf"
	"database/sql"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg, err := loadAndValidateConfig()
	if err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	dbStore, err := initializeDatabase(cfg)
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}(dbStore.DB)

	grpcServer := initializeGRPCServer(dbStore)

	startGRPCServer(grpcServer, cfg.GRPCPort)
}
func loadAndValidateConfig() (*config.Config, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func initializeDatabase(cfg *config.Config) (*store.DBStore, error) {
	dbStore, err := store.InitDB(cfg)
	if err != nil {
		return nil, err
	}
	return dbStore, nil
}

func initializeGRPCServer(dbStore *store.DBStore) *grpc.Server {
	artistQueries := queries.NewArtistQueries(dbStore.DB)
	userQueries := queries.NewUserQueries(dbStore.DB)
	songQueries := queries.NewSongQueries(dbStore.DB)
	playlistQueries := queries.NewPlaylistQueries(dbStore.DB)

	artistService := Services.NewArtistService(artistQueries)
	userService := Services.NewUserService(userQueries)
	songService := Services.NewSongService(songQueries)
	playlistService := Services.NewPlaylistService(playlistQueries)

	artistHandler := Handlers.NewArtistHandler(artistService)
	userHandler := Handlers.NewUserHandler(userService)
	songHandler := Handlers.NewSongHandler(songService)
	playlistHandler := Handlers.NewPlaylistHandler(playlistService)

	grpcServer := grpc.NewServer()
	pb.RegisterArtistServiceServer(grpcServer, artistHandler)
	pb.RegisterUserServiceServer(grpcServer, userHandler)
	pb.RegisterSongsServiceServer(grpcServer, songHandler)
	pb.RegisterPlaylistServiceServer(grpcServer, playlistHandler)

	return grpcServer
}

func startGRPCServer(grpcServer *grpc.Server, port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	fmt.Printf("gRPC server listening on port %s\n", port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
