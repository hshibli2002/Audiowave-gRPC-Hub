package Handlers

import (
	"Audiowave-gRPC-Hub/internal/Services"
	grpcapi "Audiowave-gRPC-Hub/pkg/grpcapi/api/protobuf"
	"context"
	"strconv"
)

type UserHandler struct {
	grpcapi.UnimplementedUserServiceServer
	Service Services.UserService
}

func NewUserHandler(service Services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// CreateUser creates a new user
func (u *UserHandler) CreateUser(ctx context.Context, req *grpcapi.CreateUserRequest) (*grpcapi.CreateUserResponse, error) {
	user, err := u.Service.CreateUser(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &grpcapi.CreateUserResponse{
		User: &grpcapi.User{
			UserId:   user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

// GetUserById reads a user by id

func (u *UserHandler) GetUserById(ctx context.Context, req *grpcapi.GetUserByIdRequest) (*grpcapi.GetUserByIdResponse, error) {
	user, err := u.Service.GetUserById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &grpcapi.GetUserByIdResponse{
		User: &grpcapi.User{
			UserId:   user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

// GetUserByUsername reads a user by username
func (u *UserHandler) GetUserByUsername(ctx context.Context, req *grpcapi.GetUsernameRequest) (*grpcapi.GetUsernameResponse, error) {
	user, err := u.Service.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	return &grpcapi.GetUsernameResponse{
		User: &grpcapi.User{
			UserId:   user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

// GetUserByEmail reads a user by email
func (u *UserHandler) GetUserByEmail(ctx context.Context, req *grpcapi.GetEmailRequest) (*grpcapi.GetEmailResponse, error) {
	user, err := u.Service.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return &grpcapi.GetEmailResponse{
		User: &grpcapi.User{
			UserId:   user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

// GetAllUsers reads all users
func (u *UserHandler) GetAllUsers(ctx context.Context, req *grpcapi.GetAllUsersRequest) (*grpcapi.GetAllUsersResponse, error) {
	users, err := u.Service.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	var grpcUsers []*grpcapi.User
	for _, user := range users {
		grpcUsers = append(grpcUsers, &grpcapi.User{
			UserId:   user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	return &grpcapi.GetAllUsersResponse{
		Success: true,
		Users:   grpcUsers,
	}, nil
}

// UpdateUsername updates a user's username
func (u *UserHandler) UpdateUsername(ctx context.Context, req *grpcapi.UpdateUsernameRequest) (*grpcapi.UpdateUsernameResponse, error) {
	err := u.Service.UpdateUsername(ctx, req.UserId, req.Username)
	if err != nil {
		return nil, err
	}

	return &grpcapi.UpdateUsernameResponse{
		Success: true,
		Message: "Username updated successfully " + req.Username + " for user " + strconv.FormatInt(req.UserId, 10),
	}, nil
}

// UpdateEmail updates a user's email
func (u *UserHandler) UpdateEmail(ctx context.Context, req *grpcapi.UpdateUserEmailRequest) (*grpcapi.UpdateUserEmailResponse, error) {
	err := u.Service.UpdateEmail(ctx, req.UserId, req.Email)
	if err != nil {
		return nil, err
	}

	return &grpcapi.UpdateUserEmailResponse{
		Success: true,
		Message: "Email updated successfully " + req.Email + " for user " + strconv.FormatInt(req.UserId, 10),
	}, nil
}

// DeleteUser deletes a user
func (u *UserHandler) DeleteUser(ctx context.Context, req *grpcapi.DeleteUserRequest) (*grpcapi.DeleteUserResponse, error) {
	err := u.Service.DeleteUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &grpcapi.DeleteUserResponse{
		Success: true,
		Message: "User deleted successfully",
	}, nil
}

// IncrementLikes increments a user's likes
func (u *UserHandler) IncrementLikes(ctx context.Context, req *grpcapi.IncrementLikesRequest) (*grpcapi.IncrementLikesResponse, error) {
	err := u.Service.IncrementLikes(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &grpcapi.IncrementLikesResponse{
		Success: true,
		Message: "Likes incremented successfully for user " + strconv.FormatInt(req.UserId, 10),
	}, nil
}

// IncrementFollowingCount increments a user's following count
func (u *UserHandler) IncrementFollowingCount(ctx context.Context, req *grpcapi.IncrementFollowingCountRequest) (*grpcapi.IncrementFollowingCountResponse, error) {
	err := u.Service.IncrementFollowingCount(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &grpcapi.IncrementFollowingCountResponse{
		Success: true,
		Message: "Following count incremented successfully for user " + strconv.FormatInt(req.UserId, 10),
	}, nil
}
