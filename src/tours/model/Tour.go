package model

import (
	"errors"
	"fmt"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type TourStatus int

const (
	Draft TourStatus = iota
	Published
	Archived
	Ready
)

type TourCategory int

const (
	Adventure TourCategory = iota
	FamilyTrips
	Cruise
	Cultural
)

type Tour struct {
	gorm.Model
	ID          int64          `json:"Id"`
	AuthorId    int            `json:"AuthorId" gorm:"not null;type:int"`
	Name        string         `json:"Name" gorm:"not null;type:string"`
	Description string         `json:"Description" gorm:"not null;type:string"`
	Difficulty  int            `json:"Difficulty" gorm:"not null;type:int"`
	Tags        pq.StringArray `json:"Tags" gorm:"type:text[]"`
	Status      TourStatus     `json:"Status"`
	Price       float64        `json:"Price"`
	Distance    float64        `json:"Distance"`
	PublishDate time.Time      `json:"PublishDate"`
	ArchiveDate time.Time      `json:"ArchiveDate"`
	Category    TourCategory   `json:"Category"`
	IsDeleted   bool           `json:"IsDeleted"`
	KeyPoints   []KeyPoint     `gorm:"foreignKey:TourId"`
	Equipment   []Equipment    `gorm:"many2many:tour_equipment;"`
	Reviews     []Review       `gorm:"foreignKey:TourId"`
	Durations   []TourDuration `json:"Durations" gorm:"type:jsonb"`
}

func NewTour(id int64, authorID int, name string, description string, tags pq.StringArray, difficulty int, archiveDate time.Time,
	publishDate time.Time, distance float64, status TourStatus, price float64, category TourCategory,
	isDeleted bool, keyPoints []KeyPoint, durations []TourDuration) (*Tour, error) {

	if tags == nil {
		tags = []string{}
	}

	tour := &Tour{
		ID:          id,
		AuthorId:    authorID,
		Name:        name,
		Description: description,
		Difficulty:  difficulty,
		Tags:        tags,
		Status:      status,
		Price:       price,
		Distance:    distance,
		PublishDate: publishDate,
		ArchiveDate: archiveDate,
		Category:    category,
		IsDeleted:   isDeleted,
		KeyPoints:   keyPoints,
		Durations:   durations,
	}

	if err := tour.Validate(); err != nil {
		return nil, err
	}

	return tour, nil
}

func (tour *Tour) Validate() error {
	if tour.Name == "" {
		return errors.New("invalid Name")
	}
	if tour.Description == "" {
		return errors.New("invalid Description")
	}
	if tour.Difficulty < 1 || tour.Difficulty > 5 {
		return errors.New("invalid Difficulty")
	}
	if len(tour.Tags) == 0 {
		return errors.New("tags cannot be empty")
	}
	if tour.Price < 0 {
		return errors.New("price cannot be negative")
	}
	if tour.Distance < 0 {
		return errors.New("distance cannot be negative")
	}
	if tour.Category < 0 || tour.Category > 3 {
		return errors.New("invalid Category")
	}
	if tour.Status < 0 || tour.Status > 3 {
		return errors.New("invalid Status")
	}

	return nil
}

func (tour *Tour) BeforeCreate(scope *gorm.DB) error {
	if tour.ID == 0 {
		var maxID int64
		if err := scope.Table("tours").Select("COALESCE(MAX(id), 0)").Row().Scan(&maxID); err != nil {
			return err
		}
		tour.ID = maxID + 1
	}

	tour.Equipment = []Equipment{}
	tour.Durations = []TourDuration{}
	tour.KeyPoints = []KeyPoint{}
	tour.Reviews = []Review{}
	return nil
}

func (tour *Tour) String() string {
	return fmt.Sprintf("Tour{ID: %d, AuthorId: %d, Name: %s, Description: %s, Difficulty: %d, Tags: %v, "+
		"Status: %d, Price: %.2f, Distance: %.2f, PublishDate: %s, ArchiveDate: %s, Category: %d, "+
		"IsDeleted: %t, KeyPoints: %v, Equipment: %v, Reviews: %v, Durations: %v}",
		tour.ID, tour.AuthorId, tour.Name, tour.Description, tour.Difficulty, tour.Tags,
		tour.Status, tour.Price, tour.Distance, tour.PublishDate, tour.ArchiveDate, tour.Category,
		tour.IsDeleted, tour.KeyPoints, tour.Equipment, tour.Reviews, tour.Durations)
}
