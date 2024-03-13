package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tours/model"
	"tours/service"

	"github.com/gorilla/mux"
)

type TourHandler struct {
	TourService *service.TourService
}

func (handler *TourHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	log.Printf("Tour with id %s", idString)

	tour, err := handler.TourService.GetById(idString)

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

func (handler *TourHandler) GetByAuthorId(writer http.ResponseWriter, req *http.Request) {
	authorId := mux.Vars(req)["authorId"]
	log.Printf("Tour with author id %s", authorId)

	tour, err := handler.TourService.GetByAuthorId(authorId)

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
	log.Printf("Get all tours")
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
	log.Printf("Get published tours")
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
	log.Printf("Create a tour")
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Request Body:", string(body))
	var tour model.Tour

	err = json.Unmarshal(body, &tour)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		http.Error(writer, "Invalid JSON", http.StatusBadRequest)
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
	id := mux.Vars(req)["id"]
	log.Printf("Tour with id %s", id)

	tour := handler.TourService.Delete(id)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(tour)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode tours in method Delete"))
		return
	}
}

func (handler *TourHandler) Update(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Update a tour")
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

func (handler *TourHandler) AddDurations(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Add durations to a tour")
	var tour model.Tour
	err := json.NewDecoder(req.Body).Decode(&tour)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourService.AddDurations(&tour)

	if err != nil {
		println("Error while updating tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourHandler) Publish(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Publish tour with id %s", id)

	tour := handler.TourService.Publish(id)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(tour)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode tours in method GetById"))
		return
	}
}
func (handler *TourHandler) Archive(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Tour with id %s", id)

	tour := handler.TourService.Archive(id)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(tour)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode tours in method GetById"))
		return
	}
}

func (handler *TourHandler) GetEquipment(writer http.ResponseWriter, req *http.Request) {
	tourId := mux.Vars(req)["tourId"]
	log.Printf("Equipment for tour with id %s", tourId)

	equipmentList, err := handler.TourService.GetEquipment(tourId)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(equipmentList)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode equipment in method GetEquipment"))
		return
	}
}

func (handler *TourHandler) AddEquipment(writer http.ResponseWriter, req *http.Request) {
	tourId := mux.Vars(req)["tourId"]
	equipmentId := mux.Vars(req)["equipmentId"]
	log.Printf("Adding equipment with id %s to tour with id %s", equipmentId, tourId)

	err := handler.TourService.AddEquipment(tourId, equipmentId)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *TourHandler) DeleteEquipment(writer http.ResponseWriter, req *http.Request) {
	tourId := mux.Vars(req)["tourId"]
	equipmentId := mux.Vars(req)["equipmentId"]
	log.Printf("Deleting equipment with id %s from tour with id %s", equipmentId, tourId)

	err := handler.TourService.DeleteEquipment(tourId, equipmentId)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
