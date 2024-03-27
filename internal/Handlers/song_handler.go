package Handlers

import (
	"Audiowave-gRPC-Hub/internal/Services"
	grpcapi "Audiowave-gRPC-Hub/pkg/grpcapi/api/protobuf"
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
)

type SongHandler struct {
	grpcapi.UnimplementedSongsServiceServer
	Service Services.SongService
}

func NewSongHandler(service Services.SongService) *SongHandler {
	return &SongHandler{Service: service}
}

// CreateSong creates a new song in the database
func (s *SongHandler) CreateSong(ctx context.Context, req *grpcapi.CreateSongRequest) (*grpcapi.CreateSongResponse, error) {
	// Parse PlaylistID
	var playlistID sql.NullInt64
	if req.PlaylistId != 0 {
		playlistID = sql.NullInt64{Int64: req.PlaylistId, Valid: true}
	}

	releaseDate, err := time.Parse("2006-01-02", req.ReleaseDate)
	if err != nil {
		// Return an error response if the date format is incorrect.
		return nil, status.Errorf(codes.InvalidArgument, "invalid release date format: %v", err)
	}

	songID, err := s.Service.CreateSong(ctx, req.Title, req.ArtistId, playlistID, req.Duration, releaseDate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating song: %v", err)
	}

	return &grpcapi.CreateSongResponse{
		Success: true,
		Message: "Song created successfully with ID: " + strconv.FormatInt(songID, 10),
	}, nil
}

func (s *SongHandler) GetSongById(ctx context.Context, req *grpcapi.GetSongByIdRequest) (*grpcapi.GetSongByIdResponse, error) {
	song, err := s.Service.GetSongByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &grpcapi.GetSongByIdResponse{
		Song: &grpcapi.Song{
			Id:         song.ID,
			Title:      song.Name,
			ArtistId:   song.ArtistID,
			Duration:   song.Duration,
			PlaylistId: song.PlaylistID.Int64,
		},
	}, nil
}

func (s *SongHandler) GetSongsByPlaylistId(ctx context.Context, req *grpcapi.GetSongsByPlaylistIdRequest) (*grpcapi.GetSongsByPlaylistIdResponse, error) {
	songs, err := s.Service.GetSongsByPlaylistID(ctx, req.PlaylistId)
	if err != nil {
		return nil, err
	}

	var grpcSongs []*grpcapi.Song
	for _, song := range songs {
		grpcSongs = append(grpcSongs, &grpcapi.Song{
			Id:         song.ID,
			Title:      song.Name,
			ArtistId:   song.ArtistID,
			Duration:   song.Duration,
			PlaylistId: song.PlaylistID.Int64,
		})
	}

	return &grpcapi.GetSongsByPlaylistIdResponse{
		Songs: grpcSongs,
	}, nil
}

func (s *SongHandler) GetSongsByArtistId(ctx context.Context, req *grpcapi.GetSongsByArtistIdRequest) (*grpcapi.GetSongsByArtistIdResponse, error) {
	songs, err := s.Service.GetSongsByArtistID(ctx, req.ArtistId)
	if err != nil {
		return nil, err
	}

	var grpcSongs []*grpcapi.Song
	for _, song := range songs {
		grpcSongs = append(grpcSongs, &grpcapi.Song{
			Id:         song.ID,
			Title:      song.Name,
			ArtistId:   song.ArtistID,
			Duration:   song.Duration,
			PlaylistId: song.PlaylistID.Int64,
		})
	}

	return &grpcapi.GetSongsByArtistIdResponse{
		Songs: grpcSongs,
	}, nil
}

func (s *SongHandler) GetAllSongs(ctx context.Context, req *grpcapi.GetAllSongsRequest) (*grpcapi.GetAllSongsResponse, error) {
	songs, err := s.Service.GetAllSongs(ctx)
	if err != nil {
		return nil, err
	}

	var grpcSongs []*grpcapi.Song
	for _, song := range songs {
		grpcSongs = append(grpcSongs, &grpcapi.Song{
			Id:         song.ID,
			Title:      song.Name,
			ArtistId:   song.ArtistID,
			Duration:   song.Duration,
			PlaylistId: song.PlaylistID.Int64,
		})
	}

	return &grpcapi.GetAllSongsResponse{
		Success: true,
		Songs:   grpcSongs,
	}, nil
}

func (s *SongHandler) UpdateSongTitle(ctx context.Context, req *grpcapi.UpdateSongTitleRequest) (*grpcapi.UpdateSongTitleResponse, error) {
	err := s.Service.UpdateSongTitle(ctx, req.Id, req.Title)
	if err != nil {
		return nil, err
	}

	return &grpcapi.UpdateSongTitleResponse{
		Success: true,
		Message: "Song title updated successfully",
	}, nil
}

func (s *SongHandler) IncrementSongLikesCount(ctx context.Context, req *grpcapi.IncrementSongLikesCountRequest) (*grpcapi.IncrementSongLikesCountResponse, error) {
	err := s.Service.IncrementSongLikesCount(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &grpcapi.IncrementSongLikesCountResponse{
		Success: true,
		Message: "Song likes count incremented successfully",
	}, nil
}

func (s *SongHandler) DecrementSongLikesCount(ctx context.Context, req *grpcapi.DecrementSongLikesCountRequest) (*grpcapi.DecrementSongLikesCountResponse, error) {
	err := s.Service.DecrementSongLikesCount(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &grpcapi.DecrementSongLikesCountResponse{
		Success: true,
		Message: "Song likes count decremented successfully",
	}, nil
}

func (s *SongHandler) RemoveSongFromPlaylist(ctx context.Context, req *grpcapi.RemoveSongFromPlaylistRequest) (*grpcapi.RemoveSongFromPlaylistResponse, error) {
	err := s.Service.RemoveSongFromPlaylist(ctx, req.GetSongId())
	if err != nil {
		return nil, err
	}

	return &grpcapi.RemoveSongFromPlaylistResponse{
		Success: true,
		Message: "Song removed from playlist successfully",
	}, nil
}

func (s *SongHandler) DeleteSong(ctx context.Context, req *grpcapi.DeleteSongRequest) (*grpcapi.DeleteSongResponse, error) {
	err := s.Service.DeleteSong(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &grpcapi.DeleteSongResponse{
		Success: true,
		Message: "Song deleted successfully",
	}, nil
}
