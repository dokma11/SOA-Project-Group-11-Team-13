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

func (service *KeyPointService) GetAllByTourId(tourId string) (*[]model.KeyPoint, error) {
	keyPoints, err := service.KeyPointRepository.GetAllByTourId(tourId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no keypoints with given tour id: %s were found", tourId))
	}
	return &keyPoints, nil
}

func (service *KeyPointService) Create(keyPoint *model.KeyPoint) error {
	err := service.KeyPointRepository.Create(keyPoint)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no keypoints were created"))
		return err
	}
	return nil
}

func (service *KeyPointService) Delete(id string) error {
	err := service.KeyPointRepository.Delete(id)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no keypoints were deleted"))
		return err
	}
	return nil
}

func (service *KeyPointService) Update(keyPoint *model.KeyPoint) error {
	err := service.KeyPointRepository.Update(keyPoint)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no keypoints were updated"))
		return err
	}
	return nil
}
