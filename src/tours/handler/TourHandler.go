package handler

import (
	"context"
	"log"
	"strconv"
	"time"
	"tours/dto"
	"tours/model"
	"tours/proto/tours"
	"tours/service"

	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
)

type TourHandler struct {
	TourService *service.TourService
	tp          *trace.TracerProvider
	CommandPublisher saga.Publisher
	tours.UnimplementedToursServiceServer

}

func NewTourHandler(tourService *service.TourService, tp *trace.TracerProvider, cp saga.Publisher) *TourHandler {
	return &TourHandler{
		TourService: tourService,
		tp:          tp,
		CommandPublisher: cp,
	}
}

func (handler *TourHandler) GetTourById(ctx context.Context, request *tours.GetTourByIdRequest) (*tours.GetTourByIdResponse, error) {
	log.Printf("Get tour by id call\n")
	_, span := handler.tp.Tracer("tours").Start(ctx, "tours-get-by-tour-id")
	defer func() { span.End() }()

	tour, _ := handler.TourService.GetById(request.ID)

	tourResponse := tours.Tour{}
	tourResponse.Id = tour.ID
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

	ret := &tours.GetTourByIdResponse{
		Tour: &tourResponse,
	}

	return ret, nil
}

func (handler *TourHandler) GetToursByAuthorId(ctx context.Context, request *tours.GetToursByAuthorIdRequest) (*tours.GetToursByAuthorIdResponse, error) {
	log.Printf("Get tours by author id handler call\n")
	_, span := handler.tp.Tracer("tours").Start(ctx, "tours-get-by-author-id")
	defer func() { span.End() }()

	toursList, _ := handler.TourService.GetByAuthorId(request.AuthorId, handler.tp, ctx)

	toursResponse := make([]*tours.Tour, len(*toursList))

	if toursList != nil && len(*toursList) > 0 {
		for i, tour := range *toursList {

			keypointList := make([]*tours.TourKeyPoint, len(tour.KeyPoints))
			if tour.KeyPoints != nil && len(tour.KeyPoints) > 0 {
				for index, kp := range tour.KeyPoints {
					keypointList[index] = &tours.TourKeyPoint{
						Id:              kp.ID,
						TourId:          kp.TourId,
						Name:            kp.Name,
						Description:     kp.Description,
						Longitude:       kp.Longitude,
						Latitude:        kp.Latitude,
						LocationAddress: kp.LocationAddress,
						ImagePath:       kp.ImagePath,
						Order:           kp.Order,
					}
				}
			}

			durationsList := make([]*tours.TourDuration, len(tour.Durations))
			if tour.Durations != nil && len(tour.Durations) > 0 {
				for index, duration := range tour.Durations {
					durationsList[index] = &tours.TourDuration{
						Duration:      int32(duration.Duration),
						TransportType: tours.TourDuration_TransportType(duration.TransportType),
					}
				}
			}

			toursResponse[i] = &tours.Tour{
				Id:          tour.ID,
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
				KeyPoints:   keypointList,
				Durations:   durationsList,
			}
		}
	}

	ret := &tours.GetToursByAuthorIdResponse{
		Tours: toursResponse,
	}

	return ret, nil
}

func (handler *TourHandler) GetAllTours(ctx context.Context, request *tours.GetAllToursRequest) (*tours.GetAllToursResponse, error) {
	log.Printf("Get all tours handler call\n")
	_, span := handler.tp.Tracer("tours").Start(ctx, "tours-get-all")
	defer func() { span.End() }()

	toursList, _ := handler.TourService.GetAll()

	toursResponse := make([]*tours.Tour, len(*toursList))

	if toursList != nil && len(*toursList) > 0 {
		for i, tour := range *toursList {

			keypointList := make([]*tours.TourKeyPoint, len(tour.KeyPoints))
			if tour.KeyPoints != nil && len(tour.KeyPoints) > 0 {
				for index, kp := range tour.KeyPoints {
					keypointList[index] = &tours.TourKeyPoint{
						Id:              kp.ID,
						TourId:          kp.TourId,
						Name:            kp.Name,
						Description:     kp.Description,
						Longitude:       kp.Longitude,
						Latitude:        kp.Latitude,
						LocationAddress: kp.LocationAddress,
						ImagePath:       kp.ImagePath,
						Order:           kp.Order,
					}
				}
			}

			durationsList := make([]*tours.TourDuration, len(tour.Durations))
			if tour.Durations != nil && len(tour.Durations) > 0 {
				for index, duration := range tour.Durations {
					durationsList[index] = &tours.TourDuration{
						Duration:      int32(duration.Duration),
						TransportType: tours.TourDuration_TransportType(duration.TransportType),
					}
				}
			}

			toursResponse[i] = &tours.Tour{
				Id:          tour.ID,
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
				KeyPoints:   keypointList,
				Durations:   durationsList,
			}
		}
	}

	ret := &tours.GetAllToursResponse{
		Tours: toursResponse,
	}

	return ret, nil
}

func (handler *TourHandler) GetPublishedTours(ctx context.Context, request *tours.GetPublishedToursRequest) (*tours.GetPublishedToursResponse, error) {
	log.Printf("Get published tours handler call\n")
	_, span := handler.tp.Tracer("tours").Start(ctx, "tours-get-published")
	defer func() { span.End() }()

	toursList, _ := handler.TourService.GetPublished()

	toursResponse := make([]*tours.Tour, len(*toursList))

	if toursList != nil && len(*toursList) > 0 {
		for i, tour := range *toursList {

			keypointList := make([]*tours.TourKeyPoint, len(tour.KeyPoints))
			if tour.KeyPoints != nil && len(tour.KeyPoints) > 0 {
				for index, kp := range tour.KeyPoints {
					keypointList[index] = &tours.TourKeyPoint{
						Id:              kp.ID,
						TourId:          kp.TourId,
						Name:            kp.Name,
						Description:     kp.Description,
						Longitude:       kp.Longitude,
						Latitude:        kp.Latitude,
						LocationAddress: kp.LocationAddress,
						ImagePath:       kp.ImagePath,
						Order:           kp.Order,
					}
				}
			}

			durationsList := make([]*tours.TourDuration, len(tour.Durations))
			if tour.Durations != nil && len(tour.Durations) > 0 {
				for index, duration := range tour.Durations {
					durationsList[index] = &tours.TourDuration{
						Duration:      int32(duration.Duration),
						TransportType: tours.TourDuration_TransportType(duration.TransportType),
					}
				}
			}

			toursResponse[i] = &tours.Tour{
				Id:          tour.ID,
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
				KeyPoints:   keypointList,
				Durations:   durationsList,
			}
		}
	}

	ret := &tours.GetPublishedToursResponse{
		Tours: toursResponse,
	}

	return ret, nil
}

func (handler *TourHandler) CreateTour(ctx context.Context, request *tours.CreateTourRequest) (*tours.CreateTourResponse, error) {
	log.Printf("Create tour handler call")
	tour := model.Tour{}

	tour.ID = request.Tour.Id
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

	keypointList := make([]model.KeyPoint, len(tour.KeyPoints))
	if tour.KeyPoints != nil && len(tour.KeyPoints) > 0 {
		for index, kp := range tour.KeyPoints {
			keypointList[index] = model.KeyPoint{
				ID:              kp.ID,
				TourId:          kp.TourId,
				Name:            kp.Name,
				Description:     kp.Description,
				Longitude:       kp.Longitude,
				Latitude:        kp.Latitude,
				LocationAddress: kp.LocationAddress,
				ImagePath:       kp.ImagePath,
				Order:           kp.Order,
			}
		}
	}

	durationsList := make([]model.TourDuration, len(tour.Durations))
	if tour.Durations != nil && len(tour.Durations) > 0 {
		for index, duration := range tour.Durations {
			durationsList[index] = model.TourDuration{
				Duration:      duration.Duration,
				TransportType: duration.TransportType,
			}
		}
	}

	tour.KeyPoints = keypointList
	tour.Durations = durationsList

	_, span := handler.tp.Tracer("tours").Start(ctx, "tours-handler-create")
	defer func() { span.End() }()

	span.AddEvent("Calling create service method")

	handler.TourService.Create(&tour, handler.tp, ctx)

	span.AddEvent("Tour successfully created")

	tourSearchSagaRequestDTO := dto.TourSearchSagaRequestDTO {
		ID: tour.ID,
		Name: tour.Name,
		Description: tour.Description,
	};
	
	handler.CommandPublisher.Publish(tourSearchSagaRequestDTO)

	return &tours.CreateTourResponse{}, nil
}

func (handler *TourHandler) DeleteTour(ctx context.Context, request *tours.DeleteTourRequest) (*tours.DeleteTourResponse, error) {
	log.Printf("Delete tour handler call\n")
	_, span := handler.tp.Tracer("tours").Start(ctx, "tours-delete")
	defer func() { span.End() }()

	handler.TourService.Delete(request.ID)
	return &tours.DeleteTourResponse{}, nil
}

func (handler *TourHandler) UpdateTour(ctx context.Context, request *tours.UpdateTourRequest) (*tours.UpdateTourResponse, error) {
	log.Printf("Update tour handler call\n")
	tour := model.Tour{}

	tour.ID = request.Tour.Id
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

	keypointList := make([]model.KeyPoint, len(request.Tour.KeyPoints))
	if request.Tour.KeyPoints != nil && len(request.Tour.KeyPoints) > 0 {
		for index, kp := range request.Tour.KeyPoints {
			keypointList[index] = model.KeyPoint{
				ID:              kp.Id,
				TourId:          kp.TourId,
				Name:            kp.Name,
				Description:     kp.Description,
				Longitude:       kp.Longitude,
				Latitude:        kp.Latitude,
				LocationAddress: kp.LocationAddress,
				ImagePath:       kp.ImagePath,
				Order:           kp.Order,
			}
		}
	}

	durationsList := make([]model.TourDuration, len(request.Tour.Durations))
	if request.Tour.Durations != nil && len(request.Tour.Durations) > 0 {
		for index, duration := range request.Tour.Durations {
			durationsList[index] = model.TourDuration{
				Duration:      int(duration.Duration),
				TransportType: model.TransportType(duration.TransportType),
			}
		}
	}

	tour.KeyPoints = keypointList
	tour.Durations = durationsList

	_, span := handler.tp.Tracer("tours").Start(ctx, "tours-update")
	defer func() { span.End() }()

	handler.TourService.Update(&tour)

	return &tours.UpdateTourResponse{}, nil
}

func (handler *TourHandler) AddToursDurations(ctx context.Context, request *tours.AddToursDurationsRequest) (*tours.AddToursDurationsResponse, error) {
	tour := model.Tour{}

	tour.ID = request.Tour.Id
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

	keypointList := make([]model.KeyPoint, len(request.Tour.KeyPoints))
	if request.Tour.KeyPoints != nil && len(request.Tour.KeyPoints) > 0 {
		for index, kp := range request.Tour.KeyPoints {
			keypointList[index] = model.KeyPoint{
				ID:              kp.Id,
				TourId:          kp.TourId,
				Name:            kp.Name,
				Description:     kp.Description,
				Longitude:       kp.Longitude,
				Latitude:        kp.Latitude,
				LocationAddress: kp.LocationAddress,
				ImagePath:       kp.ImagePath,
				Order:           kp.Order,
			}
		}
	}

	durationsList := make([]model.TourDuration, len(request.Tour.Durations))
	if request.Tour.Durations != nil && len(request.Tour.Durations) > 0 {
		for index, duration := range request.Tour.Durations {
			durationsList[index] = model.TourDuration{
				Duration:      int(duration.Duration),
				TransportType: model.TransportType(duration.TransportType),
			}
		}
	}

	tour.KeyPoints = keypointList
	tour.Durations = durationsList

	handler.TourService.AddDurations(&tour)

	return &tours.AddToursDurationsResponse{}, nil
}

func (handler *TourHandler) PublishTour(ctx context.Context, request *tours.PublishTourRequest) (*tours.PublishTourResponse, error) {
	log.Printf("Publish tour handler call\n")
	_, span := handler.tp.Tracer("tours").Start(ctx, "tours-publish")
	defer func() { span.End() }()

	handler.TourService.Publish(strconv.FormatInt(request.Tour.Id, 10))
	return &tours.PublishTourResponse{}, nil
}

func (handler *TourHandler) ArchiveTour(ctx context.Context, request *tours.ArchiveTourRequest) (*tours.ArchiveTourResponse, error) {
	log.Printf("Archive tour handler call\n")
	_, span := handler.tp.Tracer("tours").Start(ctx, "tours-archive")
	defer func() { span.End() }()

	handler.TourService.Archive(strconv.FormatInt(request.Tour.Id, 10))
	return &tours.ArchiveTourResponse{}, nil
}

func (handler *TourHandler) GetToursEquipment(ctx context.Context, request *tours.GetToursEquipmentRequest) (*tours.GetToursEquipmentResponse, error) {
	equipmentList, _ := handler.TourService.GetEquipment(request.TourId)

	equipmentResponse := make([]*tours.TourEquipment, len(equipmentList))

	if equipmentList != nil && len(equipmentList) > 0 {
		for i, eq := range equipmentList {
			equipmentResponse[i] = &tours.TourEquipment{
				Id:          eq.ID,
				Name:        eq.Name,
				Description: eq.Description,
			}
		}
	}

	ret := &tours.GetToursEquipmentResponse{
		Equipment: equipmentResponse,
	}

	return ret, nil
}

func (handler *TourHandler) AddToursEquipment(ctx context.Context, request *tours.AddToursEquipmentRequest) (*tours.AddToursEquipmentResponse, error) {
	handler.TourService.AddEquipment(request.Ids.TourId, request.Ids.EquipmentId)
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
