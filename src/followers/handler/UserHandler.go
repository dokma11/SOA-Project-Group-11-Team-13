package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"followers/model"
	"followers/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type UserHandler struct {
	UserService *service.UserService
	logger      *log.Logger
}

type KeyProduct struct{}

func NewUserHandler(l *log.Logger, s *service.UserService) *UserHandler {
	return &UserHandler{s, l}
}

func (handler *UserHandler) Create(rw http.ResponseWriter, h *http.Request) {
	user := h.Context().Value(KeyProduct{}).(*model.User)
	err := handler.UserService.Create(user)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func (handler *UserHandler) FollowUser(rw http.ResponseWriter, h *http.Request) {
	userList := h.Context().Value(KeyProduct{}).([]*model.User)
	err := handler.UserService.FollowUser(userList[0], userList[1])
	if err != nil {
		handler.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func (handler *UserHandler) GetFollowers(writer http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["userId"]
	log.Printf("User with id %s", userId)

	followers, err := handler.UserService.GetFollowers(userId)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(followers)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode followers in method GetFollowers"))
		return
	}
}

func (handler *UserHandler) GetFollowings(writer http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["userId"]
	log.Printf("User with id %s", userId)

	followers, err := handler.UserService.GetFollowings(userId)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(followers)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode followings in method GetFollowings"))
		return
	}
}

func (handler *UserHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		handler.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}

func (handler *UserHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		user := &model.User{}
		err := user.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}
		ctx := context.WithValue(h.Context(), KeyProduct{}, user)
		h = h.WithContext(ctx)
		next.ServeHTTP(rw, h)
	})
}

func (handler *UserHandler) MiddlewareUserFollowDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		var users []*model.User
		err := json.NewDecoder(h.Body).Decode(&users)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}
		ctx := context.WithValue(h.Context(), KeyProduct{}, users)
		h = h.WithContext(ctx)
		next.ServeHTTP(rw, h)
	})
}
