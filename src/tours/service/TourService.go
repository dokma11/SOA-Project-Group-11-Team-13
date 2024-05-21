package service

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/sdk/trace"
	"log"
	"strconv"
	"time"
	"tours/dto"
	"tours/model"
	"tours/repo"
)

type TourService struct {
	TourRepository *repo.TourRepository
}

func (service *TourService) GetById(id string) (*dto.TourResponseDto, error) {
	log.Printf("Get tour by id service call, Tour ID: " + id + "\n")
	tour, err := service.TourRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	if tourDto, err := service.MapToDto(&tour, &dto.TourResponseDto{}); err == nil {
		return tourDto, nil
	} else {
		fmt.Println("Error mapping tour to DTO:", err)
		return nil, err
	}
}

func (service *TourService) GetByAuthorId(authorId string, tp *trace.TracerProvider, ctx context.Context) (*[]dto.TourResponseDto, error) {
	log.Printf("Get tour by author id service call, Author ID" + authorId + "\n")
	_, span := tp.Tracer("tours").Start(ctx, "tours-service-getByAuthorId")
	span.AddEvent("GetByAuthorId")
	defer func() { span.End() }()

	tours, err := service.TourRepository.GetByAuthorId(authorId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("tours with author id %s not found", authorId))
	}

	var tourDtos []dto.TourResponseDto
	for _, tour := range tours {
		if tourDto, err := service.MapToDto(&tour, &dto.TourResponseDto{}); err == nil {
			tourDtos = append(tourDtos, *tourDto)
		} else {
			fmt.Println("Error mapping tour to DTO:", err)
		}
	}

	return &tourDtos, nil
}

func (service *TourService) GetAll() (*[]dto.TourResponseDto, error) {
	log.Printf("Get all tours service call\n")
	tours, err := service.TourRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no tours were found"))
	}

	var tourDtos []dto.TourResponseDto
	for _, tour := range tours {
		if tourDto, err := service.MapToDto(&tour, &dto.TourResponseDto{}); err == nil {
			tourDtos = append(tourDtos, *tourDto)
		} else {
			fmt.Println("Error mapping tour to DTO:", err)
		}
	}

	return &tourDtos, nil
}

func (service *TourService) GetPublished() (*[]dto.TourResponseDto, error) {
	log.Printf("Get published tours service call\n")
	tours, err := service.TourRepository.GetPublished()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no published tours were found"))
	}

	var tourDtos []dto.TourResponseDto
	for _, tour := range tours {
		if tourDto, err := service.MapToDto(&tour, &dto.TourResponseDto{}); err == nil {
			tourDtos = append(tourDtos, *tourDto)
		} else {
			fmt.Println("Error mapping tour to DTO:", err)
		}
	}

	return &tourDtos, nil
}

func (service *TourService) Create(tour *model.Tour, tp *trace.TracerProvider, ctx context.Context) error {
	log.Printf("Create tour service call\n")
	_, span := tp.Tracer("tours").Start(ctx, "tours-service-create")
	defer func() { span.End() }()

	err := service.TourRepository.Create(tour, tp, ctx)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no tours were created"))
		return err
	}
	return nil
}

func (service *TourService) Delete(id string) error {
	log.Printf("Delete tour service call, Tour ID: " + id + "\n")
	err := service.TourRepository.Delete(id)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no tours were deleted"))
		return err
	}
	return nil
}

func (service *TourService) Update(tour *model.Tour) error {
	log.Printf("Update tour service call\n")
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
	log.Printf("Publish tour service call, Tour ID: " + id + "\n")
	tourDto, _ := service.TourRepository.GetById(id)

	if tourDto.Status != model.Published {
		tour, err := model.NewTour(tourDto.ID, tourDto.AuthorId, tourDto.Name, tourDto.Description, tourDto.Tags,
			tourDto.Difficulty, tourDto.ArchiveDate, time.Now().Local().Add(time.Hour), tourDto.Distance, model.Published,
			tourDto.Price, tourDto.Category, tourDto.IsDeleted, tourDto.KeyPoints, tourDto.Durations)
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
	log.Printf("Archive tour service call, Tour ID" + id + "\n")
	tourDto, _ := service.TourRepository.GetById(id)

	if tourDto.Status == model.Published {
		tour, err := model.NewTour(tourDto.ID, tourDto.AuthorId, tourDto.Name, tourDto.Description, tourDto.Tags,
			tourDto.Difficulty, time.Now().Local().Add(time.Hour), tourDto.PublishDate, tourDto.Distance, model.Archived,
			tourDto.Price, tourDto.Category, tourDto.IsDeleted, tourDto.KeyPoints, tourDto.Durations)
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

func (service *TourService) MapToDto(tour *model.Tour, tourDto *dto.TourResponseDto) (*dto.TourResponseDto, error) {
	var durations, err = service.TourRepository.GetDurations(strconv.FormatInt(tour.ID, 10))
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("failed to obtain durations while mapping to dto"))
	}

	tourDto = &dto.TourResponseDto{
		AverageRating: 0.0,
		Tags:          tour.Tags,
		KeyPoints:     tour.KeyPoints,
		Status:        dto.TourStatus(tour.Status),
		Name:          tour.Name,
		Description:   tour.Description,
		ID:            tour.ID,
		AuthorId:      tour.AuthorId,
		Durations:     durations,
		PublishDate:   tour.PublishDate,
		ArchiveDate:   tour.ArchiveDate,
		Category:      dto.TourCategory(tour.Category),
		IsDeleted:     tour.IsDeleted,
		Price:         tour.Price,
		Distance:      tour.Distance,
		Difficulty:    tour.Difficulty,
	}

	return tourDto, nil
}
