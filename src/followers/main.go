package main

import (
	"context"
	"followers/handler"
	"followers/repo"
	"followers/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8084"
	}

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[user-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[user-store] ", log.LstdFlags)

	store, err := repo.New(storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.CloseDriverConnection(timeoutContext)
	store.CheckConnection()

	userService := service.NewUserService(logger, store)

	userHandler := handler.NewUserHandler(logger, userService)

	router := mux.NewRouter().StrictSlash(true)

	router.Use(userHandler.MiddlewareContentTypeSet)

	postUserNode := router.Methods(http.MethodPost).Subrouter()
	postUserNode.HandleFunc("/users", userHandler.Create)
	postUserNode.Use(userHandler.MiddlewareUserDeserialization)

	postUserFollowNode := router.Methods(http.MethodPost).Subrouter()
	postUserFollowNode.HandleFunc("/users/follow", userHandler.Follow)
	postUserFollowNode.Use(userHandler.MiddlewareUserFollowDeserialization)
	
	getNode := router.Methods(http.MethodGet).Subrouter()
	getNode.HandleFunc("/users/{username}", userHandler.GetByUsername)
	getNode.HandleFunc("/users/followers/{id}", userHandler.GetFollowers)
	getNode.HandleFunc("/users/followings/{id}", userHandler.GetFollowings)

	postNode := router.Methods(http.MethodPost).Subrouter()
	postNode.HandleFunc("/users/unfollow/{followerId}/{followingId}", userHandler.Unfollow)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	logger.Println("Server starting", port)
	logger.Println("Server listening on port", port)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
}
