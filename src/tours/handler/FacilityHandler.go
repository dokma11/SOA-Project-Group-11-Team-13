package handler

import (
	"context"
	"errors"
	"tours/model"
	"tours/proto/facilities"
	"tours/service"
)

type FacilityHandler struct {
	FacilityService *service.FacilityService
	facilities.UnimplementedFacilitiesServiceServer
}

func (handler *FacilityHandler) GetAllFacilities(ctx context.Context, request *facilities.GetAllFacilitiesRequest) (*facilities.GetAllFacilitiesResponse, error) {
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

	ret := &facilities.GetAllFacilitiesResponse{
		Facilities: facilitiesResponse,
	}

	return ret, nil
}

func (handler *FacilityHandler) GetFacilitiesByAuthorId(ctx context.Context, request *facilities.GetFacilitiesByAuthorIdRequest) (*facilities.GetFacilitiesByAuthorIdResponse, error) {
	facilityList, _ := handler.FacilityService.GetAllByAuthorId(request.AuthorId)

	if facilityList == nil {
		return &facilities.GetFacilitiesByAuthorIdResponse{
			Facilities: []*facilities.Facility{},
		}, nil
	}

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

	ret := &facilities.GetFacilitiesByAuthorIdResponse{
		Facilities: facilitiesResponse,
	}

	return ret, nil
}

func (handler *FacilityHandler) CreateFacility(ctx context.Context, request *facilities.CreateFacilityRequest) (*facilities.CreateFacilityResponse, error) {
	if request == nil {
		return nil, errors.New("nil request")
	}

	if request.Facility == nil {
		return nil, errors.New("nil facility in request")
	}

	facility := model.Facility{}
	facility.ID = request.Facility.ID
	facility.AuthorId = request.Facility.AuthorId
	facility.Name = request.Facility.Name
	facility.Description = request.Facility.Description
	facility.Longitude = request.Facility.Longitude
	facility.Latitude = request.Facility.Latitude
	facility.Category = model.FacilityCategory(request.Facility.Category)
	facility.ImagePath = request.Facility.ImagePath
	
	if err := handler.FacilityService.Create(&facility); err != nil {
		return nil, err
	}

	return &facilities.CreateFacilityResponse{}, nil
}

func (handler *FacilityHandler) DeleteFacility(ctx context.Context, request *facilities.DeleteFacilityRequest) (*facilities.DeleteFacilityResponse, error) {
	handler.FacilityService.Delete(request.ID)
	return &facilities.DeleteFacilityResponse{}, nil
}

func (handler *FacilityHandler) UpdateFacility(ctx context.Context, request *facilities.UpdateFacilityRequest) (*facilities.UpdateFacilityResponse, error) {
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

	return &facilities.UpdateFacilityResponse{}, nil
}
