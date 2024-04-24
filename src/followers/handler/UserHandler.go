package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"followers/model"
	"followers/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func (handler *UserHandler) Follow(rw http.ResponseWriter, h *http.Request) {
	userId := mux.Vars(h)["userId"]
	followedById := mux.Vars(h)["followedById"]

	err := handler.UserService.Follow(userId, followedById)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func (handler *UserHandler) Unfollow(writer http.ResponseWriter, req *http.Request) {
	followerId := mux.Vars(req)["followerId"]
	followingId := mux.Vars(req)["followingId"]

	fmt.Println()

	err := handler.UserService.Unfollow(followerId, followingId)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (handler *UserHandler) GetFollowers(writer http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["id"]
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
	userId := mux.Vars(req)["id"]
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

func (handler *UserHandler) GetByUsername(writer http.ResponseWriter, req *http.Request) {
	username := mux.Vars(req)["username"]
	log.Printf("User with username %s", username)

	user, err := handler.UserService.GetByUsername(username)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(user)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode user in method GetByUsername"))
		return
	}
}

func (handler *UserHandler) GetRecommendedUsers(writer http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["id"]
	log.Printf("User with username %s", userId)

	users, err := handler.UserService.GetRecommendedUsers(userId)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(users)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("error encountered while trying to encode user in method GetByUsername"))
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
