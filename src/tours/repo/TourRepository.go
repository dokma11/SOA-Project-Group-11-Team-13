package repo

import (
	"encoding/json"
	"errors"
	"hash/fnv"
	"math"
	"math/big"
	"fmt"
	"tours/dto"
	"tours/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}


func (repo *TourRepository) GetById(id string) (dto.TourResponseDto, error) {
	var tour model.Tour
	var tourDto dto.TourResponseDto

	dbResult := repo.DatabaseConnection.
		Table("tours").
		Select("id, tags, status, name, description, publish_date, archive_date, category, is_deleted, price, distance, difficulty, author_id").
		Where("id = ?", id).
		Preload("KeyPoints").
		First(&tour)
	if dbResult.Error != nil {
		return tourDto, dbResult.Error
	}

	returnValue, _ := MapToDto(repo, &tour, &tourDto)
	return *returnValue, nil
}

func (repo *TourRepository) GetByAuthorId(authorId string) ([]dto.TourResponseDto, error) {
	var tours []model.Tour

	dbResult := repo.DatabaseConnection.
		Table("tours").
		Select("id, tags, status, name, description, publish_date, archive_date, category, is_deleted, price, distance, difficulty, author_id").
		Where("author_id = ?", authorId).
		Preload("KeyPoints").
		Find(&tours)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	var tourDtos []dto.TourResponseDto

	for _, tour := range tours {
		var tourDto dto.TourResponseDto
		responseValue, _ := MapToDto(repo, &tour, &tourDto)
		tourDtos = append(tourDtos, *responseValue)
	}

	return tourDtos, nil
}

func (repo *TourRepository) GetAll() ([]dto.TourResponseDto, error) {
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.Preload("KeyPoints").Find(&tours)
	if dbResult != nil {
		return nil, dbResult.Error
	}

	var tourDtos []dto.TourResponseDto

	for _, tour := range tours {
		var tourDto dto.TourResponseDto
		responseValue, _ := MapToDto(repo, &tour, &tourDto)
		tourDtos = append(tourDtos, *responseValue)
	}

	return tourDtos, nil
}

func (repo *TourRepository) GetPublished() ([]dto.TourResponseDto, error) {
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.
		Table("tours").
		Select("id, tags, status, name, description, publish_date, archive_date, category, is_deleted, price, distance, difficulty, author_id").
		Where("status = ?", "Published").
		Preload("KeyPoints").
		Find(&tours)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	var tourDtos []dto.TourResponseDto

	for _, tour := range tours {
		var tourDto dto.TourResponseDto
		responseValue, _ := MapToDto(repo, &tour, &tourDto)
		tourDtos = append(tourDtos, *responseValue)
	}

	return tourDtos, nil
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
	dbResult := repo.DatabaseConnection.Model(&model.Tour{}).
		Where("id = ?", tour.ID).
		Omit("Durations").
		Updates(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("no tour found for update")
	}
	return nil
}

func (repo *TourRepository) AddDurations(tour *model.Tour) error {
	durationsJSON, err := json.Marshal(tour.Durations)
	if err != nil {
		return err
	}

	dbResult := repo.DatabaseConnection.Exec(
		"UPDATE tours SET durations = ? WHERE id = ?",
		string(durationsJSON),
		tour.ID,
	)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("no tour found for duration addition")
	}
	return nil
}

func MapToDto(repo *TourRepository, tour *model.Tour, tourDto *dto.TourResponseDto) (*dto.TourResponseDto, error) {
	tourDto.AverageRating = 0.0
	tourDto.Tags = tour.Tags
	tourDto.KeyPoints = tour.KeyPoints
	tourDto.Status = dto.TourStatus(tour.Status)
	tourDto.Name = tour.Name
	tourDto.Description = tour.Description
	tourDto.ID = tour.ID

	var durationsJSON []byte
	if err := repo.DatabaseConnection.Raw("SELECT durations FROM tours WHERE id = ?", tour.ID).Row().Scan(&durationsJSON); err != nil {
		fmt.Println(fmt.Sprintf("Error: Couldn't get tours durations"))
		return tourDto, err
	}

	var durations []model.TourDuration

	if len(durationsJSON) > 0 {
		if err := json.Unmarshal(durationsJSON, &durations); err != nil {
			fmt.Println(fmt.Sprintf("Error: Couldn't unmarshal tours durations"))
			return tourDto, err
		}
	}

	tourDto.Durations = durations
	tourDto.PublishDate = tour.PublishDate
	tourDto.ArchiveDate = tour.ArchiveDate
	tourDto.Category = dto.TourCategory(tour.Category)
	tourDto.IsDeleted = tour.IsDeleted
	tourDto.Price = tour.Price
	tourDto.Distance = tour.Distance
	tourDto.Difficulty = tour.Difficulty
	tourDto.AuthorId = tour.AuthorId

	return tourDto, nil
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
