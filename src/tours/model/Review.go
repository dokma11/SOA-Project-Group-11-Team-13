package model

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Review struct {
	ID            uuid.UUID `json:"id"`
	Rating        int       `json:"rating"`
	Comment       string    `json:"comment"`
	TouristId     int       `json:"touristId"`
	TourId        int       `json:"tourId"`
	TourVisitDate time.Time `json:"tourVisitDate"`
	CommentDate   time.Time `json:"commentDate"`
	Images        []string  `json:"images" gorm:"type:varchar(255)[]"`
}

func NewReview(rating int, comment string, touristId int, tourId int, images []string,
	tourVisitDate time.Time, commentDate time.Time) (*Review, error) {

	if images == nil {
		images = []string{}
	}

	review := &Review{
		Rating:        rating,
		Comment:       comment,
		TouristId:     touristId,
		TourId:        tourId,
		Images:        images,
		TourVisitDate: tourVisitDate,
		CommentDate:   commentDate,
	}

	if err := review.Validate(); err != nil {
		return nil, err
	}

	return review, nil
}

func (review *Review) Validate() error {
	if review.Rating < 1 && review.Rating > 5 {
		return errors.New("invalid Rating")
	}
	if review.Comment == "" {
		return errors.New("invalid Comment")
	}
	if len(review.Images) < 1 {
		return errors.New("invalid Images")
	}
	return nil
}

func (review *Review) BeforeCreate(scope *gorm.DB) error {
	review.ID = uuid.New()
	return nil
}
