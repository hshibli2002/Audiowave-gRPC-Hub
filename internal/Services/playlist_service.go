package Services

type PlaylistService interface {
	CreatePlaylist(creatorID int, title string, songIDs []int) (int, error)
	DeletePlaylist(playlistID int) error
	AddSong(artistID int, title string, duration int) (int, error)
	DeleteSong(songID int) error
	CheckFollowers(artistID int) (int, error)
}
