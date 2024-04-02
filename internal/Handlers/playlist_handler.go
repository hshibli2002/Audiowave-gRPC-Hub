package Handlers

import (
	"Audiowave-gRPC-Hub/internal/Services"
	grpcapi "Audiowave-gRPC-Hub/pkg/grpcapi"
	"context"
	"strconv"
)

type PlaylistHandler struct {
	grpcapi.UnimplementedPlaylistServiceServer
	Service Services.PlaylistService
}

func NewPlaylistHandler(service Services.PlaylistService) *PlaylistHandler { // HL
	return &PlaylistHandler{Service: service}
}

// CreatePlaylist creates a new playlist
func (p *PlaylistHandler) CreatePlaylist(ctx context.Context, req *grpcapi.CreatePlaylistRequest) (*grpcapi.CreatePlaylistResponse, error) {
	playlistID, err := p.Service.CreatePlaylist(ctx, req.ArtistId, req.Name, req.ReleaseDate)
	if err != nil {
		return nil, err
	}

	return &grpcapi.CreatePlaylistResponse{
		Success: true,
		Message: "Playlist created successfully with ID: " + strconv.FormatInt(playlistID, 10),
	}, nil
}

// GetPlaylistById reads a playlist by id
func (p *PlaylistHandler) GetPlaylistById(ctx context.Context, req *grpcapi.GetPlaylistByIdRequest) (*grpcapi.GetPlaylistByIdResponse, error) {
	playlist, err := p.Service.GetPlaylistByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	releaseDateString := playlist.ReleaseDate.Format("2006-01-02")
	return &grpcapi.GetPlaylistByIdResponse{
		Playlist: &grpcapi.Playlist{
			PlaylistId:  playlist.PlaylistID,
			ArtistId:    playlist.ArtistID,
			Name:        playlist.Name,
			ReleaseDate: releaseDateString,
		},
	}, nil
}

// GetPlaylistsByArtistId reads all playlists by artist id
func (p *PlaylistHandler) GetPlaylistsByArtistId(ctx context.Context, req *grpcapi.GetPlaylistByArtistIdRequest) (*grpcapi.GetPlaylistByArtistIdResponse, error) {
	playlists, err := p.Service.GetPlaylistsByArtistId(ctx, req.ArtistId)
	if err != nil {
		return nil, err
	}

	var grpcPlaylists []*grpcapi.Playlist
	for _, playlist := range playlists {
		releaseDateString := playlist.ReleaseDate.Format("2006-01-02")
		grpcPlaylist := &grpcapi.Playlist{
			PlaylistId:  playlist.PlaylistID,
			ArtistId:    playlist.ArtistID,
			Name:        playlist.Name,
			ReleaseDate: releaseDateString,
		}
		grpcPlaylists = append(grpcPlaylists, grpcPlaylist)
	}

	return &grpcapi.GetPlaylistByArtistIdResponse{
		Playlists: grpcPlaylists,
	}, nil
}

// GetAllPlaylists reads all playlists
func (p *PlaylistHandler) GetAllPlaylists(ctx context.Context, req *grpcapi.GetAllPlaylistsRequest) (*grpcapi.GetAllPlaylistsResponse, error) {
	playlists, err := p.Service.GetAllPlaylists(ctx)
	if err != nil {
		return nil, err
	}

	var grpcPlaylists []*grpcapi.Playlist
	for _, playlist := range playlists {
		releaseDateString := playlist.ReleaseDate.Format("2006-01-02")
		grpcPlaylist := &grpcapi.Playlist{
			PlaylistId:  playlist.PlaylistID,
			ArtistId:    playlist.ArtistID,
			Name:        playlist.Name,
			ReleaseDate: releaseDateString,
		}
		grpcPlaylists = append(grpcPlaylists, grpcPlaylist)
	}

	return &grpcapi.GetAllPlaylistsResponse{
		Success:   true,
		Playlists: grpcPlaylists,
	}, nil
}

// UpdatePlaylistTitle updates a playlist title
func (p *PlaylistHandler) UpdatePlaylistTitle(ctx context.Context, req *grpcapi.UpdatePlaylistTitleRequest) (*grpcapi.UpdatePlaylistTitleResponse, error) {
	err := p.Service.UpdatePlaylistTitle(ctx, req.PlaylistId, req.Name)
	if err != nil {
		return nil, err
	}

	return &grpcapi.UpdatePlaylistTitleResponse{
		Success: true,
		Message: "Playlist title updated successfully",
	}, nil
}

// AddSongToPlaylist adds a song to a playlist
func (p *PlaylistHandler) AddSongToPlaylist(ctx context.Context, req *grpcapi.SongToPlaylistRequest) (*grpcapi.SongToPlaylistResponse, error) {
	err := p.Service.AddSongToPlaylist(ctx, req.PlaylistId, req.SongId)
	if err != nil {
		return nil, err
	}

	return &grpcapi.SongToPlaylistResponse{
		Success: true,
		Message: "Song added to playlist successfully",
	}, nil
}

// RemoveSongFromPlaylist removes a song from a playlist
func (p *PlaylistHandler) RemoveSongFromPlaylist(ctx context.Context, req *grpcapi.SongToPlaylistRequest) (*grpcapi.SongToPlaylistResponse, error) {
	err := p.Service.RemoveSongFromPlaylist(ctx, req.PlaylistId, req.SongId)
	if err != nil {
		return nil, err
	}

	return &grpcapi.SongToPlaylistResponse{
		Success: true,
		Message: "Song removed from playlist successfully",
	}, nil
}

// DeletePlaylist deletes a playlist
func (p *PlaylistHandler) DeletePlaylist(ctx context.Context, req *grpcapi.DeletePlaylistRequest) (*grpcapi.DeletePlaylistResponse, error) {
	err := p.Service.DeletePlaylist(ctx, req.PlaylistId)
	if err != nil {
		return nil, err
	}

	return &grpcapi.DeletePlaylistResponse{
		Success: true,
		Message: "Playlist deleted successfully",
	}, nil
}
