package Handlers

import (
	"context"
	"mbplayer/internal/Services"
	"mbplayer/pkg/grpcapi"
)

type ArtistHandler struct {
	grpcapi.UnimplementedArtistServiceServer
	Service Services.ArtistService
}

func NewArtistHandler(service Services.ArtistService) *ArtistHandler {
	return &ArtistHandler{Service: service}
}

// CreateArtist creates a new artist
func (a *ArtistHandler) CreateArtist(ctx context.Context, req *grpcapi.CreateArtistRequest) (*grpcapi.CreateArtistResponse, error) {
	artist, err := a.Service.CreateArtist(ctx, req.Name, req.Bio)
	if err != nil {
		return nil, err
	}

	return &grpcapi.CreateArtistResponse{
		Artist: &grpcapi.Artist{
			Id:            artist.ID,
			Name:          artist.Name,
			Bio:           artist.Bio,
			FollowerCount: artist.Followers,
			LikesCount:    artist.Likes,
		},
	}, nil
}

// ReadArtistById reads an artist by id
func (a *ArtistHandler) ReadArtistById(ctx context.Context, req *grpcapi.ReadArtistByIdRequest) (*grpcapi.ReadArtistByIdResponse, error) {
	artist, err := a.Service.ReadArtistById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &grpcapi.ReadArtistByIdResponse{
		Artist: &grpcapi.Artist{
			Id:            artist.ID,
			Name:          artist.Name,
			Bio:           artist.Bio,
			FollowerCount: artist.Followers,
			LikesCount:    artist.Likes,
		},
	}, nil
}

// ReadArtistByName reads an artist by name
func (a *ArtistHandler) ReadArtistByName(ctx context.Context, req *grpcapi.ReadArtistByNameRequest) (*grpcapi.ReadArtistByNameResponse, error) {
	artist, err := a.Service.ReadArtistByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &grpcapi.ReadArtistByNameResponse{
		Artist: &grpcapi.Artist{
			Id:            artist.ID,
			Name:          artist.Name,
			Bio:           artist.Bio,
			FollowerCount: artist.Followers,
			LikesCount:    artist.Likes,
		},
	}, nil
}

// ReadArtistBioById reads an artist's bio by id
func (a *ArtistHandler) ReadArtistBioById(ctx context.Context, req *grpcapi.ReadArtistBioRequest) (*grpcapi.ReadArtistBioResponse, error) {
	bio, err := a.Service.ReadArtistBioById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &grpcapi.ReadArtistBioResponse{
		Bio: bio,
	}, nil
}

// UpdateArtistName updates an artist's name
func (a *ArtistHandler) UpdateArtistName(ctx context.Context, req *grpcapi.UpdateArtistNameRequest) (*grpcapi.UpdateArtistNameResponse, error) {
	artist, err := a.Service.UpdateArtistName(ctx, req.Id, req.Name)
	if err != nil {
		return nil, err
	}

	return &grpcapi.UpdateArtistNameResponse{
		Artist: &grpcapi.Artist{
			Id:            artist.ID,
			Name:          artist.Name,
			Bio:           artist.Bio,
			FollowerCount: artist.Followers,
			LikesCount:    artist.Likes,
		},
	}, nil
}

// UpdateArtistBio updates an artist's bio
func (a *ArtistHandler) UpdateArtistBio(ctx context.Context, req *grpcapi.UpdateArtistBioRequest) (*grpcapi.UpdateArtistBioResponse, error) {
	artist, err := a.Service.UpdateArtistBio(ctx, req.Id, req.Bio)
	if err != nil {
		return nil, err
	}

	return &grpcapi.UpdateArtistBioResponse{
		Artist: &grpcapi.Artist{
			Id:            artist.ID,
			Name:          artist.Name,
			Bio:           artist.Bio,
			FollowerCount: artist.Followers,
			LikesCount:    artist.Likes,
		},
	}, nil
}

// UpdateArtistFollowerCount updates an artist's follower count
func (a *ArtistHandler) UpdateArtistFollowerCount(ctx context.Context, req *grpcapi.UpdateArtistFollowerCountRequest) (*grpcapi.UpdateArtistFollowerCountResponse, error) {
	artist, err := a.Service.UpdateArtistFollowerCount(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &grpcapi.UpdateArtistFollowerCountResponse{
		Artist: &grpcapi.Artist{
			Id:            artist.ID,
			Name:          artist.Name,
			Bio:           artist.Bio,
			FollowerCount: artist.Followers,
			LikesCount:    artist.Likes,
		},
	}, nil
}

// UpdateArtistLikesCount updates an artist's likes count
func (a *ArtistHandler) UpdateArtistLikesCount(ctx context.Context, req *grpcapi.UpdateArtistLikesCountRequest) (*grpcapi.UpdateArtistLikesCountResponse, error) {
	artist, err := a.Service.UpdateArtistLikesCount(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &grpcapi.UpdateArtistLikesCountResponse{
		Artist: &grpcapi.Artist{
			Id:            artist.ID,
			Name:          artist.Name,
			Bio:           artist.Bio,
			FollowerCount: artist.Followers,
			LikesCount:    artist.Likes,
		},
	}, nil
}

// DeleteArtistById deletes an artist by id
func (a *ArtistHandler) DeleteArtistById(ctx context.Context, req *grpcapi.DeleteArtistByIdRequest) (*grpcapi.DeleteArtistByIdResponse, error) {
	err := a.Service.DeleteArtistById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &grpcapi.DeleteArtistByIdResponse{}, nil
}

// DeleteAllArtists deletes all artists
func (a *ArtistHandler) DeleteAllArtists(ctx context.Context, req *grpcapi.DeleteAllArtistsRequest) (*grpcapi.DeleteAllArtistsResponse, error) {
	err := a.Service.DeleteAllArtists(ctx)
	if err != nil {
		return nil, err
	}

	return &grpcapi.DeleteAllArtistsResponse{}, nil
}
