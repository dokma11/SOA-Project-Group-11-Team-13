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
	blogs, err := handler.BlogService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(blogs)
	if err != nil {
		_ = fmt.Errorf("error encountered while trying to encode blogs in method GetAll")
		return
	}
}

func (handler *BlogHandler) Create(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Create blog")
	var blog model.Blog
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.BlogService.Create(&blog)
	if err != nil {
		println("Error while creating a new blog")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *BlogHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	log.Printf("usao")
	id := mux.Vars(req)["id"]
	log.Printf("Delete blog with id %s", id)

	blog := handler.BlogService.Delete(id)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(blog)
	if err != nil {
		_ = fmt.Errorf("error encountered while trying to encode blogs in method Delete")
		return
	}

}

func (handler *BlogHandler) SearchByName(writer http.ResponseWriter, req *http.Request) {
	name := mux.Vars(req)["name"]
	log.Printf("Searching blogs with name " + name)
	blogs, err := handler.BlogService.SearchByName(name)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(blogs)
	if err != nil {
		_ = fmt.Errorf("error encountered while trying to encode blogs in method GetAll")
		return
	}
}

func (handler *BlogHandler) Publish(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Publish blog with id %s", id)

	blog, err := handler.BlogService.Publish(id)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(blog)
	if err != nil {
		_ = fmt.Errorf("error encountered while trying to encode blogs in method Publish")
		return
	}
}
