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
	err = json.NewEncoder(writer).Encode(review)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode reviews in method GetById"))
		return
	}
}

func (handler *ReviewHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	reviews, err := handler.ReviewService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(reviews)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode reviews in method GetAll"))
		return
	}
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

func (handler *ReviewHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Review with id %s", id)

	review := handler.ReviewService.Delete(id)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(review)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode key points in method GetById"))
		return
	}
}

func (handler *ReviewHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var review model.Review
	err := json.NewDecoder(req.Body).Decode(&review)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.ReviewService.Update(&review)

	if err != nil {
		println("Error while updating review")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
