package repo

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"tours/model"
)

type KeyPointRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *KeyPointRepository) GetById(id string) (model.KeyPoint, error) {
	keyPoint := model.KeyPoint{}
	dbResult := repo.DatabaseConnection.First(&keyPoint, "id = ?", id)
	if dbResult != nil {
		return keyPoint, dbResult.Error
	}
	return keyPoint, nil
}

func (repo *KeyPointRepository) GetAllByTourId(tourId int64) ([]model.KeyPoint, error) {
	var keyPoints []model.KeyPoint
	dbResult := repo.DatabaseConnection.Find(&keyPoints, "tour_id = ?", tourId)
	if dbResult != nil {
		return keyPoints, dbResult.Error
	}
	return keyPoints, nil
}

func (repo *KeyPointRepository) GetAll() ([]model.KeyPoint, error) {
	var keyPoints []model.KeyPoint
	dbResult := repo.DatabaseConnection.Find(&keyPoints)
	if dbResult != nil {
		return nil, dbResult.Error
	}
	return keyPoints, nil
}

func (repo *KeyPointRepository) Create(keyPoint *model.KeyPoint) error {
	dbResult := repo.DatabaseConnection.Create(keyPoint)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *KeyPointRepository) Delete(id uuid.UUID) error {
	dbResult := repo.DatabaseConnection.Where("id = ?", id).Delete(&model.KeyPoint{})
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("no key point found for deletion")
	}
	return nil
}
