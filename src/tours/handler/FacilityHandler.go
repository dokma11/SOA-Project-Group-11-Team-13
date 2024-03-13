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

type FacilityHandler struct {
	FacilityService *service.FacilityService
}

func (handler *FacilityHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Get all facilities")
	facilities, err := handler.FacilityService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(facilities)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode facilities in method GetAll"))
		return
	}
}

func (handler *FacilityHandler) GetAllByAuthorId(writer http.ResponseWriter, req *http.Request) {
	authorId := mux.Vars(req)["authorId"]
	log.Printf("Facility with author id %s", authorId)

	facilities, err := handler.FacilityService.GetAllByAuthorId(authorId)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(facilities)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode facilities in method GetAllByAuthorId"))
		return
	}
}

func (handler *FacilityHandler) Create(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Create a facility")
	var facility model.Facility
	err := json.NewDecoder(req.Body).Decode(&facility)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.FacilityService.Create(&facility)

	if err != nil {
		println("Error while creating a new facility")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *FacilityHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Facility with id %s", id)

	facility := handler.FacilityService.Delete(id)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(facility)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode facility in method Delete"))
		return
	}
}

func (handler *FacilityHandler) Update(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Update a facility")
	var facility model.Facility
	err := json.NewDecoder(req.Body).Decode(&facility)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.FacilityService.Update(&facility)

	if err != nil {
		println("Error while updating facility")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
