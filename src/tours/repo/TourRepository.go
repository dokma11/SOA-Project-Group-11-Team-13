package repo

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hash/fnv"
	"math"
	"math/big"
	"tours/dto"
	"tours/model"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRepository) GetById(id string) (model.Tour, error) {
	tour := model.Tour{}
	dbResult := repo.DatabaseConnection.First(&tour, "id = ?", id)
	if dbResult != nil {
		return tour, dbResult.Error
	}
	return tour, nil
}

func (repo *TourRepository) GetByAuthorId(authorId string) ([]dto.TourResponseDto, error) {
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.Where("author_id = ?", authorId).Find(&tours)

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
	dbResult := repo.DatabaseConnection.Where("status = ?", "Published").Find(&tours)
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
