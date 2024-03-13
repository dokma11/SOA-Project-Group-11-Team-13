package model

import (
	"errors"
	"fmt"
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
	Order           int64   `json:"Order"`
}

func NewKeyPoint(tourId int64, name, description, locationAddress, imagePath string, longitude, latitude float64,
	order int64) (*KeyPoint, error) {
	keyPoint := &KeyPoint{
		TourId:          tourId,
		Name:            name,
		Description:     description,
		Longitude:       longitude,
		Latitude:        latitude,
		LocationAddress: locationAddress,
		ImagePath:       imagePath,
		Order:           order,
	}

	if err := keyPoint.Validate(); err != nil {
		return nil, err
	}

	return keyPoint, nil
}

func (keyPoint *KeyPoint) Validate() error {
	if keyPoint.Name == "" {
		return errors.New("invalid Name. Name cannot be empty")
	}
	if keyPoint.Description == "" {
		return errors.New("invalid Description. Description cannot be empty")
	}
	if keyPoint.Longitude < -180 || keyPoint.Longitude > 180 {
		return errors.New("invalid Longitude")
	}
	if keyPoint.Latitude < -90 || keyPoint.Latitude > 90 {
		return errors.New("invalid Latitude")
	}
	if keyPoint.LocationAddress == "" {
		return errors.New("invalid Location Address. Location Address cannot be empty")
	}
	if keyPoint.ImagePath == "" {
		return errors.New("invalid Image Path. ImagePath cannot be empty")
	}
	if keyPoint.Order < 0 {
		return errors.New("invalid Order. Order cannot be negative")
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

func (keyPoint *KeyPoint) String() string {
	return fmt.Sprintf("KeyPoint{ID: %d, TourId: %d, Name: %s, Description: %s, "+
		"Longitude: %f, Latitude: %f, LocationAddress: %s, ImagePath: %s, Order: %d}",
		keyPoint.ID, keyPoint.TourId, keyPoint.Name, keyPoint.Description, keyPoint.Longitude,
		keyPoint.Latitude, keyPoint.LocationAddress, keyPoint.ImagePath, keyPoint.Order)
}
