package model

import (
	"errors"
	"gorm.io/gorm"
)

type Equipment struct {
	gorm.Model
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tours       []Tour `gorm:"many2many:tour_equipment;"`
}

func NewEquipment(name string, description string) (*Equipment, error) {
	equipment := &Equipment{
		Name:        name,
		Description: description,
	}

	if err := equipment.Validate(); err != nil {
		return nil, err
	}

	return equipment, nil
}

func (equipment *Equipment) Validate() error {
	if equipment.Name == "" {
		return errors.New("invalid Equipment Name")
	}
	if equipment.Description == "" {
		return errors.New("invalid Equipment Description")
	}
	return nil
}

func (equipment *Equipment) BeforeCreate(scope *gorm.DB) error {
	if equipment.ID == 0 {
		var maxID int64
		if err := scope.Table("equipment").Select("COALESCE(MAX(id), 0)").Row().Scan(&maxID); err != nil {
			return err
		}
		equipment.ID = maxID + 1
	}
	return nil
}
