package Models

import (
	"database/sql"
	"time"
)

type Song struct {
	ID          int64         `json:"song_id"`
	Name        string        `json:"name"`
	ArtistID    int64         `json:"artist_id"`   // Foreign key to the Artist
	PlaylistID  sql.NullInt64 `json:"playlist_id"` // Can be null if not part of a Playlist
	Duration    int32         `json:"duration"`
	ReleaseDate time.Time     `json:"release_date"`
	Likes       int32         `json:"likes_count"`
}
