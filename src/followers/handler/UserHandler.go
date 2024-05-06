package handler

import (
	"context"
	"followers/model"
	"followers/proto/followers"
	"followers/service"
	"log"
)

type UserHandler struct {
	UserService *service.UserService
	logger      *log.Logger
	followers.UnimplementedFollowersServiceServer
}

type KeyProduct struct{}

func NewUserHandler(l *log.Logger, s *service.UserService) *UserHandler {
	return &UserHandler{s, l, followers.UnimplementedFollowersServiceServer{}} //Proveriti samo sus je
}

func (handler UserHandler) GetByUsername(ctx context.Context, request *followers.GetUserByUsernameRequest) (*followers.GetUserByUsernameResponse, error) {
	user, _ := handler.UserService.GetByUsername(request.Username)

	userResponse := followers.User{}
	userResponse.ID = user.ID
	userResponse.Username = user.Username
	userResponse.Password = user.Password
	userResponse.IsActive = user.IsActive
	userResponse.ProfilePicture = user.ProfilePicture
	userResponse.Role = followers.User_Role(user.Role)

	return &followers.GetUserByUsernameResponse{
		User: &userResponse,
	}, nil
}

func (handler UserHandler) GetFollowers(ctx context.Context, request *followers.GetFollowersRequest) (*followers.GetFollowersResponse, error) {
	users, _ := handler.UserService.GetFollowers(request.ID)
	userResponse := make([]*followers.User, len(*users))

	if users != nil && len(*users) > 0 {
		for i, user := range *users {
			userResponse[i] = &followers.User{
				ID:             user.ID,
				Username:       user.Username,
				Password:       user.Password,
				IsActive:       user.IsActive,
				ProfilePicture: user.ProfilePicture,
				Role:           followers.User_Role(user.Role),
			}
		}
	}

	ret := &followers.GetFollowersResponse{
		Users: userResponse,
	}

	return ret, nil
}

func (handler UserHandler) GetFollowings(ctx context.Context, request *followers.GetFollowingsRequest) (*followers.GetFollowingsResponse, error) {
	users, _ := handler.UserService.GetFollowings(request.ID)

	userResponse := make([]*followers.User, len(*users))

	if users != nil && len(*users) > 0 {
		for i, user := range *users {
			userResponse[i] = &followers.User{
				ID:             user.ID,
				Username:       user.Username,
				Password:       user.Password,
				IsActive:       user.IsActive,
				ProfilePicture: user.ProfilePicture,
				Role:           followers.User_Role(user.Role),
			}
		}
	}

	ret := &followers.GetFollowingsResponse{
		Users: userResponse,
	}

	return ret, nil
}

func (handler UserHandler) GetRecommendedUsers(ctx context.Context, request *followers.GetRecommendedUsersRequest) (*followers.GetRecommendedUsersResponse, error) {
	users, _ := handler.UserService.GetRecommendedUsers(request.ID)

	userResponse := make([]*followers.User, len(*users))

	if users != nil && len(*users) > 0 {
		for i, user := range *users {
			userResponse[i] = &followers.User{
				ID:             user.ID,
				Username:       user.Username,
				Password:       user.Password,
				IsActive:       user.IsActive,
				ProfilePicture: user.ProfilePicture,
				Role:           followers.User_Role(user.Role),
			}
		}
	}

	ret := &followers.GetRecommendedUsersResponse{
		Users: userResponse,
	}

	return ret, nil
}

func (handler UserHandler) Follow(ctx context.Context, request *followers.FollowRequest) (*followers.FollowResponse, error) {
	handler.UserService.Follow(request.FollowingId, request.FollowerId)
	return &followers.FollowResponse{}, nil
}

func (handler UserHandler) Unfollow(ctx context.Context, request *followers.UnfollowRequest) (*followers.UnfollowResponse, error) {
	handler.UserService.Unfollow(request.FollowerId, request.FollowingId)
	return &followers.UnfollowResponse{}, nil
}

func (handler UserHandler) Create(ctx context.Context, request *followers.CreateRequest) (*followers.CreateResponse, error) {
	user := model.User{}
	user.ID = request.User.ID
	user.Username = request.User.Username
	user.Password = request.User.Password
	user.IsActive = request.User.IsActive
	user.ProfilePicture = request.User.ProfilePicture
	user.Role = model.UserRole(request.User.Role)

	handler.UserService.Create(&user)

	return &followers.CreateResponse{}, nil
}
