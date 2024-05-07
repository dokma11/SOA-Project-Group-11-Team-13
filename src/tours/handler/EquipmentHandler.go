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

func (handler *EquipmentHandler) GetById(ctx context.Context, request *equipment.GetByIdRequest) (*equipment.GetByIdResponse, error) {
	equipmentList, _ := handler.EquipmentService.GetById(request.ID)

	equipmentResponse := equipment.Equipment{}
	equipmentResponse.ID = equipmentList.ID
	equipmentResponse.Name = equipmentList.Name
	equipmentResponse.Description = equipmentList.Description
	//equipmentResponse.Tours = equipmentList.Tours				Treba proveriti

	ret := &equipment.GetByIdResponse{
		Equipment: &equipmentResponse,
	}

	return ret, nil
}

func (handler *EquipmentHandler) GetAll(ctx context.Context, request *equipment.GetAllRequest) (*equipment.GetAllResponse, error) {
	equipmentList, _ := handler.EquipmentService.GetAll()

	equipmentResponse := make([]*equipment.Equipment, len(*equipmentList))

	if equipmentList != nil && len(*equipmentList) > 0 {
		for i, eq := range *equipmentList {
			equipmentResponse[i] = &equipment.Equipment{
				ID:          eq.ID,
				Name:        eq.Name,
				Description: eq.Description,
				//Tours: eq.Tours,	Treba proveriti
			}
		}
	}

	ret := &equipment.GetAllResponse{
		Equipment: equipmentResponse,
	}

	return ret, nil
}

func (handler *EquipmentHandler) Create(ctx context.Context, request *equipment.CreateRequest) (*equipment.CreateResponse, error) {
	equipmentResponse := model.Equipment{}
	equipmentResponse.ID = request.Equipment.ID
	equipmentResponse.Name = request.Equipment.Name
	equipmentResponse.Description = request.Equipment.Description
	//equipmentResponse.Tours = request.Equipment.Tours				Treba proveriti

	handler.EquipmentService.Create(&equipmentResponse)

	return &equipment.CreateResponse{}, nil
}
