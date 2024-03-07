package repo

import (
	"tours/model"

	"gorm.io/gorm"
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

func (repo *KeyPointRepository) GetAll() ([]model.KeyPoint, error) {
	var keyPoint []model.KeyPoint
	dbResult := repo.DatabaseConnection.Find(&keyPoint)
	if dbResult != nil {
		return nil, dbResult.Error
	}
	return keyPoint, nil
}

func (repo *KeyPointRepository) Create(keyPoint *model.KeyPoint) error {
	dbResult := repo.DatabaseConnection.Create(keyPoint)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
