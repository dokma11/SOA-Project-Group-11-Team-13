package model

import (
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
	Status BlogStatus
	AuthorId int
	// ClubId int - verovatno ne mora
	Comments []Comment `gorm:"foreignKey:BlogId"`
	Votes []Vote `gorm:"foreignKey:BlogId"`
	Recommendations []BlogRecommendation
	// VisibilityPolicy BlogVisibilityPolicy - mozda necemo morati komplikovati
}

func (blog *Blog) BeforeCreate(scope *gorm.DB) error {
	if blog.ID == 0 {
		var maxID int
		if err := scope.Table("blogs").Select("COALESCE(MAX(id), 0)").Row().Scan(&maxID); err != nil {
			return err
		}
		blog.ID = maxID + 1
	}
	return nil
}
