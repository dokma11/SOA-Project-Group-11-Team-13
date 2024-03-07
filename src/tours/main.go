package main

import (
	"fmt"
	"log"
	"net/http"
	"tours/handler"
	"tours/model"
	"tours/repo"
	"tours/service"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "host=localhost user=postgres password=super dbname=soa-gorm port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	err = database.AutoMigrate(&model.Tour{})
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("tour auto migrations failed"))
		return nil
	}

	err = database.AutoMigrate(&model.KeyPoint{})
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("key point auto migrations failed"))
		return nil
	}

	err = database.AutoMigrate(&model.Review{})
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("review auto migrations failed"))
		return nil
	}

	err = database.AutoMigrate(&model.Equipment{})
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("equipment auto migrations failed"))
		return nil
	}

	return database
}

func startServer(tourHandler *handler.TourHandler, keyPointHandler *handler.KeyPointHandler,
	reviewHandler *handler.ReviewHandler, equipmentHandler *handler.EquipmentHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tours/{id}", tourHandler.GetById).Methods("GET")
	router.HandleFunc("/tours", tourHandler.Create).Methods("POST")
	router.HandleFunc("/tours", tourHandler.GetAll).Methods("GET")
	router.HandleFunc("/tours/published", tourHandler.GetPublished).Methods("GET")
	router.HandleFunc("/tours/{id}", tourHandler.Delete).Methods("DELETE")
	router.HandleFunc("/tours", tourHandler.Update).Methods("UPDATE")

	router.HandleFunc("/keyPoints/{id}", keyPointHandler.GetById).Methods("GET")
	router.HandleFunc("/keyPoints", keyPointHandler.Create).Methods("POST")
	router.HandleFunc("/keyPoints", keyPointHandler.GetAll).Methods("GET")
	router.HandleFunc("/keyPoints/{tourId}", keyPointHandler.GetAllByTourId).Methods("GET")
	router.HandleFunc("/keyPoints/{id}", keyPointHandler.Delete).Methods("DELETE")
	router.HandleFunc("/keyPoints", keyPointHandler.Update).Methods("UPDATE")

	router.HandleFunc("/reviews/{id}", reviewHandler.GetById).Methods("GET")
	router.HandleFunc("/reviews", reviewHandler.Create).Methods("POST")
	router.HandleFunc("/reviews", reviewHandler.GetAll).Methods("GET")
	router.HandleFunc("/reviews/{id}", reviewHandler.Delete).Methods("DELETE")
	router.HandleFunc("/reviews", reviewHandler.Update).Methods("UPDATE")

	router.HandleFunc("/equipment/{id}", equipmentHandler.GetById).Methods("GET")
	router.HandleFunc("/equipment", equipmentHandler.Create).Methods("POST")
	router.HandleFunc("/equipment", equipmentHandler.GetAll).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	tourRepository := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepository: tourRepository}
	tourHandler := &handler.TourHandler{TourService: tourService}

	keyPointRepository := &repo.KeyPointRepository{DatabaseConnection: database}
	keyPointService := &service.KeyPointService{KeyPointRepository: keyPointRepository}
	keyPointHandler := &handler.KeyPointHandler{KeyPointService: keyPointService}

	reviewRepository := &repo.ReviewRepository{DatabaseConnection: database}
	reviewService := &service.ReviewService{ReviewRepository: reviewRepository}
	reviewHandler := &handler.ReviewHandler{ReviewService: reviewService}

	equipmentRepository := &repo.EquipmentRepository{DatabaseConnection: database}
	equipmentService := &service.EquipmentService{EquipmentRepository: equipmentRepository}
	equipmentHandler := &handler.EquipmentHandler{EquipmentService: equipmentService}

	startServer(tourHandler, keyPointHandler, reviewHandler, equipmentHandler)
}
