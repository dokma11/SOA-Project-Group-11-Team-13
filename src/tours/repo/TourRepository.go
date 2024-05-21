package repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/gorm"
	"log"
	"tours/model"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRepository) GetById(id string) (model.Tour, error) {
	log.Printf("Get tour by id repository call\n")
	var tour model.Tour
	dbResult := repo.DatabaseConnection.
		Preload("KeyPoints").
		Preload("Equipment").
		Where("id = ?", id).
		Omit("Durations").
		First(&tour)
	if dbResult.Error != nil {
		return tour, dbResult.Error
	}
	return tour, nil
}

func (repo *TourRepository) GetByAuthorId(authorId string, tp *trace.TracerProvider, ctx context.Context) ([]model.Tour, error) {
	log.Printf("Get tour by author id repository call\n")
	_, span := tp.Tracer("tours").Start(ctx, "tours-repository-getByAuthorId")
	defer func() { span.End() }()

	var tours []model.Tour
	dbResult := repo.DatabaseConnection.
		Preload("KeyPoints").
		Preload("Equipment").
		Where("author_id = ?", authorId).
		Omit("Durations").
		Find(&tours)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return tours, nil
}

func (repo *TourRepository) GetAll() ([]model.Tour, error) {
	log.Printf("Get all tours repository call\n")
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.
		Preload("KeyPoints").
		Preload("Equipment").
		Find(&tours)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return tours, nil
}

func (repo *TourRepository) GetPublished() ([]model.Tour, error) {
	log.Printf("Get published tours repository call\n")
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.
		Preload("KeyPoints").
		Preload("Equipment").
		Where("status = ?", int64(model.Published)).
		Omit("Durations").
		Find(&tours)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return tours, nil
}

func (repo *TourRepository) Create(tour *model.Tour, tp *trace.TracerProvider, ctx context.Context) error {
	log.Printf("Create tour repository call\n")
	_, span := tp.Tracer("tours").Start(ctx, "tours-repository-create")
	defer func() { span.End() }()

	dbResult := repo.DatabaseConnection.Create(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) Delete(id string) error {
	log.Printf("Delete tour repository call\n")
	var tour model.Tour
	result := repo.DatabaseConnection.
		Preload("KeyPoints").
		Preload("Equipment").
		Where("id = ?", id).
		Omit("Durations").
		First(&tour)
	if result.Error != nil {
		return result.Error
	}

	dbResult := repo.DatabaseConnection.Delete(&tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("no tour found for deletion")
	}

	for _, keyPoint := range tour.KeyPoints {
		dbResult := repo.DatabaseConnection.Delete(&keyPoint)
		if dbResult.Error != nil {
			return dbResult.Error
		}
	}

	for _, equipment := range tour.Equipment {
		err := repo.DatabaseConnection.
			Model(&tour).
			Association("Equipment").
			Delete(&equipment)
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *TourRepository) Update(tour *model.Tour) error {
	log.Printf("Update tour repository call\n")
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
		"UPDATE tours SET durations = ?, distance = ? WHERE id = ?",
		string(durationsJSON),
		tour.Distance,
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

func (repo *TourRepository) GetEquipment(tourId string) ([]model.Equipment, error) {
	var tour model.Tour

	dbResult := repo.DatabaseConnection.
		Preload("Equipment").
		Where("id = ?", tourId).
		Omit("Durations").
		First(&tour)

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return tour.Equipment, nil
}

func (repo *TourRepository) AddEquipment(tourId string, equipmentId string) error {
	var tour model.Tour
	var equipment model.Equipment

	dbResult := repo.DatabaseConnection.
		Where("id = ?", tourId).
		Omit("Durations").
		First(&tour)

	if err := repo.DatabaseConnection.First(&equipment, "id = ?", equipmentId).Error; err != nil {
		return err
	}

	tour.Equipment = append(tour.Equipment, equipment)

	dbResult = repo.DatabaseConnection.Save(&tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *TourRepository) DeleteEquipment(tourId string, equipmentId string) error {
	var tour model.Tour
	var equipment model.Equipment

	dbResult := repo.DatabaseConnection.
		Where("id = ?", tourId).
		Preload("Equipment").
		Omit("Durations").
		First(&tour)

	if err := repo.DatabaseConnection.First(&equipment, "id = ?", equipmentId).Error; err != nil {
		return err
	}

	var updatedEquipment []model.Equipment
	for _, e := range tour.Equipment {
		if e.ID != equipment.ID {
			updatedEquipment = append(updatedEquipment, e)
		}
	}
	tour.Equipment = updatedEquipment

	err := repo.DatabaseConnection.
		Model(&tour).
		Association("Equipment").
		Delete(&equipment)
	if err != nil {
		return errors.New("error while deleting tours equipment")
	}

	dbResult = repo.DatabaseConnection.Save(&tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *TourRepository) GetDurations(id string) ([]model.TourDuration, error) {
	var durationsJSON []byte
	if err := repo.DatabaseConnection.
		Raw("SELECT durations FROM tours WHERE id = ?", id).
		Row().
		Scan(&durationsJSON); err != nil {
		fmt.Println(fmt.Sprintf("Error: Couldn't get tours durations"))
		return nil, err
	}

	var durations []model.TourDuration

	if len(durationsJSON) > 0 {
		if err := json.Unmarshal(durationsJSON, &durations); err != nil {
			fmt.Println(fmt.Sprintf("Error: Couldn't unmarshal tours durations"))
			return nil, err
		}
	}

	return durations, nil
}
