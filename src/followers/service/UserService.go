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
