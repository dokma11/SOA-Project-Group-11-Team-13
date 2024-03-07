package service

import (
	"fmt"
	"github.com/google/uuid"
	"tours/model"
	"tours/repo"
)

type TourService struct {
	TourRepository *repo.TourRepository
}

func (service *TourService) GetById(id string) (*model.Tour, error) {
	tour, err := service.TourRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &tour, nil
}

func (service *TourService) GetAll() (*[]model.Tour, error) {
	tours, err := service.TourRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no tours were found"))
	}
	return &tours, nil
}

func (service *TourService) GetPublished() (*[]model.Tour, error) {
	tours, err := service.TourRepository.GetPublished()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no published tours were found"))
	}
	return &tours, nil
}

func (service *TourService) Create(tour *model.Tour) error {
	err := service.TourRepository.Create(tour)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no tours were created"))
		return err
	}
	return nil
}

func (service *TourService) Delete(id uuid.UUID) error {
	err := service.TourRepository.Delete(id)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no tours were deleted"))
		return err
	}
	return nil
}

func (service *TourService) Update(tour *model.Tour) error {
	err := service.TourRepository.Update(tour)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no tours were updated"))
		return err
	}
	return nil
}
