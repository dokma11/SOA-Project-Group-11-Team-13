package service

import (
	"fmt"
	"tours/model"
	"tours/repo"
)

type ReviewService struct {
	ReviewRepository *repo.ReviewRepository
}

func (service *ReviewService) GetById(id string) (*model.Review, error) {
	review, err := service.ReviewRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("review with id %s not found", id))
	}
	return &review, nil
}

func (service *ReviewService) GetAll() (*[]model.Review, error) {
	reviews, err := service.ReviewRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no tours were found"))
	}
	return &reviews, nil
}

func (service *ReviewService) Create(review *model.Review) error {
	err := service.ReviewRepository.CreateReview(review)
	if err != nil {
		return err
	}
	return nil
}
