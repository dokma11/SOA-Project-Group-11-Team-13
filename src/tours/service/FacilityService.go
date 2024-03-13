package service

import (
	"fmt"
	"tours/model"
	"tours/repo"
)

type FacilityService struct {
	FacilityRepository *repo.FacilityRepository
}

func (service *FacilityService) GetAll() (*[]model.Facility, error) {
	facilities, err := service.FacilityRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no facilities were found"))
	}
	return &facilities, nil
}

func (service *FacilityService) GetAllByAuthorId(authorId string) (*[]model.Facility, error) {
	facilities, err := service.FacilityRepository.GetAllByAuthorId(authorId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no facilities with given author id: %s were found", authorId))
	}
	return &facilities, nil
}

func (service *FacilityService) Create(facility *model.Facility) error {
	err := service.FacilityRepository.Create(facility)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no facilities were created"))
		return err
	}
	return nil
}

func (service *FacilityService) Delete(id string) error {
	err := service.FacilityRepository.Delete(id)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no facilities were deleted"))
		return err
	}
	return nil
}

func (service *FacilityService) Update(facility *model.Facility) error {
	err := service.FacilityRepository.Update(facility)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no facilities were updated"))
		return err
	}
	return nil
}
