package service

import (
	"fmt"
	"tours/model"
	"tours/repo"
)

type KeyPointService struct {
	KeyPointRepository *repo.KeyPointRepository
}

func (service *KeyPointService) GetById(id string) (*model.KeyPoint, error) {
	keyPoint, err := service.KeyPointRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &keyPoint, nil
}

func (service *KeyPointService) GetAll() (*[]model.KeyPoint, error) {
	keyPoints, err := service.KeyPointRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no keypoints were found"))
	}
	return &keyPoints, nil
}

func (service *KeyPointService) Create(keyPoint *model.KeyPoint) error {
	err := service.KeyPointRepository.Create(keyPoint)
	if err != nil {
		return err
	}
	return nil
}
