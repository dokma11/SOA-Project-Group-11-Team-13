package service

import (
	"blogs/model"
	"blogs/repo"
	"fmt"
)

type BlogRecommendationService struct {
	BlogRecommendationRepository *repo.BlogRecommendationRepository
}

func (service *BlogRecommendationService) GetById(id string) (*model.BlogRecommendation, error) {
	blogRecommendation, err := service.BlogRecommendationRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("blog recommendation with id %s not found", id))
	}
	return &blogRecommendation, nil
}

func (service *BlogRecommendationService) GetAll() (*[]model.BlogRecommendation, error) {
	blogRecommendations, err := service.BlogRecommendationRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("no blog recommendations were found")
	}
	return &blogRecommendations, nil
}

func (service *BlogRecommendationService) Create(blogRecommendation *model.BlogRecommendation) error {
	err := service.BlogRecommendationRepository.Save(blogRecommendation)
	if err != nil {
		return err
	}
	return nil
}

func (service *BlogRecommendationService) GetByReceiverId(receiverId int) (*[]model.BlogRecommendation, error) {
	blogRecommendations, _ := service.BlogRecommendationRepository.GetByReceiverId(receiverId)
	return &blogRecommendations, nil
}