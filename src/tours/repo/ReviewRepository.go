package repo

import (
	"gorm.io/gorm"
	"tours/model"
)

type ReviewRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *ReviewRepository) GetById(id string) (model.Review, error) {
	review := model.Review{}
	dbResult := repo.DatabaseConnection.First(&review, "id = ?", id)
	if dbResult != nil {
		return review, dbResult.Error
	}
	return review, nil
}

func (repo *ReviewRepository) GetAll() ([]model.Review, error) {
	var reviews []model.Review
	dbResult := repo.DatabaseConnection.Find(&reviews)
	if dbResult != nil {
		return nil, dbResult.Error
	}
	return reviews, nil
}

func (repo *ReviewRepository) CreateReview(review *model.Review) error {
	dbResult := repo.DatabaseConnection.Create(review)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
