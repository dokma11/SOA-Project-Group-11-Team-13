package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tours/model"
	"tours/service"
)

type TourHandler struct {
	TourService *service.TourService
}

func (handler *TourHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Tour with id %s", id)
	tour, err := handler.TourService.GetById(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(tour)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode tours in method GetById"))
		return
	}
}

func (handler *TourHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	tours, err := handler.TourService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(tours)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode tours in method GetAll"))
		return
	}
}

func (handler *TourHandler) GetPublished(writer http.ResponseWriter, req *http.Request) {
	tours, err := handler.TourService.GetPublished()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(tours)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode tours in method GetPublished"))
		return
	}
}

func (handler *TourHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	err := json.NewDecoder(req.Body).Decode(&tour)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourService.Create(&tour)
	if err != nil {
		println("Error while creating a new tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	log.Printf("Tour with id %s", idString)
	id, err := uuid.Parse(idString)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to parse id in method Delete"))
		return
	}
	keyPoint := handler.TourService.Delete(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(keyPoint)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode tours in method Delete"))
		return
	}
}

func (handler *TourHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	err := json.NewDecoder(req.Body).Decode(&tour)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourService.Update(&tour)
	if err != nil {
		println("Error while updating tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
