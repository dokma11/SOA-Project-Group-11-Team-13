package repo

import (
	"errors"
	"hash/fnv"
	"math"
	"math/big"
	"tours/dto"
	"tours/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRepository) GetById(id string) (model.Tour, error) {
	tour := model.Tour{}
	dbResult := repo.DatabaseConnection.Preload("KeyPoints").First(&tour, "id = ?", id)
	if dbResult != nil {
		return tour, dbResult.Error
	}
	return tour, nil
}

func (repo *TourRepository) GetByAuthorId(authorId string) ([]dto.TourResponseDto, error) {
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.Where("author_id = ?", authorId).Preload("KeyPoints").Find(&tours)

	var tourDtos []dto.TourResponseDto

	for _, tour := range tours {
		var tourDto dto.TourResponseDto
		tourDto.AverageRating = 0.0
		tourDto.Tags = tour.Tags
		tourDto.KeyPoints = tour.KeyPoints
		tourDto.Status = dto.TourStatus(tour.Status)
		tourDto.Name = tour.Name
		tourDto.Description = tour.Description

		tourDto.ID = tour.ID

		tourDto.Durations = tour.Durations
		tourDto.PublishDate = tour.PublishDate
		tourDto.ArchiveDate = tour.ArchiveDate
		tourDto.Category = dto.TourCategory(tour.Category)
		tourDto.IsDeleted = tour.IsDeleted
		tourDto.Price = tour.Price
		tourDto.Distance = tour.Distance
		tourDto.Difficulty = tour.Difficulty
		tourDto.AuthorId = tour.AuthorId

		tourDtos = append(tourDtos, tourDto)
	}

	if dbResult != nil {
		return tourDtos, dbResult.Error
	}
	return tourDtos, nil
}

func (repo *TourRepository) GetAll() ([]model.Tour, error) {
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.Find(&tours)
	if dbResult != nil {
		return nil, dbResult.Error
	}
	return tours, nil
}

func (repo *TourRepository) GetPublished() ([]model.Tour, error) {
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.Where("status = ?", "Published").Preload("KeyPoints").Find(&tours)
	if dbResult != nil {
		return nil, dbResult.Error
	}
	return tours, nil
}

func (repo *TourRepository) Create(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Create(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Where("id = ?", id).Delete(&model.Tour{})
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("no tour found for deletion")
	}
	return nil
}

func (repo *TourRepository) Update(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Model(&model.Tour{}).Where("id = ?", tour.ID).Updates(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("no tour found for update")
	}
	return nil
}

func uuidToInt64(u uuid.UUID) int64 {
	hash := fnv.New64a()
	hash.Write(u[:])

	hashValue := big.NewInt(0)
	hashValue.SetUint64(hash.Sum64())

	maxInt64 := big.NewInt(math.MaxInt64)
	hashValue.Mod(hashValue, maxInt64)
	int64Value := hashValue.Int64()

	return int64Value
}

func (repo *TourRepository) GetEquipment(tourId string) ([]model.Equipment, error) {
	/*println("USAO ZA GET equipmnet ture: ")
	var equipmentList []model.Equipment
	dbResult := repo.DatabaseConnection.Model(&model.Tour{}).
		Where("id = ?", tourId).
		Find(&equipmentList)

	if len(equipmentList) == 0 {
		return equipmentList, nil
	}
	if dbResult.Error != nil {
		println("EROR U TOUR REPO ZA GET EQUIPMENT")
		return nil, dbResult.Error
	}
	return equipmentList, nil*/
	//println("USAO ZA GET EQUIPMENT TOUR: ")

	var tour model.Tour
	dbResult := repo.DatabaseConnection.Preload("Equipment").First(&tour, tourId)

	if dbResult.Error != nil {
		//println("ERROR IN TOUR REPO FOR GET EQUIPMENT")
		return nil, dbResult.Error
	}

	return tour.Equipment, nil
}

func (repo *TourRepository) AddEquipment(tourId string, equipmentId string) error {
	//println("USAO ZA ADD equip: ")
	var tour model.Tour
	var equipment model.Equipment

	// Load the tour and equipment entities
	if err := repo.DatabaseConnection.Preload("Equipment").First(&tour, "id = ?", tourId).Error; err != nil {
		return err
	}
	if err := repo.DatabaseConnection.First(&equipment, "id = ?", equipmentId).Error; err != nil {
		return err
	}

	// Add equipment to the tour
	tour.Equipment = append(tour.Equipment, equipment)

	// Save the changes
	dbResult := repo.DatabaseConnection.Save(&tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *TourRepository) DeleteEquipment(tourId string, equipmentId string) error {
	println("USAO ZA remove equip: ")
	var tour model.Tour
	var equipment model.Equipment

	// Load the tour and equipment entities
	if err := repo.DatabaseConnection.Preload("Equipment").First(&tour, "id = ?", tourId).Error; err != nil {
		return err
	}
	if err := repo.DatabaseConnection.First(&equipment, "id = ?", equipmentId).Error; err != nil {
		return err
	}

	// Remove equipment from the tour
	var updatedEquipment []model.Equipment
	for _, e := range tour.Equipment {
		if e.ID != equipment.ID {
			updatedEquipment = append(updatedEquipment, e)
		}
	}
	tour.Equipment = updatedEquipment

	// Delete the association
	repo.DatabaseConnection.Model(&tour).Association("Equipment").Delete(&equipment)

	// Save the changes
	dbResult := repo.DatabaseConnection.Save(&tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}
