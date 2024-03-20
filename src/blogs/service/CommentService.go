package service

import (
	"blogs/model"
	"blogs/repo"
	"fmt"
)

type CommentService struct {
	CommentRepository *repo.CommentRepository
}

func (service *CommentService) GetById(id string) (*model.Comment, error) {
	comment, err := service.CommentRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("comment with id %s not found", id))
	}
	return &comment, nil
}

func (service *CommentService) GetAll() (*[]model.Comment, error) {
	comments, err := service.CommentRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("no comments were found")
	}
	return &comments, nil
}