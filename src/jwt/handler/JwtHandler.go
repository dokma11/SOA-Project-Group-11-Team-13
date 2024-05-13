package handler

import (
	//"encoding/json"
	"context"
	"fmt"
	"log"

	//"net/http"
	jwtPb "jwt/proto/jwt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtHandler struct {
	jwtPb.UnimplementedJwtServiceServer // Embed the unimplemented interface
}

func NewJwtHandler() *JwtHandler {
	return &JwtHandler{}
}

func (handler *JwtHandler) GenerateToken(ctx context.Context, req *jwtPb.GenerateTokenRequest) (*jwtPb.GenerateTokenResponse, error) {
	jti := uuid.New().String()

	// Define your JWT claims
	claims := jwt.MapClaims{
		"id":       req.UserId,
		"username": req.Username,
		"personId": req.PersonId,
		"exp":      time.Now().Add(time.Hour * 24 * 100).Unix(),
		"iss":      "explorer",
		"aud":      "explorer-front.com",
		"http://schemas.microsoft.com/ws/2008/06/identity/claims/role": req.Role,
		"jti": jti,
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte("explorer_secret_key"))
	if err != nil {
		return nil, fmt.Errorf("error generating JWT token: %v", err)
	}

	log.Printf("Generating JWT for user with id=%s, username=%s, personId=%s", req.UserId, req.Username, req.PersonId)

	// Return the generated JWT token in the gRPC response
	return &jwtPb.GenerateTokenResponse{
		Token: tokenStr,
	}, nil
}

/*
type KeyProduct struct{}


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

*/
