package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tours/model"
	"tours/service"
)

type ReviewHandler struct {
	ReviewService *service.ReviewService
}

func (handler *ReviewHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Review with id %s", id)
	review, err := handler.ReviewService.GetById(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(review)
}

func (handler *ReviewHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	reviews, err := handler.ReviewService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(reviews)
}

func (handler *ReviewHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var review model.Review
	err := json.NewDecoder(req.Body).Decode(&review)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.ReviewService.Create(&review)
	if err != nil {
		println("Error while creating a new review")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
