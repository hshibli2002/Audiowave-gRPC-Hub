package Models

import "time"

type Song struct {
	ID           int64     `json:"song_id"`
	Title        string    `json:"title"`
	ArtistID     int64     `json:"artist_id"`
	Duration     int32     `json:"duration"`
	CreationTime time.Time `json:"created_at"`
}
