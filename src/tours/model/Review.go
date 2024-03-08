package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Review struct {
	gorm.Model
	ID            int64     `json:"id"`
	Rating        int       `json:"rating"`
	Comment       string    `json:"comment"`
	TouristId     int       `json:"touristId"`
	TourId        int64     `json:"tourId"`
	TourVisitDate time.Time `json:"tourVisitDate"`
	CommentDate   time.Time `json:"commentDate"`
	Images        []string  `json:"images" gorm:"type:varchar(255)[]"`
}

func NewReview(rating int, comment string, touristId int, tourId int64, images []string,
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
	if review.ID == 0 {
		var maxID int64
		if err := scope.Table("reviews").Select("COALESCE(MAX(id), 0)").Row().Scan(&maxID); err != nil {
			return err
		}
		review.ID = maxID + 1
	}
	return nil
}
