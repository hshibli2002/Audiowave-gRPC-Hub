package Handlers

import (
	"context"
	"mbplayer/internal/Services"
	pb "mbplayer/pkg/grpcapi" // import the generated protobuf code
)

type artistServer struct {
	pb.UnimplementedArtistServiceServer
	service Services.ArtistService
}

func NewArtistServer(service Services.ArtistService) pb.ArtistServiceServer {
	return &artistServer{service: service}
}

//TO DO Implement the gRPC server methods

func (s *artistServer) CreatePlaylist(ctx context.Context, req *pb.CreatePlaylistRequest) (*pb.CreatePlaylistResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *artistServer) DeletePlaylist(ctx context.Context, req *pb.DeletePlaylistRequest) (*pb.DeletePlaylistResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *artistServer) AddSong(ctx context.Context, req *pb.AddSongRequest) (*pb.AddSongResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *artistServer) DeleteSong(ctx context.Context, req *pb.DeleteSongRequest) (*pb.DeleteSongResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *artistServer) CheckFollowers(ctx context.Context, req *pb.CheckFollowersRequest) (*pb.CheckFollowersResponse, error) {
	//TODO implement me
	panic("implement me")
}

// ArtistService is the interface that provides methods for the artist service.
