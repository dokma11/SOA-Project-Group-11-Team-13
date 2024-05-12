package main

import (
	"jwt/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func startServer(tourHandler *handler.JwtHandler) {
	router := mux.NewRouter().StrictSlash(true)

	initializeJwtRoutes(router, tourHandler)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8085", router))
}

func initializeJwtRoutes(router *mux.Router, jwtHandler *handler.JwtHandler) {
	router.HandleFunc("/jwt", jwtHandler.Create).Methods("GET")
}

func main() {
	jwtHandler := &handler.JwtHandler{}

	startServer(jwtHandler)
}
