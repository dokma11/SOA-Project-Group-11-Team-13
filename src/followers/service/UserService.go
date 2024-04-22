package service

import (
	"fmt"
	"followers/model"
	"followers/repo"
	"log"
)

type UserService struct {
	UserRepository *repo.UserRepository
	logger         *log.Logger
}

func NewUserService(l *log.Logger, r *repo.UserRepository) *UserService {
	return &UserService{r, l}
}

func (service *UserService) Create(user *model.User) error {
	err := service.UserRepository.Create(user)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no users were created"))
		return err
	}
	return nil
}

func (service *UserService) Unfollow(userId string, followingId string) error {
	err := service.UserRepository.DeleteFollowConnectionBetweenUsers(userId, followingId)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no users were unfollowed"))
		return err
	}
	return nil
}

func (service *UserService) Follow(userId string, followedById string) error {
	err := service.UserRepository.CreateFollowConnectionBetweenUsers(userId, followedById)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no follow connections were created"))
		return err
	}
	return nil
}

func (service *UserService) GetFollowers(userId string) (*[]model.User, error) {
	followers, err := service.UserRepository.GetFollowers(userId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no users with given id: %s were found", userId))
	}
	return &followers, nil
}

func (service *UserService) GetFollowings(userId string) (*[]model.User, error) {
	followings, err := service.UserRepository.GetFollowings(userId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no users with given id: %s were found", userId))
	}
	return &followings, nil
}

func (service *UserService) GetByUsername(username string) (*model.User, error) {
	user, err := service.UserRepository.GetByUsername(username)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no users with given username: %s were found", username))
	}
	return &user, nil
}

func (service *UserService) GetRecommendedUsers(userId string) (*[]model.User, error) {
	users, err := service.UserRepository.GetRecommendedUsers(userId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no users with given id: %s were found", userId))
	}
	return &users, nil
}
