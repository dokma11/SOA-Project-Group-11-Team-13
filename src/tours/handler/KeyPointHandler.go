package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"tours/model"
	"tours/service"

	"github.com/gorilla/mux"
)

type KeyPointHandler struct {
	KeyPointService *service.KeyPointService
}

func (handler *KeyPointHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("KeyPoint with id %s", id)
	keyPoint, err := handler.KeyPointService.GetById(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(keyPoint)
}

func (handler *KeyPointHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("KeyPoint with id %s", id)
	keyPoints, err := handler.KeyPointService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(keyPoints)
}

func (handler *KeyPointHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var keyPoint model.KeyPoint
	err := json.NewDecoder(req.Body).Decode(&keyPoint)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.KeyPointService.Create(&keyPoint)
	if err != nil {
		println("Error while creating a new keyPoint")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
