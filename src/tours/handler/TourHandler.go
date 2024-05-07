package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
	"tours/model"
	"tours/proto/tours"
	"tours/service"
)

type TourHandler struct {
	TourService *service.TourService
	tours.UnimplementedToursServiceServer
}

func (handler *TourHandler) GetTourById(ctx context.Context, request *tours.GetTourByIdRequest) (*tours.GetTourByIdResponse, error) {
	tour, _ := handler.TourService.GetById(request.ID)

	tourResponse := tours.Tour{}
	tourResponse.ID = tour.ID
	tourResponse.AuthorId = int32(tour.AuthorId)
	tourResponse.Name = tour.Name
	tourResponse.Description = tour.Description
	tourResponse.Difficulty = int32(tour.Difficulty)
	tourResponse.Tags = tour.Tags
	tourResponse.Status = tours.Tour_TourStatus(tour.Status)
	tourResponse.Price = tour.Price
	tourResponse.Distance = tour.Distance
	tourResponse.PublishDate = TimeToProtoTimestamp(tour.PublishDate)
	tourResponse.ArchiveDate = TimeToProtoTimestamp(tour.ArchiveDate)
	tourResponse.Category = tours.Tour_TourCategory(tour.Category)
	tourResponse.IsDeleted = tour.IsDeleted
	//tourResponse.KeyPoints = tour.KeyPoints	PROVERITI SAMO DA LI TREBA I STA TREBA
	//tourResponse.Equipment = tour.Equipment
	//tourResponse.Reviews = tour.Reviews
	//tourResponse.Durations = tour.Durations

	ret := &tours.GetTourByIdResponse{
		Tour: &tourResponse,
	}

	return ret, nil
}

func (handler *TourHandler) GetToursByAuthorId(ctx context.Context, request *tours.GetToursByAuthorIdRequest) (*tours.GetToursByAuthorIdResponse, error) {
	toursList, _ := handler.TourService.GetByAuthorId(request.AuthorId)

	toursResponse := make([]*tours.Tour, len(*toursList))

	if toursList != nil && len(*toursList) > 0 {
		for i, tour := range *toursList {
			toursResponse[i] = &tours.Tour{
				ID:          tour.ID,
				AuthorId:    int32(tour.AuthorId),
				Name:        tour.Name,
				Description: tour.Description,
				Difficulty:  int32(tour.Difficulty),
				Tags:        tour.Tags,
				Status:      tours.Tour_TourStatus(tour.Status),
				Price:       tour.Price,
				Distance:    tour.Distance,
				PublishDate: TimeToProtoTimestamp(tour.PublishDate),
				ArchiveDate: TimeToProtoTimestamp(tour.ArchiveDate),
				Category:    tours.Tour_TourCategory(tour.Category),
				IsDeleted:   tour.IsDeleted,
				//KeyPoints : tour.KeyPoints,	PROVERITI SAMO DA LI TREBA I STA TREBA
				//Equipment : tour.Equipment,
				//Reviews : tour.Reviews,
				//Durations : tour.Durations,
			}
		}
	}

	ret := &tours.GetToursByAuthorIdResponse{
		Tours: toursResponse,
	}

	return ret, nil
}

func (handler *TourHandler) GetAllTours(ctx context.Context, request *tours.GetAllToursRequest) (*tours.GetAllToursResponse, error) {
	toursList, _ := handler.TourService.GetAll()

	toursResponse := make([]*tours.Tour, len(*toursList))

	if toursList != nil && len(*toursList) > 0 {
		for i, tour := range *toursList {
			toursResponse[i] = &tours.Tour{
				ID:          tour.ID,
				AuthorId:    int32(tour.AuthorId),
				Name:        tour.Name,
				Description: tour.Description,
				Difficulty:  int32(tour.Difficulty),
				Tags:        tour.Tags,
				Status:      tours.Tour_TourStatus(tour.Status),
				Price:       tour.Price,
				Distance:    tour.Distance,
				PublishDate: TimeToProtoTimestamp(tour.PublishDate),
				ArchiveDate: TimeToProtoTimestamp(tour.ArchiveDate),
				Category:    tours.Tour_TourCategory(tour.Category),
				IsDeleted:   tour.IsDeleted,
				//KeyPoints : tour.KeyPoints,	PROVERITI SAMO DA LI TREBA I STA TREBA
				//Equipment : tour.Equipment,
				//Reviews : tour.Reviews,
				//Durations : tour.Durations,
			}
		}
	}

	ret := &tours.GetAllToursResponse{
		Tours: toursResponse,
	}

	return ret, nil
}

func (handler *TourHandler) GetPublishedTours(ctx context.Context, request *tours.GetPublishedToursRequest) (*tours.GetPublishedToursResponse, error) {
	toursList, _ := handler.TourService.GetPublished()

	toursResponse := make([]*tours.Tour, len(*toursList))

	if toursList != nil && len(*toursList) > 0 {
		for i, tour := range *toursList {
			toursResponse[i] = &tours.Tour{
				ID:          tour.ID,
				AuthorId:    int32(tour.AuthorId),
				Name:        tour.Name,
				Description: tour.Description,
				Difficulty:  int32(tour.Difficulty),
				Tags:        tour.Tags,
				Status:      tours.Tour_TourStatus(tour.Status),
				Price:       tour.Price,
				Distance:    tour.Distance,
				PublishDate: TimeToProtoTimestamp(tour.PublishDate),
				ArchiveDate: TimeToProtoTimestamp(tour.ArchiveDate),
				Category:    tours.Tour_TourCategory(tour.Category),
				IsDeleted:   tour.IsDeleted,
				//KeyPoints : tour.KeyPoints,	PROVERITI SAMO DA LI TREBA I STA TREBA
				//Equipment : tour.Equipment,
				//Reviews : tour.Reviews,
				//Durations : tour.Durations,
			}
		}
	}

	ret := &tours.GetPublishedToursResponse{
		Tours: toursResponse,
	}

	return ret, nil
}

func (handler *TourHandler) CreateTour(ctx context.Context, request *tours.CreateTourRequest) (*tours.CreateTourResponse, error) {
	tour := model.Tour{}

	tour.ID = request.Tour.ID
	tour.AuthorId = int(request.Tour.AuthorId)
	tour.Name = request.Tour.Name
	tour.Description = request.Tour.Description
	tour.Difficulty = int(request.Tour.Difficulty)
	tour.Tags = request.Tour.Tags
	tour.Status = model.TourStatus(request.Tour.Status)
	tour.Price = request.Tour.Price
	tour.Distance = request.Tour.Distance
	tour.PublishDate, _ = ProtoTimestampToTime(request.Tour.PublishDate)
	tour.ArchiveDate, _ = ProtoTimestampToTime(request.Tour.ArchiveDate)
	tour.Category = model.TourCategory(request.Tour.Category)
	tour.IsDeleted = request.Tour.IsDeleted
	//tourResponse.KeyPoints = request.Tour.KeyPoints	PROVERITI SAMO DA LI TREBA I STA TREBA
	//tourResponse.Equipment = request.Tour.Equipment
	//tourResponse.Reviews = request.Tour.Reviews
	//tourResponse.Durations = request.Tour.Durations

	handler.TourService.Create(&tour)

	return &tours.CreateTourResponse{}, nil
}

func (handler *TourHandler) DeleteTour(ctx context.Context, request *tours.DeleteTourRequest) (*tours.DeleteTourResponse, error) {
	handler.TourService.Delete(request.ID)
	return &tours.DeleteTourResponse{}, nil
}

func (handler *TourHandler) UpdateTour(ctx context.Context, request *tours.UpdateTourRequest) (*tours.UpdateTourResponse, error) {
	tour := model.Tour{}

	tour.ID = request.Tour.ID
	tour.AuthorId = int(request.Tour.AuthorId)
	tour.Name = request.Tour.Name
	tour.Description = request.Tour.Description
	tour.Difficulty = int(request.Tour.Difficulty)
	tour.Tags = request.Tour.Tags
	tour.Status = model.TourStatus(request.Tour.Status)
	tour.Price = request.Tour.Price
	tour.Distance = request.Tour.Distance
	tour.PublishDate, _ = ProtoTimestampToTime(request.Tour.PublishDate)
	tour.ArchiveDate, _ = ProtoTimestampToTime(request.Tour.ArchiveDate)
	tour.Category = model.TourCategory(request.Tour.Category)
	tour.IsDeleted = request.Tour.IsDeleted
	//tourResponse.KeyPoints = request.Tour.KeyPoints	PROVERITI SAMO DA LI TREBA I STA TREBA
	//tourResponse.Equipment = request.Tour.Equipment
	//tourResponse.Reviews = request.Tour.Reviews
	//tourResponse.Durations = request.Tour.Durations

	handler.TourService.Update(&tour)

	return &tours.UpdateTourResponse{}, nil
}

func (handler *TourHandler) AddToursDurations(ctx context.Context, request *tours.AddToursDurationsRequest) (*tours.AddToursDurationsResponse, error) {
	tour := model.Tour{}

	tour.ID = request.Tour.ID
	tour.AuthorId = int(request.Tour.AuthorId)
	tour.Name = request.Tour.Name
	tour.Description = request.Tour.Description
	tour.Difficulty = int(request.Tour.Difficulty)
	tour.Tags = request.Tour.Tags
	tour.Status = model.TourStatus(request.Tour.Status)
	tour.Price = request.Tour.Price
	tour.Distance = request.Tour.Distance
	tour.PublishDate, _ = ProtoTimestampToTime(request.Tour.PublishDate)
	tour.ArchiveDate, _ = ProtoTimestampToTime(request.Tour.ArchiveDate)
	tour.Category = model.TourCategory(request.Tour.Category)
	tour.IsDeleted = request.Tour.IsDeleted
	//tourResponse.KeyPoints = request.Tour.KeyPoints	PROVERITI SAMO DA LI TREBA I STA TREBA
	//tourResponse.Equipment = request.Tour.Equipment
	//tourResponse.Reviews = request.Tour.Reviews
	//tourResponse.Durations = request.Tour.Durations

	handler.TourService.AddDurations(&tour)

	return &tours.AddToursDurationsResponse{}, nil
}

func (handler *TourHandler) PublishTour(ctx context.Context, request *tours.PublishTourRequest) (*tours.PublishTourResponse, error) {
	handler.TourService.Publish(request.ID)
	return &tours.PublishTourResponse{}, nil
}

func (handler *TourHandler) ArchiveTour(ctx context.Context, request *tours.ArchiveTourRequest) (*tours.ArchiveTourResponse, error) {
	handler.TourService.Archive(request.ID)
	return &tours.ArchiveTourResponse{}, nil
}

func (handler *TourHandler) GetToursEquipment(ctx context.Context, request *tours.GetToursEquipmentRequest) (*tours.GetToursEquipmentResponse, error) {
	equipmentList, _ := handler.TourService.GetEquipment(request.TourId)

	equipmentResponse := make([]*tours.TourEquipment, len(equipmentList))

	if equipmentList != nil && len(equipmentList) > 0 {
		for i, eq := range equipmentList {
			equipmentResponse[i] = &tours.TourEquipment{
				ID:          eq.ID,
				Name:        eq.Name,
				Description: eq.Description,
				//Tours: eq.Tours,	Treba proveriti
			}
		}
	}

	ret := &tours.GetToursEquipmentResponse{
		Equipment: equipmentResponse,
	}

	return ret, nil
}

func (handler *TourHandler) AddToursEquipment(ctx context.Context, request *tours.AddToursEquipmentRequest) (*tours.AddToursEquipmentResponse, error) {
	handler.TourService.AddEquipment(request.TourId, request.EquipmentId)
	return &tours.AddToursEquipmentResponse{}, nil
}

func (handler *TourHandler) DeleteToursEquipment(ctx context.Context, request *tours.DeleteToursEquipmentRequest) (*tours.DeleteToursEquipmentResponse, error) {
	handler.TourService.DeleteEquipment(request.TourId, request.EquipmentId)
	return &tours.DeleteToursEquipmentResponse{}, nil
}

func TimeToProtoTimestamp(t time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(t)
	return ts
}

func ProtoTimestampToTime(ts *timestamp.Timestamp) (time.Time, error) {
	return ptypes.Timestamp(ts)
}
