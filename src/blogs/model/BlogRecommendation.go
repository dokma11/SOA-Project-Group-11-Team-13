package model

import (
	"gorm.io/gorm"
)

type BlogRecommendation struct {
	gorm.Model
	ID int
	BlogId int
	RecommenderId int
	RecommendationReceiverId int
	Blog Blog
}

func (blogRecommendation *BlogRecommendation) BeforeCreate(scope *gorm.DB) error {
	if blogRecommendation.ID == 0 {
		var maxID int
		if err := scope.Table("blog_recommendations").Select("COALESCE(MAX(id), 0)").Row().Scan(&maxID); err != nil {
			return err
		}
		blogRecommendation.ID = maxID + 1
	}
	return nil
}
