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
		return nil, fmt.Errorf(fmt.Sprintf("no reviews were found"))
	}
	return &reviews, nil
}

func (service *ReviewService) Create(review *model.Review) error {
	err := service.ReviewRepository.Create(review)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no reviews were created"))
		return err
	}
	return nil
}

func (service *ReviewService) Delete(id string) error {
	err := service.ReviewRepository.Delete(id)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no reviews were deleted"))
		return err
	}
	return nil
}

func (service *ReviewService) Update(review *model.Review) error {
	err := service.ReviewRepository.Update(review)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no reviews were updated"))
		return err
	}
	return nil
}
