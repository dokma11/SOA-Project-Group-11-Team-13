package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID int		
	AuthorId int
	BlogId int
	Text string
	CreatedAt time.Time
	UpdatedAt time.Time
}