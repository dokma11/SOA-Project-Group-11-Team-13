package model

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TourStatus string

const (
	Draft     TourStatus = "Draft"
	Archived  TourStatus = "Archived"
	Published TourStatus = "Published"
	Ready     TourStatus = "Ready"
)

type Tour struct {
	ID          uuid.UUID  `json:"id"`
	AuthorId    int        `json:"authorId"`
	Name        string     `json:"name" gorm:"not null;type:string"`
	Description string     `json:"description"`
	Difficulty  int        `json:"difficulty"`
	Tags        []string   `json:"tags" gorm:"type:varchar(255)[]"`
	Status      TourStatus `json:"status"`
	Price       float64    `json:"price"`
	Distance    float64    `json:"distance"`
	PublishDate time.Time  `json:"publishDate"`
	ArchiveDate time.Time  `json:"archiveDate"`
	//Equipments []Equipment
	//KeyPoints  []KeyPoint
	Reviews []Review
	//Durations  []TourDuration
}

func NewTour(authorID int, name, description string, tags []string, difficulty int, archiveDate,
	publishDate time.Time, distance float64, status TourStatus, price float64) (*Tour, error) {

	if tags == nil {
		tags = []string{}
	}

	tour := &Tour{
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

	return nil
}

func (tour *Tour) BeforeCreate(scope *gorm.DB) error {
	tour.ID = uuid.New()
	return nil
}
