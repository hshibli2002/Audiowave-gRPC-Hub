package Handlers

import (
	"context"
	"mbplayer/internal/Services"
	pb "mbplayer/pkg/grpcapi" // import the generated protobuf code
)

type playlistServer struct {
	pb.UnimplementedPlaylistServiceServer
	service Services.PlaylistService
}

func NewPlaylistServer(service Services.PlaylistService) pb.PlaylistServiceServer {
	return &playlistServer{service: service}
}

//TO DO Implement the gRPC server methods

func (s *playlistServer) CreatePlaylist(ctx context.Context, req *pb.CreatePlaylistRequest) (*pb.CreatePlaylistResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *playlistServer) DeletePlaylist(ctx context.Context, req *pb.DeletePlaylistRequest) (*pb.DeletePlaylistResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *playlistServer) AddSong(ctx context.Context, req *pb.AddSongRequest) (*pb.AddSongResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *playlistServer) DeleteSong(ctx context.Context, req *pb.DeleteSongRequest) (*pb.DeleteSongResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *playlistServer) CheckFollowers(ctx context.Context, req *pb.CheckFollowersRequest) (*pb.CheckFollowersResponse, error) {
	//TODO implement me
	panic("implement me")
}

// PlaylistService is the interface that provides methods for the playlist service.
