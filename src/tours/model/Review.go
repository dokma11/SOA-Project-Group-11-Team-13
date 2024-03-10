package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Review struct {
	gorm.Model
	ID            int64     `json:"id"`
	Rating        int       `json:"rating" gorm:"not null;type:int"`
	Comment       string    `json:"comment" gorm:"not null;type:string"`
	TouristId     int       `json:"touristId" gorm:"not null;type:int"`
	TourId        int64     `json:"tourId" gorm:"not null;type:int64"`
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
		return errors.New("invalid Rating. Rating's value range is from 1 to 5")
	}
	if review.Comment == "" {
		return errors.New("invalid Comment. Comment cannot be empty")
	}
	if len(review.Images) < 1 {
		return errors.New("invalid Images. Images cannot be empty")
	}
	if review.TourVisitDate.IsZero() {
		return errors.New("invalid Tour Visit Date. Tour Visit Date cannot be empty")
	}
	if review.CommentDate.IsZero() {
		return errors.New("invalid Comment Date. Comment Date cannot be empty")
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

func (review *Review) String() string {
	imagesStr := strings.Join(review.Images, ", ")
	return fmt.Sprintf("Review{ID: %d, Rating: %d, Comment: %s, TouristId: %d, TourId: %d, "+
		"TourVisitDate: %s, CommentDate: %s, Images: [%s]}",
		review.ID, review.Rating, review.Comment, review.TouristId, review.TourId,
		review.TourVisitDate.Format("2006-01-02"), review.CommentDate.Format("2006-01-02"), imagesStr)
}
