package queries

import (
	"context"
	"database/sql"
	"fmt"
	"mbplayer/internal/Models"
	"time"
)

type SongQueries struct {
	db *sql.DB
}

func NewSongQueries(db *sql.DB) *SongQueries {
	return &SongQueries{db: db}
}

// CreateSong creates a new song in the database
func (q *SongQueries) CreateSong(ctx context.Context, name string, artistID int64, playlistID sql.NullInt64, duration float32, releaseDate time.Time) (int64, error) {
	var err error
	var songID int64
	query := `
        INSERT INTO mpdb.songs (title, artist_id, duration, release_year, likes_count, playlist_id)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `
	if playlistID.Valid {
		// Include playlist_id in the query
		err = q.db.QueryRowContext(ctx, query, name, artistID, duration, releaseDate, 0, playlistID.Int64).Scan(&songID)
	} else {
		// Exclude playlist_id from the query
		err = q.db.QueryRowContext(ctx, query, name, artistID, duration, releaseDate, 0, nil).Scan(&songID)
	}

	return songID, err
}

// GetSongByID retrieves a song from the database by its ID
func (q *SongQueries) GetSongByID(ctx context.Context, id int64) (*Models.Song, error) {
	query := `SELECT * FROM songs WHERE id = $1`
	var song Models.Song
	err := q.db.QueryRowContext(ctx, query, id).Scan(&song.ID, &song.Name, &song.ArtistID, &song.PlaylistID)
	if err != nil {
		return nil, err
	}
	return &song, nil
}

// GetSongsByPlaylistID retrieves all songs from the database by their playlist ID
func (q *SongQueries) GetSongsByPlaylistID(ctx context.Context, playlistID int64) ([]*Models.Song, error) {
	query := `SELECT * FROM songs WHERE playlist_id = $1`
	rows, err := q.db.QueryContext(ctx, query, playlistID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var songs []*Models.Song
	for rows.Next() {
		var song Models.Song
		err := rows.Scan(&song.ID, &song.Name, &song.ArtistID, &song.PlaylistID)
		if err != nil {
			return nil, err
		}
		songs = append(songs, &song)
	}
	return songs, nil
}

// GetSongsByArtistID retrieves all songs from the database by their artist ID
func (q *SongQueries) GetSongsByArtistID(ctx context.Context, artistID int64) ([]*Models.Song, error) {
	query := `SELECT * FROM songs WHERE artist_id = $1`
	rows, err := q.db.QueryContext(ctx, query, artistID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var songs []*Models.Song
	for rows.Next() {
		var song Models.Song
		err := rows.Scan(&song.ID, &song.Name, &song.ArtistID, &song.PlaylistID)
		if err != nil {
			return nil, err
		}
		songs = append(songs, &song)
	}
	return songs, nil
}

// GetSongByTitle retrieves a song from the database by its title
func (q *SongQueries) GetSongByTitle(ctx context.Context, title string) (*Models.Song, error) {
	query := `SELECT * FROM songs WHERE title = $1`
	var song Models.Song
	err := q.db.QueryRowContext(ctx, query, title).Scan(&song.ID, &song.Name, &song.ArtistID, &song.PlaylistID)
	if err != nil {
		return nil, err
	}
	return &song, nil
}

// GetAllSongs retrieves all songs from the database
func (q *SongQueries) GetAllSongs(ctx context.Context) ([]*Models.Song, error) {
	query := `SELECT * FROM songs`
	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query all songs: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)

	var songs []*Models.Song
	for rows.Next() {
		var song Models.Song
		if err := rows.Scan(&song.ID, &song.Name, &song.ArtistID, &song.PlaylistID); err != nil {
			return nil, fmt.Errorf("scan song: %w", err)
		}
		songs = append(songs, &song)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows final error: %w", err)
	}
	return songs, nil
}

// UpdateSongTitle updates the title of a song in the database
func (q *SongQueries) UpdateSongTitle(ctx context.Context, id int64, title string) error {
	query := `UPDATE songs SET title = $1 WHERE id = $2`
	_, err := q.db.ExecContext(ctx, query, title, id)
	if err != nil {
		return err
	}
	return nil
}

// IncrementSongLikesCount increments the likes count of a song in the database
func (q *SongQueries) IncrementSongLikesCount(ctx context.Context, id int64) error {
	query := `UPDATE songs SET likes = likes + 1 WHERE id = $1`
	_, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

// DecrementSongLikesCount decrements the likes count of a song in the database
func (q *SongQueries) DecrementSongLikesCount(ctx context.Context, id int64) error {
	query := `UPDATE songs SET likes = likes - 1 WHERE id = $1`
	_, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

// RemoveSongFromPlaylist removes a song from a playlist in the database
func (q *SongQueries) RemoveSongFromPlaylist(ctx context.Context, id int64) error {
	query := `UPDATE songs SET playlist_id = NULL WHERE id = $1`
	_, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSong deletes a song from the database
func (q *SongQueries) DeleteSong(ctx context.Context, id int64) error {
	query := `DELETE FROM songs WHERE id = $1`
	_, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
