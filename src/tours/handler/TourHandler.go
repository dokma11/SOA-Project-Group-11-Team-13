package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"tours/dto"
	"tours/model"
	"tours/service"
)

type TourHandler struct {
	TourService *service.TourService
}

func (handler *TourHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	log.Printf("Tour with id %s", idString)

	tour, err := handler.TourService.GetById(idString)

	var tourDto dto.TourResponseDto
	tourDto.AverageRating = 0.0
	tourDto.Tags = tour.Tags
	tourDto.KeyPoints = tour.KeyPoints
	tourDto.Status = dto.TourStatus(tour.Status)
	tourDto.Name = tour.Name
	tourDto.Description = tour.Description

	tourDto.ID = tour.ID

	tourDto.Durations = tour.Durations
	tourDto.PublishDate = tour.PublishDate
	tourDto.ArchiveDate = tour.ArchiveDate
	tourDto.Category = dto.TourCategory(tour.Category)
	tourDto.IsDeleted = tour.IsDeleted
	tourDto.Price = tour.Price
	tourDto.Distance = tour.Distance
	tourDto.Difficulty = tour.Difficulty
	tourDto.AuthorId = tour.AuthorId

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(tourDto)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode tours in method GetById"))
		return
	}
}

func (handler *TourHandler) GetByAuthorId(writer http.ResponseWriter, req *http.Request) {
	authorIdString := mux.Vars(req)["authorId"]
	log.Printf("Tour with author id %s", authorIdString)

	tour, err := handler.TourService.GetByAuthorId(authorIdString)

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
	idString := mux.Vars(req)["id"]
	log.Printf("Tour with id %s", idString)

	tour := handler.TourService.Delete(idString)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(tour)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode tours in method Delete"))
		return
	}
}

func (handler *TourHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	//err := json.NewDecoder(req.Body).Decode(&tour)
	/*
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println("Error reading request body:", err)
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body, &tour)
	*/
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
	idString := mux.Vars(req)["id"]
	log.Printf("Publish tour with id %s", idString)

	tour := handler.TourService.Publish(idString)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(tour)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode tours in method GetById"))
		return
	}
}
func (handler *TourHandler) Archive(writer http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	log.Printf("Tour with id %s", idString)

	tour := handler.TourService.Archive(idString)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(tour)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode tours in method GetById"))
		return
	}
}
