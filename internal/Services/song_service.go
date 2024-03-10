package Services

import (
	"context"
	"fmt"
	"mbplayer/internal/store/Queries"
	pb "mbplayer/pkg/grpcapi"
	"time"
)

type SongService interface {
	CreateSong(ctx context.Context, title string, artistId int64, duration int32) (*pb.Song, error)
	GetSongById(ctx context.Context, songId int64) (*pb.Song, error)
	GetsSongsByArtistId(ctx context.Context, artistId int64) ([]*pb.Song, error)
	UpdateSongTitle(ctx context.Context, songId int64, title string) (*pb.Song, error)
	UpdateSongArtist(ctx context.Context, songId int64, artistId int64) (*pb.Song, error)
	DeleteSong(ctx context.Context, songId int64) (*pb.Song, error)
}

type songServiceImpl struct {
	Queries *queries.SongQueries
}

func NewSongService(queries *queries.SongQueries) SongService {
	return &songServiceImpl{Queries: queries}
}

func (s *songServiceImpl) CreateSong(ctx context.Context, title string, artistId int64, duration int32) (*pb.Song, error) {
	song, err := s.Queries.CreateSong(ctx, title, &artistId, duration)
	if err != nil {
		return nil, fmt.Errorf("failed to create song: %v", err)
	}
	return &pb.Song{
		Id:        song.ID,
		Title:     song.Title,
		ArtistId:  song.ArtistID,
		Duration:  song.Duration,
		CreatedAt: song.CreationTime.Format(time.RFC3339),
	}, nil
}

func (s *songServiceImpl) GetSongById(ctx context.Context, songId int64) (*pb.Song, error) {
	song, err := s.Queries.GetSongById(ctx, songId)
	if err != nil {
		return nil, fmt.Errorf("failed to get song by ID: %v", err)
	}
	return &pb.Song{
		Id:        song.ID,
		Title:     song.Title,
		ArtistId:  song.ArtistID,
		Duration:  song.Duration,
		CreatedAt: song.CreationTime.Format(time.RFC3339),
	}, nil
}

func (s *songServiceImpl) GetsSongsByArtistId(ctx context.Context, artistId int64) ([]*pb.Song, error) {
	songs, err := s.Queries.GetsSongsByArtistId(ctx, artistId)
	if err != nil {
		return nil, fmt.Errorf("failed to get songs by artist ID: %v", err)
	}
	var pbSongs []*pb.Song
	for _, song := range songs {
		pbSongs = append(pbSongs, &pb.Song{
			Id:        song.ID,
			Title:     song.Title,
			ArtistId:  song.ArtistID,
			Duration:  song.Duration,
			CreatedAt: song.CreationTime.Format(time.RFC3339),
		})
	}
	return pbSongs, nil
}

func (s *songServiceImpl) UpdateSongTitle(ctx context.Context, songId int64, title string) (*pb.Song, error) {
	song, err := s.Queries.UpdateSongTitle(ctx, songId, title)
	if err != nil {
		return nil, fmt.Errorf("failed to update song title: %v", err)
	}
	return &pb.Song{
		Id:        song.ID,
		Title:     song.Title,
		ArtistId:  song.ArtistID,
		Duration:  song.Duration,
		CreatedAt: song.CreationTime.Format(time.RFC3339),
	}, nil
}

func (s *songServiceImpl) UpdateSongArtist(ctx context.Context, songId int64, artistId int64) (*pb.Song, error) {
	song, err := s.Queries.UpdateSongArtist(ctx, songId, artistId)
	if err != nil {
		return nil, fmt.Errorf("failed to update song artist: %v", err)
	}
	return &pb.Song{
		Id:        song.ID,
		Title:     song.Title,
		ArtistId:  song.ArtistID,
		Duration:  song.Duration,
		CreatedAt: song.CreationTime.Format(time.RFC3339),
	}, nil
}

func (s *songServiceImpl) DeleteSong(ctx context.Context, songId int64) (*pb.Song, error) {
	err := s.Queries.DeleteSong(ctx, songId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete song: %v", err)
	}
	return nil, nil
}
