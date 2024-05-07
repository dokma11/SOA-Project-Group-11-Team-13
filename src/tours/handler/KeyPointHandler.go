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

func (handler *KeyPointHandler) GetById(ctx context.Context, request *keyPoints.KeyPointGetByIdRequest) (*keyPoints.KeyPointGetByIdResponse, error) {
	keyPoint, _ := handler.KeyPointService.GetById(request.ID)

	keyPointsResponse := keyPoints.KeyPoint{}
	keyPointsResponse.ID = keyPoint.ID
	keyPointsResponse.TourId = keyPoint.TourId
	keyPointsResponse.Name = keyPoint.Name
	keyPointsResponse.Description = keyPoint.Description
	keyPointsResponse.Longitude = keyPoint.Longitude
	keyPointsResponse.Latitude = keyPoint.Latitude
	keyPointsResponse.LocationAddress = keyPoint.LocationAddress
	keyPointsResponse.ImagePath = keyPoint.ImagePath
	keyPointsResponse.Order = keyPoint.Order

	ret := &keyPoints.KeyPointGetByIdResponse{
		KeyPoint: &keyPointsResponse,
	}

	return ret, nil
}

func (handler *KeyPointHandler) GetAll(ctx context.Context, request *keyPoints.KeyPointGetAllRequest) (*keyPoints.KeyPointGetAllResponse, error) {
	keyPointList, _ := handler.KeyPointService.GetAll()

	keyPointsResponse := make([]*keyPoints.KeyPoint, len(*keyPointList))

	if keyPointList != nil && len(*keyPointList) > 0 {
		for i, keyPoint := range *keyPointList {
			keyPointsResponse[i] = &keyPoints.KeyPoint{
				ID:              keyPoint.ID,
				TourId:          keyPoint.TourId,
				Name:            keyPoint.Name,
				Description:     keyPoint.Description,
				Longitude:       keyPoint.Longitude,
				Latitude:        keyPoint.Latitude,
				LocationAddress: keyPoint.LocationAddress,
				ImagePath:       keyPoint.ImagePath,
				Order:           keyPoint.Order,
			}
		}
	}

	ret := &keyPoints.KeyPointGetAllResponse{
		KeyPoints: keyPointsResponse,
	}

	return ret, nil
}

func (handler *KeyPointHandler) GetAllByTourId(ctx context.Context, request *keyPoints.KeyPointGetByTourIdRequest) (*keyPoints.KeyPointGetByTourIdResponse, error) {
	keyPointList, _ := handler.KeyPointService.GetAllByTourId(request.TourId)

	keyPointsResponse := make([]*keyPoints.KeyPoint, len(*keyPointList))

	if keyPointList != nil && len(*keyPointList) > 0 {
		for i, keyPoint := range *keyPointList {
			keyPointsResponse[i] = &keyPoints.KeyPoint{
				ID:              keyPoint.ID,
				TourId:          keyPoint.TourId,
				Name:            keyPoint.Name,
				Description:     keyPoint.Description,
				Longitude:       keyPoint.Longitude,
				Latitude:        keyPoint.Latitude,
				LocationAddress: keyPoint.LocationAddress,
				ImagePath:       keyPoint.ImagePath,
				Order:           keyPoint.Order,
			}
		}
	}

	ret := &keyPoints.KeyPointGetByTourIdResponse{
		KeyPoints: keyPointsResponse,
	}

	return ret, nil
}

func (handler *KeyPointHandler) Create(ctx context.Context, request *keyPoints.KeyPointCreateRequest) (*keyPoints.KeyPointCreateResponse, error) {
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

	return &keyPoints.KeyPointCreateResponse{}, nil
}

func (handler *KeyPointHandler) Delete(ctx context.Context, request *keyPoints.KeyPointDeleteRequest) (*keyPoints.KeyPointDeleteResponse, error) {
	handler.KeyPointService.Delete(request.ID)
	return &keyPoints.KeyPointDeleteResponse{}, nil
}

func (handler *KeyPointHandler) Update(ctx context.Context, request *keyPoints.KeyPointUpdateRequest) (*keyPoints.KeyPointUpdateResponse, error) {
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

	return &keyPoints.KeyPointUpdateResponse{}, nil
}
