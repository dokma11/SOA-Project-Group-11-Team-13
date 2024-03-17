package model

import (
	"time"

	"gorm.io/gorm"
)

type BlogStatus int

const (
	Draft BlogStatus = iota
	Published
	Closed
	Active
	Famous
)

type Blog struct {
	gorm.Model
	ID int		
	Title string
	Description string
	Date time.Time
	Status BlogStatus
	AuthorId int
	ClubId int
	Comments []Comment `gorm:"foreignKey:BlogId"`
	Votes []Vote `gorm:"foreignKey:BlogId"`
	// VisibilityPolicy BlogVisibilityPolicy - mozda necemo morati komplikovati
}