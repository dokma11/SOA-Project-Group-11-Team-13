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

func (handler *FacilityHandler) GetAll(ctx context.Context, request *facilities.GetAllFacilitiesRequest) (*facilities.GetAllFacilitiesResponse, error) {
	facilityList, _ := handler.FacilityService.GetAll()
	var facilitiesResponse []*facilities.Facility

	for i, f := range *facilityList {
		facilitiesResponse[i].ID = f.ID
		facilitiesResponse[i].AuthorId = f.AuthorId
		facilitiesResponse[i].Name = f.Name
		facilitiesResponse[i].Description = f.Description
		facilitiesResponse[i].Longitude = f.Longitude
		facilitiesResponse[i].Latitude = f.Latitude
		facilitiesResponse[i].Category = facilities.Facility_Category(f.Category)
		facilitiesResponse[i].ImagePath = f.ImagePath
	}

	return &facilities.GetAllFacilitiesResponse{
		Facilities: facilitiesResponse,
	}, nil
}

func (handler *FacilityHandler) GetAllByAuthorId(ctx context.Context, request *facilities.GetFacilitiesByAuthorIdRequest) (*facilities.GetFacilitiesByAuthorIdResponse, error) {
	facilityList, _ := handler.FacilityService.GetAllByAuthorId(request.AuthorId)
	var facilitiesResponse []*facilities.Facility

	for i, f := range *facilityList {
		facilitiesResponse[i].ID = f.ID
		facilitiesResponse[i].AuthorId = f.AuthorId
		facilitiesResponse[i].Name = f.Name
		facilitiesResponse[i].Description = f.Description
		facilitiesResponse[i].Longitude = f.Longitude
		facilitiesResponse[i].Latitude = f.Latitude
		facilitiesResponse[i].Category = facilities.Facility_Category(f.Category)
		facilitiesResponse[i].ImagePath = f.ImagePath
	}

	return &facilities.GetFacilitiesByAuthorIdResponse{
		Facilities: facilitiesResponse,
	}, nil
}

func (handler *FacilityHandler) Create(ctx context.Context, request *facilities.CreateFacilityRequest) (*facilities.CreateFacilityResponse, error) {
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

	return &facilities.CreateFacilityResponse{}, nil
}

func (handler *FacilityHandler) Delete(ctx context.Context, request *facilities.DeleteFacilityRequest) (*facilities.DeleteFacilityResponse, error) {
	handler.FacilityService.Delete(request.ID)
	return &facilities.DeleteFacilityResponse{}, nil
}

func (handler *FacilityHandler) Update(ctx context.Context, request *facilities.UpdateFacilityRequest) (*facilities.UpdateFacilityResponse, error) {
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
