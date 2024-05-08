package repo

import (
	"errors"
	"gorm.io/gorm"
	"tours/model"
)

type FacilityRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *FacilityRepository) GetAll() ([]model.Facility, error) {
	var facilities []model.Facility
	dbResult := repo.DatabaseConnection.Find(&facilities)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return facilities, nil
}

func (repo *FacilityRepository) GetAllByAuthorId(authorId string) ([]model.Facility, error) {
	var facilities []model.Facility
	dbResult := repo.DatabaseConnection.Find(&facilities, "author_id = ?", authorId)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return facilities, nil
}

func (repo *FacilityRepository) Create(facility *model.Facility) error {
	dbResult := repo.DatabaseConnection.Create(facility)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *FacilityRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Where("id = ?", id).Delete(&model.Facility{})
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("no facility found for deletion")
	}
	return nil
}

func (repo *FacilityRepository) Update(facility *model.Facility) error {
	dbResult := repo.DatabaseConnection.Model(&model.Facility{}).Where("id = ?", facility.ID).Updates(facility)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("no facility found for update")
	}
	return nil
}
