package queries

import (
	"Audiowave-gRPC-Hub/internal/Models"
	"context"
	"database/sql"
)

type PlaylistQueries struct {
	db *sql.DB
}

func NewPlaylistQueries(db *sql.DB) *PlaylistQueries {
	return &PlaylistQueries{db: db}
}

// CreatePlaylist creates a new playlist in the database
func (q *PlaylistQueries) CreatePlaylist(ctx context.Context, artistID int64, name string, releaseDate string) (int64, error) {
	var err error
	var playlistID int64
	query := `
		INSERT INTO mpdb.playlists (name, artist_id, release_date)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err = q.db.QueryRowContext(ctx, query, name, artistID, releaseDate).Scan(&playlistID)
	return playlistID, err
}

// GetPlaylistByID retrieves a playlist from the database by its ID
func (q *PlaylistQueries) GetPlaylistByID(ctx context.Context, id int64) (*Models.Playlist, error) {
	query := `SELECT * FROM playlists WHERE id = $1`
	var playlist Models.Playlist
	err := q.db.QueryRowContext(ctx, query, id).Scan(&playlist.PlaylistID, &playlist.Name, &playlist.ArtistID, &playlist.ReleaseDate, &playlist.SongsCount, &playlist.LikesCount)
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}

// GetPlaylistsByArtistID retrieves all playlists from the database by their artist ID
func (q *PlaylistQueries) GetPlaylistsByArtistID(ctx context.Context, artistID int64) ([]*Models.Playlist, error) {
	query := `SELECT * FROM playlists WHERE artist_id = $1`
	rows, err := q.db.QueryContext(ctx, query, artistID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var playlists []*Models.Playlist
	for rows.Next() {
		var playlist Models.Playlist
		err = rows.Scan(&playlist.PlaylistID, &playlist.Name, &playlist.ArtistID, &playlist.ReleaseDate, &playlist.SongsCount, &playlist.LikesCount)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, &playlist)
	}

	return playlists, nil
}

// GetPlaylistByTitle retrieves a playlist from the database by its title using context
func (q *PlaylistQueries) GetPlaylistByTitle(ctx context.Context, title string) (*Models.Playlist, error) {
	query := `SELECT * FROM playlists WHERE name = $1`
	var playlist Models.Playlist
	err := q.db.QueryRowContext(ctx, query, title).Scan(&playlist.PlaylistID, &playlist.Name, &playlist.ArtistID, &playlist.ReleaseDate, &playlist.SongsCount, &playlist.LikesCount)
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}

// GetPlaylistByArtistId retrieves a playlist from the database by its artist ID
func (q *PlaylistQueries) GetPlaylistByArtistId(ctx context.Context, artistID int64) (*Models.Playlist, error) {
	query := `SELECT * FROM playlists WHERE artist_id = $1`
	var playlist Models.Playlist
	err := q.db.QueryRow(query, artistID).Scan(&playlist.PlaylistID, &playlist.Name, &playlist.ArtistID, &playlist.ReleaseDate, &playlist.SongsCount, &playlist.LikesCount)
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}

// GetAllPlaylists retrieves all playlists from the database
func (q *PlaylistQueries) GetAllPlaylists(ctx context.Context) ([]*Models.Playlist, error) {
	query := `SELECT * FROM playlists`
	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var playlists []*Models.Playlist
	for rows.Next() {
		var playlist Models.Playlist
		err = rows.Scan(&playlist.PlaylistID, &playlist.Name, &playlist.ArtistID, &playlist.ReleaseDate, &playlist.SongsCount, &playlist.LikesCount)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, &playlist)
	}

	return playlists, nil
}

// UpdatePlaylistTitle updates the title of a playlist in the database
func (q *PlaylistQueries) UpdatePlaylistTitle(ctx context.Context, id int64, title string) error {
	query := `UPDATE playlists SET name = $1 WHERE id = $2`
	_, err := q.db.ExecContext(ctx, query, title, id)
	return err
}

// AddSongToPlaylist adds a song to a playlist in the database, increments both playlist song count and sets the song playlist_id to the playlist id
func (q *PlaylistQueries) AddSongToPlaylist(ctx context.Context, playlistID int64, songID int64) error {
	query := `UPDATE songs SET playlist_id = $1 WHERE id = $2`
	_, err := q.db.ExecContext(ctx, query, playlistID, songID)
	if err != nil {
		return err
	}

	query = `UPDATE playlists SET songs_count = songs_count + 1 WHERE id = $1`
	_, err = q.db.ExecContext(ctx, query, playlistID)
	return err
}

// RemoveSongFromPlaylist removes a song from a playlist in the database, decrements both playlist song count and sets the song playlist_id to NULL
func (q *PlaylistQueries) RemoveSongFromPlaylist(ctx context.Context, playlistID int64, songID int64) error {
	query := `UPDATE songs SET playlist_id = NULL WHERE id = $1`
	_, err := q.db.ExecContext(ctx, query, songID)
	if err != nil {
		return err
	}

	query = `UPDATE playlists SET songs_count = songs_count - 1 WHERE id = $1`
	_, err = q.db.ExecContext(ctx, query, playlistID)
	return err

}

// DeletePlaylist deletes a playlist from the database
func (q *PlaylistQueries) DeletePlaylist(ctx context.Context, id int64) error {
	query := `DELETE FROM playlists WHERE id = $1`
	_, err := q.db.ExecContext(ctx, query, id)
	return err
}
