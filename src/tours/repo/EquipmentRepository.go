package repo

import (
	"tours/model"

	"gorm.io/gorm"
)

type EquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EquipmentRepository) GetById(id string) (model.Equipment, error) {
	equipment := model.Equipment{}
	dbResult := repo.DatabaseConnection.First(&equipment, "id = ?", id)
	if dbResult.Error != nil {
		return equipment, dbResult.Error
	}
	return equipment, nil
}

func (repo *EquipmentRepository) GetAll() ([]model.Equipment, error) {
	//println("USAO ZA GET ALL: ")
	var equipment []model.Equipment
	dbResult := repo.DatabaseConnection.Find(&equipment)

	/*for _, e := range equipment {
		//println("Ime: " + e.Name)
	}*/

	if dbResult.Error != nil {
		//println("EROR U EQUIPMENT REPO ZA GET ALL")
		return nil, dbResult.Error
	}

	return equipment, nil
}

func (repo *EquipmentRepository) Create(review *model.Equipment) error {
	//println("USAO ZA CREATE: ")
	dbResult := repo.DatabaseConnection.Create(review)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	//println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
