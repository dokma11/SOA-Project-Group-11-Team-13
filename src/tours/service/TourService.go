package service

import (
	"fmt"
	"github.com/google/uuid"
	"time"
	"tours/dto"
	"tours/model"
	"tours/repo"
)

type TourService struct {
	TourRepository *repo.TourRepository
}

func (service *TourService) GetById(id uuid.UUID) (*model.Tour, error) {
	tour, err := service.TourRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &tour, nil
}

func (service *TourService) GetByAuthorId(authorId int) (*[]dto.TourResponseDto, error) {
	tours, err := service.TourRepository.GetByAuthorId(authorId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("tours with author id %s not found", authorId))
	}
	return &tours, nil
}

func (service *TourService) GetAll() (*[]model.Tour, error) {
	tours, err := service.TourRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no tours were found"))
	}
	return &tours, nil
}

func (service *TourService) GetPublished() (*[]model.Tour, error) {
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

func (service *TourService) Delete(id uuid.UUID) error {
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

func (service *TourService) Publish(id uuid.UUID) error {
	tour, err := service.TourRepository.GetById(id)

	if tour.Status != model.Published {
		tour.Status = model.Published
		tour.PublishDate = time.Now().Local() // moram proveriti da li ovako ili bez local
		err = service.TourRepository.Update(&tour)
		if err != nil {
			_ = fmt.Errorf(fmt.Sprintf("no tours were published"))
			return err
		}
		return nil
	}

	_ = fmt.Errorf(fmt.Sprintf("can not publish an already published tour"))
	return nil
}

func (service *TourService) Archive(id uuid.UUID) error {
	tour, err := service.TourRepository.GetById(id)

	if tour.Status == model.Published {
		tour.Status = model.Archived
		tour.ArchiveDate = time.Now().Local() // moram proveriti da li ovako ili bez local
		err = service.TourRepository.Update(&tour)
		if err != nil {
			_ = fmt.Errorf(fmt.Sprintf("no tours were archived"))
			return err
		}
		return nil
	}

	_ = fmt.Errorf(fmt.Sprintf("can not archive selected tour"))
	return nil
}
