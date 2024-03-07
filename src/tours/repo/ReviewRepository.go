package repo

import (
	"errors"
	"github.com/google/uuid"
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

func (repo *ReviewRepository) Create(review *model.Review) error {
	dbResult := repo.DatabaseConnection.Create(review)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *ReviewRepository) Delete(id uuid.UUID) error {
	dbResult := repo.DatabaseConnection.Where("id = ?", id).Delete(&model.Review{})
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("no review found for deletion")
	}
	return nil
}

func (repo *ReviewRepository) Update(review *model.Review) error {
	dbResult := repo.DatabaseConnection.Model(&model.Review{}).Where("id = ?", review.ID).Updates(review)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("no review found for update")
	}
	return nil
}
