package queries

import (
	"context"
	"database/sql"
	models "mbplayer/internal/Models" // Import the models package
	_ "time"
)

type SongQueries struct {
	db *sql.DB
}

func NewSongQueries(db *sql.DB) *SongQueries {
	return &SongQueries{db: db}
}

func (q *SongQueries) CreateSong(ctx context.Context, title string, artistIdPtr *int64, duration int32) (*models.Song, error) {
	var (
		query    string
		artistId sql.NullInt64 // using sql.NullInt32 to allow NULL values
	)

	if artistIdPtr != nil {
		artistId = sql.NullInt64{Int64: *artistIdPtr, Valid: true}
	}

	// Use COALESCE to set the artist_id to NULL if the input artistId is 0 or invalid
	query = `
        INSERT INTO songs (title, artist_id, duration, created_at)
        VALUES ($1, COALESCE(NULLIF($2, 0), NULL), $3, NOW())
        RETURNING song_id, title, artist_id, duration, created_at;
    `

	var song models.Song
	err := q.db.QueryRowContext(ctx, query, title, artistId, duration).Scan(
		&song.ID,
		&song.Title,
		&artistId,
		&song.Duration,
		&song.CreationTime,
	)
	if err != nil {
		return nil, err
	}

	if artistId.Valid {
		song.ArtistID = artistId.Int64
	} else {
		song.ArtistID = 0
	}

	return &song, nil
}

func (q *SongQueries) GetSongById(ctx context.Context, songId int64) (*models.Song, error) {
	var song models.Song
	query := `SELECT song_id, title, artist_id, duration, created_at FROM songs WHERE song_id = $1;`
	err := q.db.QueryRowContext(ctx, query, songId).Scan(
		&song.ID,
		&song.Title,
		&song.ArtistID,
		&song.Duration,
		&song.CreationTime,
	)
	if err != nil {
		return nil, err
	}
	return &song, nil
}

func (q *SongQueries) GetsSongsByArtistId(ctx context.Context, artistId int64) ([]*models.Song, error) {
	var songs []*models.Song
	query := `SELECT song_id, title, artist_id, duration, created_at FROM songs WHERE artist_id = $1;`
	rows, err := q.db.QueryContext(ctx, query, artistId)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var song models.Song
		if err := rows.Scan(&song.ID, &song.Title, &song.ArtistID, &song.Duration, &song.CreationTime); err != nil {
			return nil, err
		}
		songs = append(songs, &song)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return songs, nil
}

func (q *SongQueries) UpdateSongTitle(ctx context.Context, songId int64, title string) (*models.Song, error) {
	query := `UPDATE songs SET title = $2 WHERE song_id = $1 RETURNING song_id, title, artist_id, duration, created_at;`
	var song models.Song
	err := q.db.QueryRowContext(ctx, query, songId, title).Scan(&song.ID, &song.Title, &song.ArtistID, &song.Duration, &song.CreationTime)
	if err != nil {
		return nil, err
	}
	return &song, nil
}

func (q *SongQueries) UpdateSongArtist(ctx context.Context, songId, artistId int64) (*models.Song, error) {
	query := `UPDATE songs SET artist_id = $2 WHERE song_id = $1 RETURNING song_id, title, artist_id, duration, created_at;`
	var song models.Song
	err := q.db.QueryRowContext(ctx, query, songId, artistId).Scan(&song.ID, &song.Title, &song.ArtistID, &song.Duration, &song.CreationTime)
	if err != nil {
		return nil, err
	}
	return &song, nil
}

func (q *SongQueries) DeleteSong(ctx context.Context, songId int64) error {
	query := `DELETE FROM songs WHERE song_id = $1;`
	_, err := q.db.ExecContext(ctx, query, songId)
	return err
}
