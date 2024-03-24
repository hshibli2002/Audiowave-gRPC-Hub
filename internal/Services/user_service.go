package Services

import (
	"context"
	"mbplayer/internal/Models"
	queries "mbplayer/internal/store/Queries"
)

type UserService interface {
	CreateUser(ctx context.Context, name string, email string, password string) (*Models.User, error)
	GetUserById(ctx context.Context, id int64) (*Models.User, error)
	GetUserByUsername(ctx context.Context, name string) (*Models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*Models.User, error)
	UpdateUsername(ctx context.Context, id int64, username string) error
	UpdateEmail(ctx context.Context, id int64, email string) error
	DeleteUser(ctx context.Context, id int64) error
	IncrementLikes(ctx context.Context, id int64) error
	IncrementFollowingCount(ctx context.Context, id int64) error
}

type UserServiceImpl struct {
	Queries *queries.UserQueries
}

func NewUserService(queries *queries.UserQueries) UserService {
	return &UserServiceImpl{Queries: queries}
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, name string, email string, password string) (*Models.User, error) {
	return u.Queries.CreateUser(ctx, name, email, password)
}

func (u *UserServiceImpl) GetUserById(ctx context.Context, id int64) (*Models.User, error) {
	return u.Queries.GetUserById(ctx, id)
}

func (u *UserServiceImpl) GetUserByUsername(ctx context.Context, Username string) (*Models.User, error) {
	return u.Queries.GetUserByUsername(ctx, Username)
}

func (u *UserServiceImpl) GetUserByEmail(ctx context.Context, email string) (*Models.User, error) {
	return u.Queries.GetUserByEmail(ctx, email)
}

func (u *UserServiceImpl) UpdateUsername(ctx context.Context, id int64, username string) error {
	return u.Queries.UpdateUsername(ctx, id, username)
}

func (u *UserServiceImpl) UpdateEmail(ctx context.Context, id int64, email string) error {
	return u.Queries.UpdateEmail(ctx, id, email)
}

func (u *UserServiceImpl) DeleteUser(ctx context.Context, id int64) error {
	return u.Queries.DeleteUser(ctx, id)
}

func (u *UserServiceImpl) IncrementLikes(ctx context.Context, id int64) error {
	return u.Queries.IncrementLikes(ctx, id)
}

func (u *UserServiceImpl) IncrementFollowingCount(ctx context.Context, id int64) error {
	return u.Queries.IncrementFollowingCount(ctx, id)
}
