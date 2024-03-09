package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"tours/dto"
	"tours/model"
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
		Scan(&tour)
	if dbResult.Error != nil {
		return tourDto, dbResult.Error
	}

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

func (repo *TourRepository) GetByAuthorId(authorId string) ([]dto.TourResponseDto, error) {
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.
		Table("tours").
		Select("id, tags, status, name, description, publish_date, archive_date, category, is_deleted, price, distance, difficulty, author_id").
		Where("author_id = ?", authorId).
		Scan(&tours)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

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

		var durationsJSON []byte
		if err := repo.DatabaseConnection.Raw("SELECT durations FROM tours WHERE id = ?", tour.ID).Row().Scan(&durationsJSON); err != nil {
			return nil, err
		}

		var durations []model.TourDuration

		if len(durationsJSON) > 0 {
			if err := json.Unmarshal(durationsJSON, &durations); err != nil {
				return nil, err
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

		tourDtos = append(tourDtos, tourDto)
	}

	return tourDtos, nil
}

func (repo *TourRepository) GetAll() ([]dto.TourResponseDto, error) {
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.Find(&tours)
	if dbResult != nil {
		return nil, dbResult.Error
	}

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

		var durationsJSON []byte
		if err := repo.DatabaseConnection.Raw("SELECT durations FROM tours WHERE id = ?", tour.ID).Row().Scan(&durationsJSON); err != nil {
			return nil, err
		}

		var durations []model.TourDuration

		if len(durationsJSON) > 0 {
			if err := json.Unmarshal(durationsJSON, &durations); err != nil {
				return nil, err
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

		tourDtos = append(tourDtos, tourDto)
	}

	return tourDtos, nil
}

func (repo *TourRepository) GetPublished() ([]dto.TourResponseDto, error) {
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.
		Table("tours").
		Select("id, tags, status, name, description, publish_date, archive_date, category, is_deleted, price, distance, difficulty, author_id").
		Where("status = ?", "Published").
		Scan(&tours)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

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

		var durationsJSON []byte
		if err := repo.DatabaseConnection.Raw("SELECT durations FROM tours WHERE id = ?", tour.ID).Row().Scan(&durationsJSON); err != nil {
			return nil, err
		}

		var durations []model.TourDuration

		if len(durationsJSON) > 0 {
			if err := json.Unmarshal(durationsJSON, &durations); err != nil {
				return nil, err
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

		tourDtos = append(tourDtos, tourDto)
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
	if len(tour.Durations) > 0 {
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
			return errors.New("no tour found for update")
		}
		return nil
	} else {
		dbResult := repo.DatabaseConnection.Model(&model.Tour{}).Where("id = ?", tour.ID).Updates(tour)
		if dbResult.Error != nil {
			return dbResult.Error
		}
		if dbResult.RowsAffected == 0 {
			return errors.New("no tour found for update")
		}
		return nil
	}
}
