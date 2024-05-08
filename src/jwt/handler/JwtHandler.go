package handler

import (
	"jwt/model"
	"jwt/service"
	"log"
	"net/http"
)

type JwtHandler struct {
	UserService *service.UserService
	logger      *log.Logger
}

type KeyProduct struct{}

func NewUserHandler(l *log.Logger, s *service.UserService) *JwtHandler {
	return &JwtHandler{s, l}
}

func (handler *JwtHandler) Create(rw http.ResponseWriter, h *http.Request) {
	user := h.Context().Value(KeyProduct{}).(*model.User)
	err := handler.UserService.Create(user)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}
