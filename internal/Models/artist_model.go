package Models

type Artist struct {
	ID        int64  `json:"artist_id"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	Followers int32  `json:"followers_count"`
	Likes     int    `json:"likes_count"`
}
