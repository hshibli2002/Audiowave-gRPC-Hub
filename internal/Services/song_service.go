package Services

import (
	"context"
	"database/sql"
	"mbplayer/internal/Models"
	queries "mbplayer/internal/store/Queries"
	"time"
)

type SongService interface {
	CreateSong(ctx context.Context, name string, artistID int64, playlistID sql.NullInt64, duration float32, releaseDate time.Time) (int64, error)
	GetSongByID(ctx context.Context, id int64) (*Models.Song, error)
	GetSongsByArtistID(ctx context.Context, artistID int64) ([]*Models.Song, error)
	GetSongByTitle(ctx context.Context, title string) (*Models.Song, error)
	GetSongsByPlaylistID(ctx context.Context, playlistID int64) ([]*Models.Song, error)
	GetAllSongs(ctx context.Context) ([]*Models.Song, error)
	UpdateSongTitle(ctx context.Context, id int64, title string) error
	IncrementSongLikesCount(ctx context.Context, id int64) error
	DecrementSongLikesCount(ctx context.Context, id int64) error
	RemoveSongFromPlaylist(ctx context.Context, songID int64) error
	DeleteSong(ctx context.Context, id int64) error
}

type SongServiceImpl struct {
	Queries *queries.SongQueries
}

func NewSongService(queries *queries.SongQueries) SongService {
	return &SongServiceImpl{Queries: queries}
}

func (s *SongServiceImpl) CreateSong(ctx context.Context, name string, artistID int64, playlistID sql.NullInt64, duration float32, releaseDate time.Time) (int64, error) {
	return s.Queries.CreateSong(ctx, name, artistID, playlistID, duration, releaseDate)
}

func (s *SongServiceImpl) GetSongByID(ctx context.Context, id int64) (*Models.Song, error) {
	return s.Queries.GetSongByID(ctx, id)
}

func (s *SongServiceImpl) GetSongsByArtistID(ctx context.Context, artistID int64) ([]*Models.Song, error) {
	return s.Queries.GetSongsByArtistID(ctx, artistID)
}

func (s *SongServiceImpl) GetSongByTitle(ctx context.Context, title string) (*Models.Song, error) {
	return s.Queries.GetSongByTitle(ctx, title)
}

func (s *SongServiceImpl) GetSongsByPlaylistID(ctx context.Context, playlistID int64) ([]*Models.Song, error) {
	return s.Queries.GetSongsByPlaylistID(ctx, playlistID)
}

func (s *SongServiceImpl) GetAllSongs(ctx context.Context) ([]*Models.Song, error) {
	return s.Queries.GetAllSongs(ctx)
}

func (s *SongServiceImpl) UpdateSongTitle(ctx context.Context, id int64, title string) error {
	return s.Queries.UpdateSongTitle(ctx, id, title)
}

func (s *SongServiceImpl) IncrementSongLikesCount(ctx context.Context, id int64) error {
	return s.Queries.IncrementSongLikesCount(ctx, id)
}

func (s *SongServiceImpl) DecrementSongLikesCount(ctx context.Context, id int64) error {
	return s.Queries.DecrementSongLikesCount(ctx, id)
}

func (s *SongServiceImpl) RemoveSongFromPlaylist(ctx context.Context, songID int64) error {
	return s.Queries.RemoveSongFromPlaylist(ctx, songID)
}

func (s *SongServiceImpl) DeleteSong(ctx context.Context, id int64) error {
	return s.Queries.DeleteSong(ctx, id)
}
