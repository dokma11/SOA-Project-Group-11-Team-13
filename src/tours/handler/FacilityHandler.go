package handler

import (
	"context"
	"tours/model"
	"tours/proto/facilities"
	"tours/service"
)

type FacilityHandler struct {
	FacilityService *service.FacilityService
	facilities.UnimplementedFacilitiesServiceServer
}

func (handler *FacilityHandler) GetAll(ctx context.Context, request *facilities.GetAllRequest) (*facilities.GetAllResponse, error) {
	facilityList, _ := handler.FacilityService.GetAll()

	facilitiesResponse := make([]*facilities.Facility, len(*facilityList))

	if facilityList != nil && len(*facilityList) > 0 {
		for i, f := range *facilityList {
			facilitiesResponse[i] = &facilities.Facility{
				ID:          f.ID,
				AuthorId:    f.AuthorId,
				Name:        f.Name,
				Description: f.Description,
				Longitude:   f.Longitude,
				Latitude:    f.Latitude,
				Category:    facilities.Facility_Category(f.Category),
				ImagePath:   f.ImagePath,
			}
		}
	}

	ret := &facilities.GetAllResponse{
		Facilities: facilitiesResponse,
	}

	return ret, nil
}

func (handler *FacilityHandler) GetAllByAuthorId(ctx context.Context, request *facilities.GetByAuthorIdRequest) (*facilities.GetByAuthorIdResponse, error) {
	facilityList, _ := handler.FacilityService.GetAllByAuthorId(request.AuthorId)

	facilitiesResponse := make([]*facilities.Facility, len(*facilityList))

	if facilityList != nil && len(*facilityList) > 0 {
		for i, f := range *facilityList {
			facilitiesResponse[i] = &facilities.Facility{
				ID:          f.ID,
				AuthorId:    f.AuthorId,
				Name:        f.Name,
				Description: f.Description,
				Longitude:   f.Longitude,
				Latitude:    f.Latitude,
				Category:    facilities.Facility_Category(f.Category),
				ImagePath:   f.ImagePath,
			}
		}
	}

	ret := &facilities.GetByAuthorIdResponse{
		Facilities: facilitiesResponse,
	}

	return ret, nil
}

func (handler *FacilityHandler) Create(ctx context.Context, request *facilities.CreateRequest) (*facilities.CreateResponse, error) {
	facility := model.Facility{}
	facility.ID = request.Facility.ID
	facility.AuthorId = request.Facility.AuthorId
	facility.Name = request.Facility.Name
	facility.Description = request.Facility.Description
	facility.Longitude = request.Facility.Longitude
	facility.Latitude = request.Facility.Latitude
	facility.Category = model.FacilityCategory(request.Facility.Category)
	facility.ImagePath = request.Facility.ImagePath

	handler.FacilityService.Create(&facility)

	return &facilities.CreateResponse{}, nil
}

func (handler *FacilityHandler) Delete(ctx context.Context, request *facilities.DeleteRequest) (*facilities.DeleteResponse, error) {
	handler.FacilityService.Delete(request.ID)
	return &facilities.DeleteResponse{}, nil
}

func (handler *FacilityHandler) Update(ctx context.Context, request *facilities.UpdateRequest) (*facilities.UpdateResponse, error) {
	facility := model.Facility{}
	facility.ID = request.Facility.ID
	facility.AuthorId = request.Facility.AuthorId
	facility.Name = request.Facility.Name
	facility.Description = request.Facility.Description
	facility.Longitude = request.Facility.Longitude
	facility.Latitude = request.Facility.Latitude
	facility.Category = model.FacilityCategory(request.Facility.Category)
	facility.ImagePath = request.Facility.ImagePath

	handler.FacilityService.Update(&facility)

	return &facilities.UpdateResponse{}, nil
}
