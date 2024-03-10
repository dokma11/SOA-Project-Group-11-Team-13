package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type FacilityCategory int

const (
	Restaurant FacilityCategory = iota
	ParkingLot
	Toilet
	Hospital
	Cafe
	Pharmacy
	ExchangeOffice
	BusStop
	Shop
	Other
)

type Facility struct {
	gorm.Model
	ID          int64            `json:"Id"`
	AuthorId    int64            `json:"AuthorId"`
	Name        string           `json:"Name"`
	Description string           `json:"Description"`
	Longitude   float64          `json:"Longitude"`
	Latitude    float64          `json:"Latitude"`
	Category    FacilityCategory `json:"Category"`
	ImagePath   string           `json:"ImagePath"`
}

func NewFacility(id int64, authorId int64, name string, description string, longitude float64, latitude float64,
	category FacilityCategory, imagePath string) (*Facility, error) {
	facility := &Facility{
		ID:          id,
		AuthorId:    authorId,
		Name:        name,
		Description: description,
		Longitude:   longitude,
		Latitude:    latitude,
		Category:    category,
		ImagePath:   imagePath,
	}

	if err := facility.Validate(); err != nil {
		return nil, err
	}

	return facility, nil
}

func (facility *Facility) Validate() error {
	if facility.Name == "" {
		return errors.New("invalid name. Name cannot be empty")
	}
	if facility.Description == "" {
		return errors.New("invalid description. Description cannot be empty")
	}
	if facility.ImagePath == "" {
		return errors.New("invalid image path. Image path cannot be empty")
	}
	if facility.Longitude < -180 || facility.Longitude > 180 {
		return errors.New("invalid Longitude")
	}
	if facility.Latitude < -90 || facility.Latitude > 90 {
		return errors.New("invalid Latitude")
	}
	if facility.Category < 0 || facility.Category > 9 {
		return errors.New("invalid category")
	}
	return nil
}

func (facility *Facility) BeforeCreate(scope *gorm.DB) error {
	if facility.ID == 0 {
		var maxID int64
		if err := scope.Table("facilities").Select("COALESCE(MAX(id), 0)").Row().Scan(&maxID); err != nil {
			return err
		}
		facility.ID = maxID + 1
	}
	return nil
}

func (facility *Facility) String() string {
	return fmt.Sprintf("KeyPoint{ID: %d, AuthorId: %d, Name: %s, Description: %s, "+
		"Longitude: %f, Latitude: %f, ImagePath: %s, Category: %d}",
		facility.ID, facility.AuthorId, facility.Name, facility.Description, facility.Longitude,
		facility.Latitude, facility.ImagePath, facility.Category)
}
