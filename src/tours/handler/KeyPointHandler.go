package handler

import (
	"context"
	"tours/model"
	keyPoints "tours/proto/keypoints"
	"tours/service"
)

type KeyPointHandler struct {
	KeyPointService *service.KeyPointService
	keyPoints.UnimplementedKeyPointsServiceServer
}

func (handler *KeyPointHandler) GetById(ctx context.Context, request *keyPoints.GetKeyPointByIdRequest) (*keyPoints.GetKeyPointByIdResponse, error) {
	keyPoint, _ := handler.KeyPointService.GetById(request.ID)

	var keyPointsResponse keyPoints.KeyPoint
	keyPointsResponse.ID = keyPoint.ID
	keyPointsResponse.TourId = keyPoint.TourId
	keyPointsResponse.Name = keyPoint.Name
	keyPointsResponse.Description = keyPoint.Description
	keyPointsResponse.Longitude = keyPoint.Longitude
	keyPointsResponse.Latitude = keyPoint.Latitude
	keyPointsResponse.LocationAddress = keyPoint.LocationAddress
	keyPointsResponse.ImagePath = keyPoint.ImagePath
	keyPointsResponse.Order = keyPoint.Order

	return &keyPoints.GetKeyPointByIdResponse{
		KeyPoint: &keyPointsResponse,
	}, nil
}

func (handler *KeyPointHandler) GetAll(ctx context.Context, request *keyPoints.GetAllKeyPointsRequest) (*keyPoints.GetAllKeyPointsResponse, error) {
	keyPointList, _ := handler.KeyPointService.GetAll()

	var keyPointsResponse []*keyPoints.KeyPoint

	for i, keyPoint := range *keyPointList {
		keyPointsResponse[i].ID = keyPoint.ID
		keyPointsResponse[i].TourId = keyPoint.TourId
		keyPointsResponse[i].Name = keyPoint.Name
		keyPointsResponse[i].Description = keyPoint.Description
		keyPointsResponse[i].Longitude = keyPoint.Longitude
		keyPointsResponse[i].Latitude = keyPoint.Latitude
		keyPointsResponse[i].LocationAddress = keyPoint.LocationAddress
		keyPointsResponse[i].ImagePath = keyPoint.ImagePath
		keyPointsResponse[i].Order = keyPoint.Order
	}

	return &keyPoints.GetAllKeyPointsResponse{
		KeyPoints: keyPointsResponse,
	}, nil
}

func (handler *KeyPointHandler) GetAllByTourId(ctx context.Context, request *keyPoints.GetKeyPointsByTourIdRequest) (*keyPoints.GetKeyPointsByTourIdResponse, error) {
	keyPointList, _ := handler.KeyPointService.GetAllByTourId(request.TourId)

	var keyPointsResponse []*keyPoints.KeyPoint

	for i, keyPoint := range *keyPointList {
		keyPointsResponse[i].ID = keyPoint.ID
		keyPointsResponse[i].TourId = keyPoint.TourId
		keyPointsResponse[i].Name = keyPoint.Name
		keyPointsResponse[i].Description = keyPoint.Description
		keyPointsResponse[i].Longitude = keyPoint.Longitude
		keyPointsResponse[i].Latitude = keyPoint.Latitude
		keyPointsResponse[i].LocationAddress = keyPoint.LocationAddress
		keyPointsResponse[i].ImagePath = keyPoint.ImagePath
		keyPointsResponse[i].Order = keyPoint.Order
	}

	return &keyPoints.GetKeyPointsByTourIdResponse{
		KeyPoints: keyPointsResponse,
	}, nil
}

func (handler *KeyPointHandler) Create(ctx context.Context, request *keyPoints.CreateKeyPointRequest) (*keyPoints.CreateKeyPointResponse, error) {
	keyPoint := model.KeyPoint{}

	keyPoint.ID = request.KeyPoint.ID
	keyPoint.TourId = request.KeyPoint.TourId
	keyPoint.Name = request.KeyPoint.Name
	keyPoint.Description = request.KeyPoint.Description
	keyPoint.Longitude = request.KeyPoint.Longitude
	keyPoint.Latitude = request.KeyPoint.Latitude
	keyPoint.LocationAddress = request.KeyPoint.LocationAddress
	keyPoint.ImagePath = request.KeyPoint.ImagePath
	keyPoint.Order = request.KeyPoint.Order

	handler.KeyPointService.Create(&keyPoint)

	return &keyPoints.CreateKeyPointResponse{}, nil
}

func (handler *KeyPointHandler) Delete(ctx context.Context, request *keyPoints.DeleteKeyPointRequest) (*keyPoints.DeleteKeyPointResponse, error) {
	handler.KeyPointService.Delete(request.ID)
	return &keyPoints.DeleteKeyPointResponse{}, nil
}

func (handler *KeyPointHandler) Update(ctx context.Context, request *keyPoints.UpdateKeyPointRequest) (*keyPoints.UpdateKeyPointResponse, error) {
	keyPoint := model.KeyPoint{}

	keyPoint.ID = request.KeyPoint.ID
	keyPoint.TourId = request.KeyPoint.TourId
	keyPoint.Name = request.KeyPoint.Name
	keyPoint.Description = request.KeyPoint.Description
	keyPoint.Longitude = request.KeyPoint.Longitude
	keyPoint.Latitude = request.KeyPoint.Latitude
	keyPoint.LocationAddress = request.KeyPoint.LocationAddress
	keyPoint.ImagePath = request.KeyPoint.ImagePath
	keyPoint.Order = request.KeyPoint.Order

	handler.KeyPointService.Update(&keyPoint)

	return &keyPoints.UpdateKeyPointResponse{}, nil
}
