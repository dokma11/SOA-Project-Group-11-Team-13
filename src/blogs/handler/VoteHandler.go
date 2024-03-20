package handler

import (
	"blogs/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type VoteHandler struct {
	VoteService *service.VoteService
}

func (handler *VoteHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Vote with id %s", id)

	review, err := handler.VoteService.GetById(id)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(review)
	if err != nil {
		_ = fmt.Errorf("error encountered while trying to encode votes in method GetById")
		return
	}
}

func (handler *VoteHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Get all votes")
	tours, err := handler.VoteService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(tours)
	if err != nil {
		_ = fmt.Errorf("error encountered while trying to encode votes in method GetAll")
		return
	}
}
