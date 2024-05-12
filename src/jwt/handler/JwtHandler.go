package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtHandler struct { }

type KeyProduct struct{}

func NewUserHandler(l *log.Logger) *JwtHandler {
	return &JwtHandler{}
}

func (handler *JwtHandler) Create(writer http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()

	userId := q.Get("id")
	username := q.Get("username")
	personId := q.Get("personId")
	role := q.Get("role")

	jti := uuid.New().String()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userId,
		"username": username,
		"personId": personId,
		"exp": time.Now().Add(time.Hour * 24 * 100).Unix(),
		"iss": "explorer",
		"aud": "explorer-front.com",
		"http://schemas.microsoft.com/ws/2008/06/identity/claims/role": role,
		"jti": jti,
	})

	tokenStr, err := token.SignedString([]byte("explorer_secret_key"))

	if err != nil {
		return
	}

	log.Println("Generating JWT for user with id=" + userId + ", username=" + username + ", personId=" + personId);
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(tokenStr)
	if err != nil {
		_ = fmt.Errorf("error encountered while trying to encode jwt in method Create")
		return
	}
}
