package Services

import (
	"context"
	"mbplayer/internal/Models"
	queries "mbplayer/internal/store/Queries"
)

type ArtistService interface {
	CreateArtist(ctx context.Context, name string, bio string) (*Models.Artist, error)
	ReadArtistById(ctx context.Context, id int64) (*Models.Artist, error)
	ReadArtistByName(ctx context.Context, name string) (*Models.Artist, error)
	ReadArtistBioById(ctx context.Context, id int64) (string, error)
	ReadArtistFollowerCount(ctx context.Context, id int64) (int32, error)
	ReadArtistLikesCount(ctx context.Context, id int64) (int32, error)
	UpdateArtistName(ctx context.Context, id int64, name string) (*Models.Artist, error)
	UpdateArtistBio(ctx context.Context, id int64, bio string) (*Models.Artist, error)
	UpdateArtistFollowerCount(ctx context.Context, id int64) (*Models.Artist, error)
	UpdateArtistLikesCount(ctx context.Context, id int64) (*Models.Artist, error)
	DeleteArtistById(ctx context.Context, id int64) error
}

type ArtistServiceImpl struct {
	Queries *queries.ArtistQueries
}

// NewArtistService returns a new ArtistService
func NewArtistService(queries *queries.ArtistQueries) ArtistService {
	return &ArtistServiceImpl{Queries: queries}
}

func (a *ArtistServiceImpl) CreateArtist(ctx context.Context, name string, bio string) (*Models.Artist, error) {
	return a.Queries.CreateArtist(ctx, name, bio)
}

func (a *ArtistServiceImpl) ReadArtistById(ctx context.Context, id int64) (*Models.Artist, error) {
	return a.Queries.ReadArtistById(ctx, id)
}

func (a *ArtistServiceImpl) ReadArtistByName(ctx context.Context, name string) (*Models.Artist, error) {
	return a.Queries.ReadArtistByName(ctx, name)
}

func (a *ArtistServiceImpl) ReadArtistBioById(ctx context.Context, id int64) (string, error) {
	return a.Queries.ReadArtistBioById(ctx, id)
}

func (a *ArtistServiceImpl) ReadArtistFollowerCount(ctx context.Context, id int64) (int32, error) {
	return a.Queries.ReadArtistFollowerCount(ctx, id)
}

func (a *ArtistServiceImpl) ReadArtistLikesCount(ctx context.Context, id int64) (int32, error) {
	return a.Queries.ReadArtistLikesCount(ctx, id)
}

func (a *ArtistServiceImpl) UpdateArtistName(ctx context.Context, id int64, name string) (*Models.Artist, error) {
	return a.Queries.UpdateArtistName(ctx, id, name)
}

func (a *ArtistServiceImpl) UpdateArtistBio(ctx context.Context, id int64, bio string) (*Models.Artist, error) {
	return a.Queries.UpdateArtistBio(ctx, id, bio)
}

func (a *ArtistServiceImpl) UpdateArtistFollowerCount(ctx context.Context, id int64) (*Models.Artist, error) {
	return a.Queries.UpdateArtistFollowerCount(ctx, id)
}

func (a *ArtistServiceImpl) UpdateArtistLikesCount(ctx context.Context, id int64) (*Models.Artist, error) {
	return a.Queries.UpdateArtistLikesCount(ctx, id)
}

func (a *ArtistServiceImpl) DeleteArtistById(ctx context.Context, id int64) error {
	return a.Queries.DeleteArtistById(ctx, id)
}
