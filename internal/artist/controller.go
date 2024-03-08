package artist

import "mbplayer/internal/store"

// serviceImpl implements the ArtistService interface.
type serviceImpl struct {
	store *store.Store
}

func NewArtistService(store *store.Store) ArtistService {
	return &serviceImpl{store: store}
}

func (s serviceImpl) CreatePlaylist(creatorID int, title string, songIDs []int) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s serviceImpl) DeletePlaylist(playlistID int) error {
	//TODO implement me
	panic("implement me")
}

func (s serviceImpl) AddSong(artistID int, title string, duration int) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s serviceImpl) DeleteSong(songID int) error {
	//TODO implement me
	panic("implement me")
}

func (s serviceImpl) CheckFollowers(artistID int) (int, error) {
	//TODO implement me
	panic("implement me")
}
