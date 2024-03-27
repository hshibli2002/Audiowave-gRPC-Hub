package Services

import (
	"Audiowave-gRPC-Hub/internal/Models"
	queries "Audiowave-gRPC-Hub/internal/store/Queries"
	"context"
)

type PlaylistService interface {
	CreatePlaylist(ctx context.Context, artistID int64, name string, releaseDate string) (int64, error)
	GetPlaylistByID(ctx context.Context, id int64) (*Models.Playlist, error)
	GetPlaylistByTitle(ctx context.Context, title string) (*Models.Playlist, error)
	GetPlaylistsByArtistId(ctx context.Context, artistID int64) ([]*Models.Playlist, error)
	GetAllPlaylists(ctx context.Context) ([]*Models.Playlist, error)
	UpdatePlaylistTitle(ctx context.Context, id int64, title string) error
	AddSongToPlaylist(ctx context.Context, playlistID int64, songID int64) error
	RemoveSongFromPlaylist(ctx context.Context, playlistID int64, songID int64) error
	DeletePlaylist(ctx context.Context, id int64) error
}

type PlaylistServiceImpl struct {
	Queries *queries.PlaylistQueries
}

func NewPlaylistService(queries *queries.PlaylistQueries) PlaylistService {
	return &PlaylistServiceImpl{Queries: queries}
}

func (s *PlaylistServiceImpl) CreatePlaylist(ctx context.Context, artistID int64, name string, releaseDate string) (int64, error) {
	return s.Queries.CreatePlaylist(ctx, artistID, name, releaseDate)
}

func (s *PlaylistServiceImpl) GetPlaylistByID(ctx context.Context, id int64) (*Models.Playlist, error) {
	return s.Queries.GetPlaylistByID(ctx, id)
}

func (s *PlaylistServiceImpl) GetPlaylistByTitle(ctx context.Context, title string) (*Models.Playlist, error) {
	return s.Queries.GetPlaylistByTitle(ctx, title)
}

func (s *PlaylistServiceImpl) GetPlaylistsByArtistId(ctx context.Context, artistID int64) ([]*Models.Playlist, error) {
	return s.Queries.GetPlaylistsByArtistID(ctx, artistID)
}

func (s *PlaylistServiceImpl) GetAllPlaylists(ctx context.Context) ([]*Models.Playlist, error) {
	return s.Queries.GetAllPlaylists(ctx)
}

func (s *PlaylistServiceImpl) UpdatePlaylistTitle(ctx context.Context, id int64, title string) error {
	return s.Queries.UpdatePlaylistTitle(ctx, id, title)
}

func (s *PlaylistServiceImpl) AddSongToPlaylist(ctx context.Context, playlistID int64, songID int64) error {
	return s.Queries.AddSongToPlaylist(ctx, playlistID, songID)
}

func (s *PlaylistServiceImpl) RemoveSongFromPlaylist(ctx context.Context, playlistID int64, songID int64) error {
	return s.Queries.RemoveSongFromPlaylist(ctx, playlistID, songID)
}

func (s *PlaylistServiceImpl) DeletePlaylist(ctx context.Context, id int64) error {
	return s.Queries.DeletePlaylist(ctx, id)
}
