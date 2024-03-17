package handler

import (
	"blogs/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type BlogHandler struct {
	BlogService *service.BlogService
}

func (handler *BlogHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Blog with id %s", id)

	review, err := handler.BlogService.GetById(id)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(review)
	if err != nil {
		_ = fmt.Errorf("error encountered while trying to encode blogs in method GetById")
		return
	}
}

func (handler *BlogHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Get all blogs")
	tours, err := handler.BlogService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(tours)
	if err != nil {
		_ = fmt.Errorf("error encountered while trying to encode blogs in method GetAll")
		return
	}
}
