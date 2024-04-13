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

func (service *CommentService) GetByBlogId(id string, page int, pageSize int) ([]model.Comment, int, error) {
	comments, totalCount, err := service.CommentRepository.GetByBlogId(id, page, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("comments with blog id %s not found", id)
	}
	return comments, totalCount, nil
}

func (service *CommentService) GetAll() (*[]model.Comment, error) {
	comments, err := service.CommentRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("no comments were found")
	}
	return &comments, nil
}

func (service *CommentService) Create(comment *model.Comment) error {
	err := service.CommentRepository.Create(comment)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no comments were created"))
		return err
	}
	return nil
}

func (service *CommentService) Delete(id string) error {
	err := service.CommentRepository.Delete(id)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no comments were deleted"))
		return err
	}
	return nil
}

func (service *CommentService) Update(comment *model.Comment) error {
	err := service.CommentRepository.Update(comment)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no keypoints were updated"))
		return err
	}
	return nil
}
