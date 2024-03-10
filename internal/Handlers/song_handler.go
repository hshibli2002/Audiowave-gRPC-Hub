package Handlers

import (
	"context"
	"mbplayer/internal/Services"
	pb "mbplayer/pkg/grpcapi"
)

type songServer struct {
	service Services.SongService
	pb.UnimplementedSongsServiceServer
}

func NewSongServer(service Services.SongService) pb.SongsServiceServer {
	return &songServer{service: service}
}

func (s *songServer) GetSongById(ctx context.Context, req *pb.GetSongByIdRequest) (*pb.GetSongByIdResponse, error) {
	song, err := s.service.GetSongById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	// Construct the response based on your protobuf definition
	response := &pb.GetSongByIdResponse{
		Id:        song.Id,
		Title:     song.Title,
		ArtistId:  song.ArtistId,
		Duration:  song.Duration,
		CreatedAt: song.CreatedAt,
	}
	return response, nil
}

func (s *songServer) CreateSong(ctx context.Context, req *pb.CreateSongRequest) (*pb.CreateSongResponse, error) {
	song, err := s.service.CreateSong(ctx, req.Title, req.ArtistId, req.Duration)
	if err != nil {
		return nil, err
	}
	return &pb.CreateSongResponse{
		Id:        song.Id,
		Title:     song.Title,
		ArtistId:  song.ArtistId,
		Duration:  song.Duration,
		CreatedAt: song.CreatedAt,
	}, nil
}

func (s *songServer) GetsSongsByArtistId(ctx context.Context, req *pb.GetSongsByArtistIdRequest) (*pb.GetSongsByArtistIdResponse, error) {
	songs, err := s.service.GetsSongsByArtistId(ctx, req.ArtistId)
	if err != nil {
		return nil, err
	}
	return &pb.GetSongsByArtistIdResponse{Songs: songs}, nil
}

func (s *songServer) UpdateSongTitle(ctx context.Context, req *pb.UpdateSongTitleRequest) (*pb.UpdateSongTitleResponse, error) {
	song, err := s.service.UpdateSongTitle(ctx, req.Id, req.Title)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateSongTitleResponse{
		Id:       song.Id,
		Title:    song.Title,
		ArtistId: song.ArtistId,
	}, nil
}

func (s *songServer) UpdateSongArtist(ctx context.Context, req *pb.UpdateSongArtistRequest) (*pb.UpdateSongArtistResponse, error) {
	song, err := s.service.UpdateSongArtist(ctx, req.Id, req.ArtistId)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateSongArtistResponse{
		Id:        song.Id,
		Title:     song.Title,
		ArtistId:  song.ArtistId,
		Duration:  song.Duration,
		CreatedAt: song.CreatedAt,
	}, nil
}

func (s *songServer) DeleteSong(ctx context.Context, req *pb.DeleteSongRequest) (*pb.DeleteSongResponse, error) {
	err, _ := s.service.DeleteSong(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteSongResponse{}, nil
}
