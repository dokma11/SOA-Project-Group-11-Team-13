package model

import (
	"errors"
	"gorm.io/gorm"
)

type KeyPoint struct {
	gorm.Model
	ID              int64   `json:"Id"`
	TourId          int64   `json:"TourId"`
	Name            string  `json:"Name"`
	Description     string  `json:"Description"`
	Longitude       float64 `json:"Longitude"`
	Latitude        float64 `json:"Latitude"`
	LocationAddress string  `json:"LocationAddress"`
	ImagePath       string  `json:"ImagePath"`
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
