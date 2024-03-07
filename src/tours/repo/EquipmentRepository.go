package repo

import (
	"gorm.io/gorm"
	"tours/model"
)

type EquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EquipmentRepository) GetById(id string) (model.Equipment, error) {
	equipment := model.Equipment{}
	dbResult := repo.DatabaseConnection.First(&equipment, "id = ?", id)
	if dbResult != nil {
		return equipment, dbResult.Error
	}
	return equipment, nil
}

func (repo *EquipmentRepository) GetAll() ([]model.Equipment, error) {
	var equipment []model.Equipment
	dbResult := repo.DatabaseConnection.Find(&equipment)
	if dbResult != nil {
		return nil, dbResult.Error
	}
	return equipment, nil
}

func (repo *EquipmentRepository) Create(review *model.Equipment) error {
	dbResult := repo.DatabaseConnection.Create(review)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
