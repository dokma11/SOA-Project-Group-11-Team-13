package handler

import (
	"context"
	"tours/model"
	"tours/proto/equipment"
	"tours/service"
)

type EquipmentHandler struct {
	EquipmentService *service.EquipmentService
	equipment.UnimplementedEquipmentServiceServer
}

func (handler *EquipmentHandler) GetById(ctx context.Context, request *equipment.GetEquipmentByIdRequest) (*equipment.GetEquipmentByIdResponse, error) {
	equipmentList, _ := handler.EquipmentService.GetById(request.ID)

	equipmentResponse := equipment.Equipment{}
	equipmentResponse.ID = equipmentList.ID
	equipmentResponse.Name = equipmentList.Name
	equipmentResponse.Description = equipmentList.Description
	//equipmentResponse.Tours = equipmentList.Tours				Treba proveriti

	return &equipment.GetEquipmentByIdResponse{
		Equipment: &equipmentResponse,
	}, nil
}

func (handler *EquipmentHandler) GetAll(ctx context.Context, request *equipment.GetAllEquipmentRequest) (*equipment.GetAllEquipmentResponse, error) {
	equipmentList, _ := handler.EquipmentService.GetAll()
	var equipmentResponse []*equipment.Equipment

	for i, eq := range *equipmentList {
		equipmentResponse[i].ID = eq.ID
		equipmentResponse[i].Name = eq.Name
		equipmentResponse[i].Description = eq.Description
		//equipmentResponse[i].Tours = e.Tours				Treba proveriti
	}

	return &equipment.GetAllEquipmentResponse{
		Equipment: equipmentResponse,
	}, nil
}

func (handler *EquipmentHandler) Create(ctx context.Context, request *equipment.CreateEquipmentRequest) (*equipment.CreateEquipmentResponse, error) {
	equipmentResponse := model.Equipment{}
	equipmentResponse.ID = request.Equipment.ID
	equipmentResponse.Name = request.Equipment.Name
	equipmentResponse.Description = request.Equipment.Description
	//equipmentResponse.Tours = request.Equipment.Tours				Treba proveriti

	handler.EquipmentService.Create(&equipmentResponse)

	return &equipment.CreateEquipmentResponse{}, nil
}
