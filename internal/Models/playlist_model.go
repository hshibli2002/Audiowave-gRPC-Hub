package Models

import "time"

type Playlist struct {
	PlaylistID  int64     `json:"playlist_id"`
	Name        string    `json:"name"`
	ArtistID    int64     `json:"artist_id"` // Foreign key to the Artist
	ReleaseDate time.Time `json:"release_date"`
	SongsCount  int32     `json:"songs_count"`
	LikesCount  int32     `json:"likes_count"`
}
