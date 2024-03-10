package service

import (
	"fmt"
	"time"
	"tours/dto"
	"tours/model"
	"tours/repo"
)

type TourService struct {
	TourRepository *repo.TourRepository
}

func (service *TourService) GetById(id string) (*dto.TourResponseDto, error) {
	tour, err := service.TourRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &tour, nil
}

func (service *TourService) GetByAuthorId(authorId string) (*[]dto.TourResponseDto, error) {
	tours, err := service.TourRepository.GetByAuthorId(authorId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("tours with author id %s not found", authorId))
	}
	return &tours, nil
}

func (service *TourService) GetAll() (*[]dto.TourResponseDto, error) {
	tours, err := service.TourRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no tours were found"))
	}
	return &tours, nil
}

func (service *TourService) GetPublished() (*[]dto.TourResponseDto, error) {
	tours, err := service.TourRepository.GetPublished()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no published tours were found"))
	}
	return &tours, nil
}

func (service *TourService) Create(tour *model.Tour) error {
	err := service.TourRepository.Create(tour)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no tours were created"))
		return err
	}
	return nil
}

func (service *TourService) Delete(id string) error {
	err := service.TourRepository.Delete(id)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no tours were deleted"))
		return err
	}
	return nil
}

func (service *TourService) Update(tour *model.Tour) error {
	err := service.TourRepository.Update(tour)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no tours were updated"))
		return err
	}
	return nil
}

func (service *TourService) AddDurations(tour *model.Tour) error {
	err := service.TourRepository.AddDurations(tour)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no durations were addede to tours"))
		return err
	}
	return nil
}

func (service *TourService) Publish(id string) error {
	tourDto, _ := service.TourRepository.GetById(id)

	if tourDto.Status != dto.Published {
		tour, err := model.NewTour(tourDto.ID, tourDto.AuthorId, tourDto.Name, tourDto.Description, tourDto.Tags,
			tourDto.Difficulty, tourDto.ArchiveDate, time.Now().Local().Add(time.Hour), tourDto.Distance, model.Published,
			tourDto.Price, model.TourCategory(tourDto.Category), tourDto.IsDeleted, tourDto.KeyPoints, tourDto.Durations)
		if err != nil {
			return err
		}

		err = service.TourRepository.Update(tour)
		if err != nil {
			_ = fmt.Errorf(fmt.Sprintf("no tours were published"))
			return err
		}
		return nil
	}

	_ = fmt.Errorf(fmt.Sprintf("can not publish an already published tour"))
	return nil
}

func (service *TourService) Archive(id string) error {
	tourDto, _ := service.TourRepository.GetById(id)

	if tourDto.Status == dto.Published {
		tour, err := model.NewTour(tourDto.ID, tourDto.AuthorId, tourDto.Name, tourDto.Description, tourDto.Tags,
			tourDto.Difficulty, time.Now().Local().Add(time.Hour), tourDto.PublishDate, tourDto.Distance, model.Archived,
			tourDto.Price, model.TourCategory(tourDto.Category), tourDto.IsDeleted, tourDto.KeyPoints, tourDto.Durations)
		if err != nil {
			return err
		}

		err = service.TourRepository.Update(tour)
		if err != nil {
			_ = fmt.Errorf(fmt.Sprintf("no tours were archived"))
			return err
		}
		return nil
	}

	_ = fmt.Errorf(fmt.Sprintf("can not archive selected tour"))
	return nil
}

func (service *TourService) GetEquipment(tourId string) ([]model.Equipment, error) {
	equipmentList, err := service.TourRepository.GetEquipment(tourId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("equipment for tour with id %s not found", tourId))
	}
	return equipmentList, nil
}

func (service *TourService) AddEquipment(tourId string, equipmentId string) error {
	err := service.TourRepository.AddEquipment(tourId, equipmentId)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("failed to add equipment to tour with id %s", tourId))
	}
	return nil
}

func (service *TourService) DeleteEquipment(tourId string, equipmentId string) error {
	err := service.TourRepository.DeleteEquipment(tourId, equipmentId)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("failed to delete equipment from tour with id %s", tourId))
	}
	return nil
}
