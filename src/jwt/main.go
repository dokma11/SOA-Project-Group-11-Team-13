package main

import (
	"jwt/handler"
	"jwt/repo"
	"jwt/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// func main() {
// 	port := os.Getenv("PORT")
// 	if len(port) == 0 {
// 		port = "8085"
// 	}

// 	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	logger := log.New(os.Stdout, "[user-api] ", log.LstdFlags)
// 	storeLogger := log.New(os.Stdout, "[user-store] ", log.LstdFlags)

// 	store, err := repo.New(storeLogger)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}
// 	defer store.CloseDriverConnection(timeoutContext)
// 	store.CheckConnection()

// 	userService := service.NewUserService(logger, store)

// 	userHandler := handler.NewUserHandler(logger, userService)

// 	router := mux.NewRouter().StrictSlash(true)

// 	router.Use(userHandler.MiddlewareContentTypeSet)

// 	postUserNode := router.Methods(http.MethodPost).Subrouter()
// 	postUserNode.HandleFunc("/users", userHandler.Create)
// 	postUserNode.Use(userHandler.MiddlewareUserDeserialization)

// 	postUserFollowNode := router.Methods(http.MethodPost).Subrouter()
// 	postUserFollowNode.HandleFunc("/users/follow/{userId}/{followedById}", userHandler.Follow)

// 	getNode := router.Methods(http.MethodGet).Subrouter()
// 	getNode.HandleFunc("/users/{username}", userHandler.GetByUsername).Methods("GET")

// 	getFollowersNode := router.Methods(http.MethodGet).Subrouter()
// 	getFollowersNode.HandleFunc("/users/followers/{id}", userHandler.GetFollowers)

// 	getFollowingsNode := router.Methods(http.MethodGet).Subrouter()
// 	getFollowingsNode.HandleFunc("/users/followings/{id}", userHandler.GetFollowings)

// 	getRecommendedNode := router.Methods(http.MethodGet).Subrouter()
// 	getRecommendedNode.HandleFunc("/users/recommended/{id}", userHandler.GetRecommendedUsers)

// 	deleteNode := router.Methods(http.MethodDelete).Subrouter()
// 	deleteNode.HandleFunc("/users/unfollow/{followerId}/{followingId}", userHandler.Unfollow)

// 	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}), gorillaHandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}))
// 	server := http.Server{
// 		Addr:         ":" + port,
// 		Handler:      cors(router),
// 		IdleTimeout:  120 * time.Second,
// 		ReadTimeout:  5 * time.Second,
// 		WriteTimeout: 5 * time.Second,
// 	}

// 	logger.Println("Server starting", port)
// 	logger.Println("Server listening on port", port)
// 	go func() {
// 		err := server.ListenAndServe()
// 		if err != nil {
// 			logger.Fatal(err)
// 		}
// 	}()

// 	sigCh := make(chan os.Signal)
// 	signal.Notify(sigCh, os.Interrupt)
// 	signal.Notify(sigCh, os.Kill)

// 	sig := <-sigCh
// 	logger.Println("Received terminate, graceful shutdown", sig)

// 	if server.Shutdown(timeoutContext) != nil {
// 		logger.Fatal("Cannot gracefully shutdown...")
// 	}
// 	logger.Println("Server stopped")
// }

func startServer(tourHandler *handler.TourHandler) {
	router := mux.NewRouter().StrictSlash(true)

	initializeTourRoutes(router, tourHandler)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8085", router))
}

func initializeTourRoutes(router *mux.Router, tourHandler *handler.TourHandler) {
	router.HandleFunc("/tours/published", tourHandler.GetPublished).Methods("GET")
	router.HandleFunc("/tours/{id}", tourHandler.GetById).Methods("GET")
	router.HandleFunc("/tours/authors/{authorId}", tourHandler.GetByAuthorId).Methods("GET")
	router.HandleFunc("/tours", tourHandler.GetAll).Methods("GET")
	router.HandleFunc("/tours/{tourId}/equipment", tourHandler.GetEquipment).Methods("GET")

	router.HandleFunc("/tours", tourHandler.Create).Methods("POST")
	router.HandleFunc("/tours/{tourId}/equipment/{equipmentId}", tourHandler.AddEquipment).Methods("POST")

	router.HandleFunc("/tours", tourHandler.Update).Methods("PUT")
	router.HandleFunc("/tours/durations", tourHandler.AddDurations).Methods("PUT")
	router.HandleFunc("/tours/publish/{id}", tourHandler.Publish).Methods("PUT")
	router.HandleFunc("/tours/archive/{id}", tourHandler.Archive).Methods("PUT")

	router.HandleFunc("/tours/{id}", tourHandler.Delete).Methods("DELETE")
	router.HandleFunc("/tours/{tourId}/equipment/{equipmentId}", tourHandler.DeleteEquipment).Methods("DELETE")
}

func main() {
	tourRepository := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepository: tourRepository}
	tourHandler := &handler.TourHandler{TourService: tourService}

	startServer(tourHandler, keyPointHandler, reviewHandler, equipmentHandler, facilityHandler)
}
