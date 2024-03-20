package handler

import (
	"blogs/model"
	"blogs/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type BlogRecommendationHandler struct {
	BlogRecommendationService *service.BlogRecommendationService
}

func (handler *BlogRecommendationHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Blog recommendation with id %s", id)

	review, err := handler.BlogRecommendationService.GetById(id)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(review)
	if err != nil {
		_ = fmt.Errorf("error encountered while trying to encode blog recommendations in method GetById")
		return
	}
}

func (handler *BlogRecommendationHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Get all blog recommendations")
	blogRecommendations, err := handler.BlogRecommendationService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(blogRecommendations)
	if err != nil {
		_ = fmt.Errorf("error encountered while trying to encode blog recommendations in method GetAll")
		return
	}
}


func (handler *BlogRecommendationHandler) Create(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Create blog recommendation")
	var blogRecommendation model.BlogRecommendation
	err := json.NewDecoder(req.Body).Decode(&blogRecommendation)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.BlogRecommendationService.Create(&blogRecommendation)
	if err != nil {
		println("Error while creating a new blog recommendation")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}