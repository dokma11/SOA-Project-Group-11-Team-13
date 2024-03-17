package service

import (
	"blogs/model"
	"blogs/repo"
	"fmt"
)

type BlogService struct {
	BlogRepository *repo.BlogRepository
}

func (service *BlogService) GetById(id string) (*model.Blog, error) {
	blog, err := service.BlogRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("blog with id %s not found", id))
	}
	return &blog, nil
}

func (service *BlogService) GetAll() (*[]model.Blog, error) {
	blogs, err := service.BlogRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no blogs were found"))
	}
	return &blogs, nil
}