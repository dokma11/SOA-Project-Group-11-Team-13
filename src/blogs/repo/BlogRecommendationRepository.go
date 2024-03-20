package repo

import (
	"blogs/model"

	"gorm.io/gorm"
)

type BlogRecommendationRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogRecommendationRepository) GetById(id string) (model.BlogRecommendation, error) {
	var blogRecommendation model.BlogRecommendation
	dbResult := repo.DatabaseConnection.Where("id = ?", id).First(&blogRecommendation)
	if dbResult.Error != nil {
		return blogRecommendation, dbResult.Error
	}
	return blogRecommendation, nil
}

func (repo *BlogRecommendationRepository) GetAll() ([]model.BlogRecommendation, error) {
	var blogRecommendations []model.BlogRecommendation
	dbResult := repo.DatabaseConnection.Find(&blogRecommendations)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return blogRecommendations, nil
}

func (repo *BlogRecommendationRepository) Save(blogRecommendation *model.BlogRecommendation) error {
	dbResult := repo.DatabaseConnection.Create(blogRecommendation)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
