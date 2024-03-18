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
		return nil, fmt.Errorf("no blogs were found")
	}
	return &blogs, nil
}

func (service *BlogService) Create(blog *model.Blog) error {
	err := service.BlogRepository.Save(blog)
	if err != nil {
		return err
	}
	return nil
}
