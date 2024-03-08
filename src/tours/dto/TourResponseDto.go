package dto

import (
	"github.com/lib/pq"
	"time"
	"tours/model"
)

type TourStatus int

const (
	Draft TourStatus = iota
	Archived
	Published
	Ready
)

type TourCategory int

const (
	Adventure TourCategory = iota
	FamilyTrips
	Cruise
	Cultural
)

type TourResponseDto struct {
	ID            int64                `json:"Id"`
	AuthorId      int                  `json:"AuthorId" gorm:"not null;type:int"`
	Name          string               `json:"Name" gorm:"not null;type:string"`
	Description   string               `json:"Description" gorm:"not null;type:string"`
	Difficulty    int                  `json:"Difficulty" gorm:"not null;type:int"`
	Tags          pq.StringArray       `json:"Tags" gorm:"type:text[]"`
	Status        TourStatus           `json:"Status"`
	Price         float64              `json:"Price"`
	Distance      float64              `json:"Distance"`
	AverageRating float64              `json:"AverageRating"`
	PublishDate   time.Time            `json:"PublishDate"`
	ArchiveDate   time.Time            `json:"ArchiveDate"`
	Category      TourCategory         `json:"Category"`
	KeyPoints     []model.KeyPoint     `gorm:"foreignKey:TourId"`
	Durations     []model.TourDuration `json:"Durations" gorm:"type:jsonb"`
	IsDeleted     bool                 `json:"IsDeleted"`
}
