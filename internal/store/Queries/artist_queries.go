package queries

import (
	"context"
	"database/sql"
	"mbplayer/internal/Models"
)

type ArtistQueries struct {
	db *sql.DB
}

func NewArtistQueries(db *sql.DB) *ArtistQueries {
	return &ArtistQueries{db: db}
}

func (a *ArtistQueries) CreateArtist(ctx context.Context, name string, bio string) (*Models.Artist, error) {
	query := `INSERT INTO mpdb.artists (name, bio) VALUES ($1, $2) 
                                     RETURNING artist_id, name, bio, followers_count, likes_count`
	row := a.db.QueryRowContext(ctx, query, name, bio)

	var artist Models.Artist
	err := row.Scan(&artist.ID, &artist.Name, &artist.Bio, &artist.Followers, &artist.Likes)
	if err != nil {
		return nil, err
	}

	return &artist, nil

}

func (a *ArtistQueries) ReadArtistById(ctx context.Context, id int64) (*Models.Artist, error) {
	query := `SELECT * FROM mpdb.artists WHERE artist_id = $1`
	row := a.db.QueryRowContext(ctx, query, id)

	var artist Models.Artist
	err := row.Scan(&artist.ID, &artist.Name, &artist.Bio, &artist.Followers, &artist.Likes)
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

func (a *ArtistQueries) ReadArtistByName(ctx context.Context, name string) (*Models.Artist, error) {
	query := `SELECT * FROM mpdb.artists WHERE name = $1`
	row := a.db.QueryRowContext(ctx, query, name)

	var artist Models.Artist
	err := row.Scan(&artist.ID, &artist.Name, &artist.Bio, &artist.Followers, &artist.Likes)
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

func (a *ArtistQueries) ReadArtistBioById(ctx context.Context, id int64) (string, error) {
	query := `SELECT bio FROM mpdb.artists WHERE artist_id = $1`
	row := a.db.QueryRowContext(ctx, query, id)

	var bio string
	err := row.Scan(&bio)
	if err != nil {
		return "", err
	}

	return bio, nil
}

func (a *ArtistQueries) ReadArtistFollowerCount(ctx context.Context, id int64) (int32, error) {
	query := `SELECT followers_count FROM mpdb.artists WHERE artist_id = $1`
	row := a.db.QueryRowContext(ctx, query, id)

	var followers int32
	err := row.Scan(&followers)
	if err != nil {
		return 0, err
	}

	return followers, nil
}

func (a *ArtistQueries) ReadArtistLikesCount(ctx context.Context, id int64) (int32, error) {
	query := `SELECT likes_count FROM mpdb.artists WHERE artist_id = $1`
	row := a.db.QueryRowContext(ctx, query, id)

	var likes int32
	err := row.Scan(&likes)
	if err != nil {
		return 0, err
	}

	return likes, nil
}

func (a *ArtistQueries) UpdateArtistName(ctx context.Context, id int64, name string) (*Models.Artist, error) {
	query := `UPDATE mpdb.artists SET name = $1 WHERE artist_id = $2`
	_, err := a.db.ExecContext(ctx, query, name, id)

	var artist Models.Artist
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

func (a *ArtistQueries) UpdateArtistBio(ctx context.Context, id int64, bio string) (*Models.Artist, error) {
	query := `UPDATE mpdb.artists SET bio = $1 WHERE artist_id = $2`
	_, err := a.db.ExecContext(ctx, query, bio, id)

	var artist Models.Artist
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

func (a *ArtistQueries) UpdateArtistFollowerCount(ctx context.Context, id int64) (*Models.Artist, error) {
	query := `
		UPDATE mpdb.artists
		SET followers_count = followers_count + 1
		WHERE artist_id = $1
		RETURNING artist_id, name, bio, followers_count, likes_count`

	_, err := a.db.ExecContext(ctx, query, id)

	var artist Models.Artist
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

func (a *ArtistQueries) UpdateArtistLikesCount(ctx context.Context, id int64) (*Models.Artist, error) {
	query := `
		UPDATE mpdb.artists
		SET likes_count = likes_count + 1
		WHERE artist_id = $1
		RETURNING artist_id, name, bio, followers_count, likes_count`

	_, err := a.db.ExecContext(ctx, query, id)

	var artist Models.Artist
	if err != nil {
		return nil, err
	}

	return &artist, nil

}

func (a *ArtistQueries) DeleteArtistById(ctx context.Context, id int64) error {
	query := `DELETE FROM mpdb.artists WHERE artist_id = $1`
	_, err := a.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (a *ArtistQueries) DeleteAllArtists(ctx context.Context) error {
	query := `DELETE FROM mpdb.artists`
	_, err := a.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
