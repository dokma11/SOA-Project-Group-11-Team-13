package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tours/model"
	"tours/service"
)

type EquipmentHandler struct {
	EquipmentService *service.EquipmentService
}

func (handler *EquipmentHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Equipment with id %s", id)
	equipment, err := handler.EquipmentService.GetById(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(equipment)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode equipment"))
		return
	}
}

func (handler *EquipmentHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	equipment, err := handler.EquipmentService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(equipment)
	if err != nil {
		return
	}
}

func (handler *EquipmentHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var equipment model.Equipment
	err := json.NewDecoder(req.Body).Decode(&equipment)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EquipmentService.Create(&equipment)
	if err != nil {
		println("Error while creating a new equipment")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
