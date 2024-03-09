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
	tourDto, err := service.TourRepository.GetById(id)
	if tourDto.Status != dto.Published {
		tourDto.Status = dto.Published
		tourDto.PublishDate = time.Now().Local()

		var tour model.Tour
		tour.Tags = tourDto.Tags
		tour.KeyPoints = tourDto.KeyPoints
		tour.Status = model.TourStatus(tourDto.Status)
		tour.Name = tourDto.Name
		tour.Description = tourDto.Description
		tour.ID = tourDto.ID
		tour.Durations = tourDto.Durations
		tour.PublishDate = tourDto.PublishDate
		tour.ArchiveDate = tourDto.ArchiveDate
		tour.Category = model.TourCategory(tourDto.Category)
		tour.IsDeleted = tourDto.IsDeleted
		tour.Price = tourDto.Price
		tour.Distance = tourDto.Distance
		tour.Difficulty = tourDto.Difficulty
		tour.AuthorId = tourDto.AuthorId

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

func (service *TourService) Archive(id string) error {
	tourDto, err := service.TourRepository.GetById(id)

	if tourDto.Status == dto.Published {
		tourDto.Status = dto.Archived
		tourDto.ArchiveDate = time.Now().Local()

		var tour model.Tour
		tour.Tags = tourDto.Tags
		tour.KeyPoints = tourDto.KeyPoints
		tour.Status = model.TourStatus(tourDto.Status)
		tour.Name = tourDto.Name
		tour.Description = tourDto.Description
		tour.ID = tourDto.ID
		tour.Durations = tourDto.Durations
		tour.PublishDate = tourDto.PublishDate
		tour.ArchiveDate = tourDto.ArchiveDate
		tour.Category = model.TourCategory(tourDto.Category)
		tour.IsDeleted = tourDto.IsDeleted
		tour.Price = tourDto.Price
		tour.Distance = tourDto.Distance
		tour.Difficulty = tourDto.Difficulty
		tour.AuthorId = tourDto.AuthorId

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
