package service

import (
	"fmt"
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

func (service *TourService) Create(tour *model.Tour) error {
	err := service.TourRepository.CreateTour(tour)
	if err != nil {
		return err
	}
	return nil
}
