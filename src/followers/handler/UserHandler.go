package handler

import (
	"context"
	"followers/model"
	"followers/service"
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
