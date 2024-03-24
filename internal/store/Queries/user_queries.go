package queries

import (
	"context"
	"database/sql"
	"mbplayer/internal/Models"
)

type UserQueries struct {
	db *sql.DB
}

func NewUserQueries(db *sql.DB) *UserQueries {
	return &UserQueries{db: db}
}

func (u *UserQueries) CreateUser(ctx context.Context, username string, email string, password string) (*Models.User, error) {
	query := `INSERT INTO mpdb.users (username, email, password) VALUES ($1, $2, $3) 
									 RETURNING user_id, username, email, password`
	row := u.db.QueryRowContext(ctx, query, username, email, password)

	var user Models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserQueries) GetUserById(ctx context.Context, id int64) (*Models.User, error) {
	query := `SELECT * FROM mpdb.users WHERE user_id = $1`
	row := u.db.QueryRowContext(ctx, query, id)

	var user Models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserQueries) GetUserByUsername(ctx context.Context, username string) (*Models.User, error) {
	query := `SELECT * FROM mpdb.users WHERE username = $1`
	row := u.db.QueryRowContext(ctx, query, username)

	var user Models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserQueries) GetUserByEmail(ctx context.Context, email string) (*Models.User, error) {
	query := `SELECT * FROM mpdb.users WHERE email = $1`
	row := u.db.QueryRowContext(ctx, query, email)

	var user Models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserQueries) UpdateUsername(ctx context.Context, id int64, username string) error {
	query := `UPDATE mpdb.users SET username = $1 WHERE user_id = $2`
	_, err := u.db.ExecContext(ctx, query, username, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserQueries) UpdateEmail(ctx context.Context, id int64, email string) error {
	query := `UPDATE mpdb.users SET email = $1 WHERE user_id = $2`
	_, err := u.db.ExecContext(ctx, query, email, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserQueries) DeleteUser(ctx context.Context, id int64) error {
	query := `DELETE FROM mpdb.users WHERE user_id = $1`
	_, err := u.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserQueries) IncrementLikes(ctx context.Context, id int64) error {
	query := `UPDATE mpdb.users SET likes_given = likes_given + 1 WHERE user_id = $1`
	_, err := u.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserQueries) IncrementFollowingCount(ctx context.Context, id int64) error {
	query := `UPDATE mpdb.users SET follows_given = users.follows_given + 1 WHERE user_id = $1`
	_, err := u.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
