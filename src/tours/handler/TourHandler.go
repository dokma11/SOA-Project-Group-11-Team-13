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

func (handler *TourHandler) GetById(ctx context.Context, request *tours.GetTourByIdRequest) (*tours.GetTourByIdResponse, error) {
	tour, _ := handler.TourService.GetById(request.ID)

	var tourResponse tours.Tour
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

	return &tours.GetTourByIdResponse{
		Tour: &tourResponse,
	}, nil
}

func (handler *TourHandler) GetByAuthorId(ctx context.Context, request *tours.GetToursByAuthorIdRequest) (*tours.GetToursByAuthorIdResponse, error) {
	toursList, _ := handler.TourService.GetByAuthorId(request.AuthorId)

	var toursResponse []*tours.Tour

	for i, tour := range *toursList {
		toursResponse[i].ID = tour.ID
		toursResponse[i].AuthorId = int32(tour.AuthorId)
		toursResponse[i].Name = tour.Name
		toursResponse[i].Description = tour.Description
		toursResponse[i].Difficulty = int32(tour.Difficulty)
		toursResponse[i].Tags = tour.Tags
		toursResponse[i].Status = tours.Tour_TourStatus(tour.Status)
		toursResponse[i].Price = tour.Price
		toursResponse[i].Distance = tour.Distance
		toursResponse[i].PublishDate = TimeToProtoTimestamp(tour.PublishDate)
		toursResponse[i].ArchiveDate = TimeToProtoTimestamp(tour.ArchiveDate)
		toursResponse[i].Category = tours.Tour_TourCategory(tour.Category)
		toursResponse[i].IsDeleted = tour.IsDeleted
		//toursResponse[i].KeyPoints = tour.KeyPoints	PROVERITI SAMO DA LI TREBA I STA TREBA
		//toursResponse[i].Equipment = tour.Equipment
		//toursResponse[i].Reviews = tour.Reviews
		//toursResponse[i].Durations = tour.Durations
	}

	return &tours.GetToursByAuthorIdResponse{
		Tours: toursResponse,
	}, nil
}

func (handler *TourHandler) GetAll(ctx context.Context, request *tours.GetAllToursRequest) (*tours.GetAllToursResponse, error) {
	toursList, _ := handler.TourService.GetAll()

	var toursResponse []*tours.Tour

	for i, tour := range *toursList {
		toursResponse[i].ID = tour.ID
		toursResponse[i].AuthorId = int32(tour.AuthorId)
		toursResponse[i].Name = tour.Name
		toursResponse[i].Description = tour.Description
		toursResponse[i].Difficulty = int32(tour.Difficulty)
		toursResponse[i].Tags = tour.Tags
		toursResponse[i].Status = tours.Tour_TourStatus(tour.Status)
		toursResponse[i].Price = tour.Price
		toursResponse[i].Distance = tour.Distance
		toursResponse[i].PublishDate = TimeToProtoTimestamp(tour.PublishDate)
		toursResponse[i].ArchiveDate = TimeToProtoTimestamp(tour.ArchiveDate)
		toursResponse[i].Category = tours.Tour_TourCategory(tour.Category)
		toursResponse[i].IsDeleted = tour.IsDeleted
		//toursResponse[i].KeyPoints = tour.KeyPoints	PROVERITI SAMO DA LI TREBA I STA TREBA
		//toursResponse[i].Equipment = tour.Equipment
		//toursResponse[i].Reviews = tour.Reviews
		//toursResponse[i].Durations = tour.Durations
	}

	return &tours.GetAllToursResponse{
		Tours: toursResponse,
	}, nil
}

func (handler *TourHandler) GetPublished(ctx context.Context, request *tours.GetPublishedToursRequest) (*tours.GetPublishedToursResponse, error) {
	toursList, _ := handler.TourService.GetPublished()

	var toursResponse []*tours.Tour

	for i, tour := range *toursList {
		toursResponse[i].ID = tour.ID
		toursResponse[i].AuthorId = int32(tour.AuthorId)
		toursResponse[i].Name = tour.Name
		toursResponse[i].Description = tour.Description
		toursResponse[i].Difficulty = int32(tour.Difficulty)
		toursResponse[i].Tags = tour.Tags
		toursResponse[i].Status = tours.Tour_TourStatus(tour.Status)
		toursResponse[i].Price = tour.Price
		toursResponse[i].Distance = tour.Distance
		toursResponse[i].PublishDate = TimeToProtoTimestamp(tour.PublishDate)
		toursResponse[i].ArchiveDate = TimeToProtoTimestamp(tour.ArchiveDate)
		toursResponse[i].Category = tours.Tour_TourCategory(tour.Category)
		toursResponse[i].IsDeleted = tour.IsDeleted
		//toursResponse[i].KeyPoints = tour.KeyPoints	PROVERITI SAMO DA LI TREBA I STA TREBA
		//toursResponse[i].Equipment = tour.Equipment
		//toursResponse[i].Reviews = tour.Reviews
		//toursResponse[i].Durations = tour.Durations
	}

	return &tours.GetPublishedToursResponse{
		Tours: toursResponse,
	}, nil
}

func (handler *TourHandler) Create(ctx context.Context, request *tours.CreateTourRequest) (*tours.CreateTourResponse, error) {
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

func (handler *TourHandler) Delete(ctx context.Context, request *tours.DeleteTourRequest) (*tours.DeleteTourResponse, error) {
	handler.TourService.Delete(request.ID)
	return &tours.DeleteTourResponse{}, nil
}

func (handler *TourHandler) Update(ctx context.Context, request *tours.UpdateTourRequest) (*tours.UpdateTourResponse, error) {
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

func (handler *TourHandler) Publish(ctx context.Context, request *tours.PublishTourRequest) (*tours.PublishTourResponse, error) {
	handler.TourService.Publish(request.ID)
	return &tours.PublishTourResponse{}, nil
}

func (handler *TourHandler) Archive(ctx context.Context, request *tours.ArchiveTourRequest) (*tours.ArchiveTourResponse, error) {
	handler.TourService.Archive(request.ID)
	return &tours.ArchiveTourResponse{}, nil
}

func (handler *TourHandler) GetEquipment(ctx context.Context, request *tours.GetToursEquipmentRequest) (*tours.GetToursEquipmentResponse, error) {
	equipmentList, _ := handler.TourService.GetEquipment(request.TourId)

	var equipmentResponse []*tours.Equipment

	for i, eq := range equipmentList {
		equipmentResponse[i].ID = eq.ID
		equipmentResponse[i].Name = eq.Name
		equipmentResponse[i].Description = eq.Description
		//equipmentResponse[i].Tours = e.Tours				Treba proveriti
	}

	return &tours.GetToursEquipmentResponse{
		Equipment: equipmentResponse,
	}, nil
}

func (handler *TourHandler) AddEquipment(ctx context.Context, request *tours.AddToursEquipmentRequest) (*tours.AddToursEquipmentResponse, error) {
	handler.TourService.AddEquipment(request.TourId, request.EquipmentId)
	return &tours.AddToursEquipmentResponse{}, nil
}

func (handler *TourHandler) DeleteEquipment(ctx context.Context, request *tours.DeleteToursEquipmentRequest) (*tours.DeleteToursEquipmentResponse, error) {
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
