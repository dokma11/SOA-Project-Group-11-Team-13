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

func (comment *Comment) BeforeCreate(scope *gorm.DB) error {
	if comment.ID == 0 {
		var maxID int
		if err := scope.Table("comments").Select("COALESCE(MAX(id), 0)").Row().Scan(&maxID); err != nil {
			return err
		}
		comment.ID = maxID + 1
	}
	return nil
}