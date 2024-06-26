package repo

import (
	"errors"
	"gorm.io/gorm"
	"tours/model"
)

type KeyPointRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *KeyPointRepository) GetById(id string) (model.KeyPoint, error) {
	keyPoint := model.KeyPoint{}
	dbResult := repo.DatabaseConnection.First(&keyPoint, "id = ?", id)
	if dbResult.Error != nil {
		return keyPoint, dbResult.Error
	}
	return keyPoint, nil
}

func (repo *KeyPointRepository) GetAllByTourId(tourId string) ([]model.KeyPoint, error) {
	var keyPoints []model.KeyPoint
	dbResult := repo.DatabaseConnection.Find(&keyPoints, "tour_id = ?", tourId)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return keyPoints, nil
}

func (repo *KeyPointRepository) GetAll() ([]model.KeyPoint, error) {
	var keyPoints []model.KeyPoint
	dbResult := repo.DatabaseConnection.Find(&keyPoints)
	if dbResult.Error != nil {
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

func (repo *KeyPointRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Where("id = ?", id).Delete(&model.KeyPoint{})
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("no key point found for deletion")
	}
	return nil
}

func (repo *KeyPointRepository) Update(keyPoint *model.KeyPoint) error {
	dbResult := repo.DatabaseConnection.Model(&model.KeyPoint{}).Where("id = ?", keyPoint.ID).Updates(keyPoint)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("no key point found for update")
	}
	return nil
}
