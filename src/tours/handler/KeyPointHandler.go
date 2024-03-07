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
	err = json.NewEncoder(writer).Encode(keyPoint)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode key points in method GetById"))
		return
	}
}

func (handler *KeyPointHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	keyPoints, err := handler.KeyPointService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(keyPoints)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode key points in method GetAll"))
		return
	}
}

func (handler *KeyPointHandler) GetAllByTourId(writer http.ResponseWriter, req *http.Request) {
	tourId := mux.Vars(req)["tourId"]
	log.Printf("KeyPoint with tour id %s", tourId)
	keyPoints, err := handler.KeyPointService.GetAllByTourId(tourId)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(keyPoints)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode key points in method GetAllByTourId"))
		return
	}
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
