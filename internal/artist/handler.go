// internal/artist/service_impl.go

package artist

import "mbplayer/internal/store"

type artistController struct {
	store *store.Store // This is a reference to your data store
}

func NewService(store *store.Store) Service {
	return &artistController{store: store}
}

func (s *artistController) CreatePlaylist(creatorID int, title string, songIDs []int) (int, error) {
	//TODO implement me
	panic("implement me")
}
func (s *artistController) DeletePlaylist(playlistID int) error {
	//TODO implement me
	panic("implement me")
}

func (s *artistController) AddSong(artistID int, title string, duration int) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s *artistController) DeleteSong(songID int) error {
	//TODO implement me
	panic("implement me")
}

func (s *artistController) CheckFollowers(artistID int) (int, error) {
	//TODO implement me
	panic("implement me")
}
