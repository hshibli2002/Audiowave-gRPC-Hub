package Handlers

import (
	"Audiowave-gRPC-Hub/internal/Services"
	grpcapi "Audiowave-gRPC-Hub/pkg/grpcapi"
	"context"
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

// ReadArtistFollowerCount reads an artist's follower count by id
func (a *ArtistHandler) ReadArtistFollowerCount(ctx context.Context, req *grpcapi.ReadArtistFollowerCountRequest) (*grpcapi.ReadArtistFollowerCountResponse, error) {
	count, err := a.Service.ReadArtistFollowerCount(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &grpcapi.ReadArtistFollowerCountResponse{
		FollowerCount: count,
	}, nil
}

// ReadArtistLikesCount reads an artist's likes count by id
func (a *ArtistHandler) ReadArtistLikesCount(ctx context.Context, req *grpcapi.ReadArtistLikesCountRequest) (*grpcapi.ReadArtistLikesCountResponse, error) {
	count, err := a.Service.ReadArtistLikesCount(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &grpcapi.ReadArtistLikesCountResponse{
		LikesCount: count,
	}, nil

}

// ReadAllArtists reads all artists
func (a *ArtistHandler) ReadAllArtists(ctx context.Context, req *grpcapi.ReadAllArtistsRequest) (*grpcapi.ReadAllArtistsResponse, error) {
	artists, err := a.Service.ReadAllArtists(ctx)
	if err != nil {
		return nil, err
	}

	var grpcArtists []*grpcapi.Artist
	for _, artist := range artists {
		grpcArtists = append(grpcArtists, &grpcapi.Artist{
			Id:            artist.ID,
			Name:          artist.Name,
			Bio:           artist.Bio,
			FollowerCount: artist.Followers,
			LikesCount:    artist.Likes,
		})
	}

	return &grpcapi.ReadAllArtistsResponse{
		Success: true,
		Artists: grpcArtists,
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
