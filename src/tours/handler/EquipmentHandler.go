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

func (handler *EquipmentHandler) GetById(ctx context.Context, request *equipment.EquipmentGetByIdRequest) (*equipment.EquipmentGetByIdResponse, error) {
	equipmentList, _ := handler.EquipmentService.GetById(request.ID)

	equipmentResponse := equipment.Equipment{}
	equipmentResponse.ID = equipmentList.ID
	equipmentResponse.Name = equipmentList.Name
	equipmentResponse.Description = equipmentList.Description

	ret := &equipment.EquipmentGetByIdResponse{
		Equipment: &equipmentResponse,
	}

	return ret, nil
}

func (handler *EquipmentHandler) GetAll(ctx context.Context, request *equipment.EquipmentGetAllRequest) (*equipment.EquipmentGetAllResponse, error) {
	equipmentList, _ := handler.EquipmentService.GetAll()

	equipmentResponse := make([]*equipment.Equipment, len(*equipmentList))

	if equipmentList != nil && len(*equipmentList) > 0 {
		for i, eq := range *equipmentList {
			equipmentResponse[i] = &equipment.Equipment{
				ID:          eq.ID,
				Name:        eq.Name,
				Description: eq.Description,
			}
		}
	}

	ret := &equipment.EquipmentGetAllResponse{
		Equipment: equipmentResponse,
	}

	return ret, nil
}

func (handler *EquipmentHandler) Create(ctx context.Context, request *equipment.EquipmentCreateRequest) (*equipment.EquipmentCreateResponse, error) {
	equipmentResponse := model.Equipment{}
	equipmentResponse.ID = request.Equipment.ID
	equipmentResponse.Name = request.Equipment.Name
	equipmentResponse.Description = request.Equipment.Description

	handler.EquipmentService.Create(&equipmentResponse)

	return &equipment.EquipmentCreateResponse{}, nil
}
