package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
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
	tourIdString := mux.Vars(req)["tourId"]
	log.Printf("KeyPoint with tour id %s", tourIdString)
	tourId, err := strconv.ParseInt(tourIdString, 10, 64)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to parse tour id in method GetAllByTourId"))
		return
	}
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

func (handler *KeyPointHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	log.Printf("KeyPoint with id %s", idString)
	id, err := uuid.Parse(idString)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to parse id in method Delete"))
		return
	}
	keyPoint := handler.KeyPointService.Delete(id)
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

func (handler *KeyPointHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var keyPoint model.KeyPoint
	err := json.NewDecoder(req.Body).Decode(&keyPoint)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.KeyPointService.Update(&keyPoint)
	if err != nil {
		println("Error while updating keyPoint")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
