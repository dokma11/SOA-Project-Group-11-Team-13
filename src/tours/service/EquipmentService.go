package service

import (
	"fmt"
	"tours/model"
	"tours/repo"
)

type EquipmentService struct {
	EquipmentRepository *repo.EquipmentRepository
}

func (service *EquipmentService) GetById(id string) (*model.Equipment, error) {
	equipment, err := service.EquipmentRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("equipment with id %s not found", id))
	}
	return &equipment, nil
}

func (service *EquipmentService) GetAll() (*[]model.Equipment, error) {
	equipment, err := service.EquipmentRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no equipment was found"))
	}
	return &equipment, nil
}

func (service *EquipmentService) Create(equipment *model.Equipment) error {
	err := service.EquipmentRepository.Create(equipment)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no equipment was created"))
		return err
	}
	return nil
}
