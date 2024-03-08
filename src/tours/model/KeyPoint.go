package model

import (
	"errors"
	"gorm.io/gorm"
)

type KeyPoint struct {
	gorm.Model
	ID              int64   `json:"id"`
	TourId          int64   `json:"tourId"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Longitude       float64 `json:"longitude"`
	Latitude        float64 `json:"latitude"`
	LocationAddress string  `json:"locationAddress"`
	ImagePath       string  `json:"imagePath"`
}

func NewKeyPoint(tourId int64, name, description, locationAddress, imagePath string, longitude, latitude float64) (*KeyPoint, error) {
	keyPoint := &KeyPoint{
		TourId:          tourId,
		Name:            name,
		Description:     description,
		Longitude:       longitude,
		Latitude:        latitude,
		LocationAddress: locationAddress,
		ImagePath:       imagePath,
	}

	if err := keyPoint.Validate(); err != nil {
		return nil, err
	}

	return keyPoint, nil
}

func (keyPoint *KeyPoint) Validate() error {
	/*if keyPoint.TourId == 0 {
		return errors.New("invalid TourId")
	}*/
	if keyPoint.Name == "" {
		return errors.New("invalid Name")
	}
	if keyPoint.Description == "" {
		return errors.New("invalid Description")
	}
	if keyPoint.Longitude < -180 || keyPoint.Longitude > 180 {
		return errors.New("invalid Longitude")
	}
	if keyPoint.Latitude < -90 || keyPoint.Latitude > 90 {
		return errors.New("invalid Latitude")
	}
	if keyPoint.LocationAddress == "" {
		return errors.New("invalid Location Address")
	}
	if keyPoint.ImagePath == "" {
		return errors.New("invalid ImagePath")
	}

	return nil
}

func (keyPoint *KeyPoint) BeforeCreate(scope *gorm.DB) error {
	if keyPoint.ID == 0 {
		var maxID int64
		if err := scope.Table("key_points").Select("COALESCE(MAX(id), 0)").Row().Scan(&maxID); err != nil {
			return err
		}
		keyPoint.ID = maxID + 1
	}
	return nil
}
