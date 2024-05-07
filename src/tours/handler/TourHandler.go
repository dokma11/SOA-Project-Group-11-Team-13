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

func (handler *TourHandler) GetById(ctx context.Context, request *tours.GetByIdRequest) (*tours.GetByIdResponse, error) {
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

	ret := &tours.GetByIdResponse{
		Tour: &tourResponse,
	}

	return ret, nil
}

func (handler *TourHandler) GetByAuthorId(ctx context.Context, request *tours.GetByAuthorIdRequest) (*tours.GetByAuthorIdResponse, error) {
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

	ret := &tours.GetByAuthorIdResponse{
		Tours: toursResponse,
	}

	return ret, nil
}

func (handler *TourHandler) GetAll(ctx context.Context, request *tours.GetAllRequest) (*tours.GetAllResponse, error) {
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

	ret := &tours.GetAllResponse{
		Tours: toursResponse,
	}

	return ret, nil
}

func (handler *TourHandler) GetPublished(ctx context.Context, request *tours.GetPublishedRequest) (*tours.GetPublishedResponse, error) {
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

	ret := &tours.GetPublishedResponse{
		Tours: toursResponse,
	}

	return ret, nil
}

func (handler *TourHandler) Create(ctx context.Context, request *tours.CreateRequest) (*tours.CreateResponse, error) {
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

	return &tours.CreateResponse{}, nil
}

func (handler *TourHandler) Delete(ctx context.Context, request *tours.DeleteRequest) (*tours.DeleteResponse, error) {
	handler.TourService.Delete(request.ID)
	return &tours.DeleteResponse{}, nil
}

func (handler *TourHandler) Update(ctx context.Context, request *tours.UpdateRequest) (*tours.UpdateResponse, error) {
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

	return &tours.UpdateResponse{}, nil
}

func (handler *TourHandler) AddDurations(ctx context.Context, request *tours.AddDurationsRequest) (*tours.AddDurationsResponse, error) {
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

	return &tours.AddDurationsResponse{}, nil
}

func (handler *TourHandler) Publish(ctx context.Context, request *tours.PublishRequest) (*tours.PublishResponse, error) {
	handler.TourService.Publish(request.ID)
	return &tours.PublishResponse{}, nil
}

func (handler *TourHandler) Archive(ctx context.Context, request *tours.ArchiveRequest) (*tours.ArchiveResponse, error) {
	handler.TourService.Archive(request.ID)
	return &tours.ArchiveResponse{}, nil
}

func (handler *TourHandler) GetEquipment(ctx context.Context, request *tours.GetEquipmentRequest) (*tours.GetEquipmentResponse, error) {
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

	ret := &tours.GetEquipmentResponse{
		Equipment: equipmentResponse,
	}

	return ret, nil
}

func (handler *TourHandler) AddEquipment(ctx context.Context, request *tours.AddEquipmentRequest) (*tours.AddEquipmentResponse, error) {
	handler.TourService.AddEquipment(request.TourId, request.EquipmentId)
	return &tours.AddEquipmentResponse{}, nil
}

func (handler *TourHandler) DeleteEquipment(ctx context.Context, request *tours.DeleteEquipmentRequest) (*tours.DeleteEquipmentResponse, error) {
	handler.TourService.DeleteEquipment(request.TourId, request.EquipmentId)
	return &tours.DeleteEquipmentResponse{}, nil
}

func TimeToProtoTimestamp(t time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(t)
	return ts
}

func ProtoTimestampToTime(ts *timestamp.Timestamp) (time.Time, error) {
	return ptypes.Timestamp(ts)
}
