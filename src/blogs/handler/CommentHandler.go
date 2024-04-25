package handler

import (
	"blogs/model"
	"blogs/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CommentHandler struct {
	CommentService *service.CommentService
}

func (handler *CommentHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Comment with id %s", id)

	review, err := handler.CommentService.GetById(id)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(review)
	if err != nil {
		_ = fmt.Errorf("error encountered while trying to encode comments in method GetById")
		return
	}
}

func (handler *CommentHandler) GetByBlogId(writer http.ResponseWriter, req *http.Request) {
	blogId := mux.Vars(req)["id"]
	log.Printf("Comment with blog id %s", blogId)
	page, _ := strconv.Atoi(req.URL.Query().Get("page"))         //novo
	pageSize, _ := strconv.Atoi(req.URL.Query().Get("pageSize")) //novo

	comments, totalCount, err := handler.CommentService.GetByBlogId(blogId, page, pageSize)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	response := struct {
		Comments   []model.Comment `json:"comments"`
		TotalCount int             `json:"totalCount"`
	}{
		Comments:   comments,
		TotalCount: totalCount,
	}

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		log.Printf("Error encountered while trying to encode comments in method GetByBlogId: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *CommentHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Get all blogs")
	tours, err := handler.CommentService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(tours)
	if err != nil {
		_ = fmt.Errorf("error encountered while trying to encode comments in method GetAll")
		return
	}
}

func (handler *CommentHandler) Create(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Creating a Comment")
	var comment model.Comment
	err := json.NewDecoder(req.Body).Decode(&comment)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.CommentService.Create(&comment)

	if err != nil {
		println("Error while creating a new comment")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *CommentHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	log.Printf("Comment with id %s is deleted", idString)

	comment := handler.CommentService.Delete(idString)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(comment)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode comments in method Delete"))
		return
	}
}

func (handler *CommentHandler) Update(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Update Comment")
	var comment model.Comment
	err := json.NewDecoder(req.Body).Decode(&comment)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.CommentService.Update(&comment)

	if err != nil {
		println("Error while updating comment")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
