package main

import (
	"blogs/handler"
	"blogs/model"
	"blogs/repo"
	"blogs/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "host=localhost user=postgres password=super dbname=soa-blogs port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	err = database.AutoMigrate(&model.Blog{})
	if err != nil {
		_ = fmt.Errorf("blog auto migrations failed")
		return nil
	}

	err = database.AutoMigrate(&model.Comment{})
	if err != nil {
		_ = fmt.Errorf("comment auto migrations failed")
		return nil
	}

	err = database.AutoMigrate(&model.Vote{})
	if err != nil {
		_ = fmt.Errorf("vote auto migrations failed")
		return nil
	}

	err = database.AutoMigrate(&model.BlogRecommendation{})
	if err != nil {
		_ = fmt.Errorf("blog_recommendation auto migrations failed")
		return nil
	}

	return database
}

func startServer(blogHandler *handler.BlogHandler, commentHandler *handler.CommentHandler, voteHandler *handler.VoteHandler, blogRecommendationHandler *handler.BlogRecommendationHandler) {
	router := mux.NewRouter().StrictSlash(true)

	initializeBlogRoutes(router, blogHandler)
	initializeCommentRoutes(router, commentHandler)
	initializeVoteRoutes(router, voteHandler)
	initializeBlogRecommendationRoutes(router, blogRecommendationHandler)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	println("Server starting")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func initializeBlogRoutes(router *mux.Router, blogHandler *handler.BlogHandler) {
	router.HandleFunc("/blogs", blogHandler.Create).Methods("POST")
	router.HandleFunc("/blogs", blogHandler.GetAll).Methods("GET")
	router.HandleFunc("/blogs/{id}", blogHandler.GetById).Methods("GET")
	router.HandleFunc("/blogs/search/{name}", blogHandler.SearchByName).Methods("GET")
}

func initializeCommentRoutes(router *mux.Router, commentHandler *handler.CommentHandler) {
	router.HandleFunc("/comments", commentHandler.GetAll).Methods("GET")
	router.HandleFunc("/comments/{id}", commentHandler.GetById).Methods("GET")
}

func initializeVoteRoutes(router *mux.Router, blogHandler *handler.VoteHandler) {
	router.HandleFunc("/votes", blogHandler.GetAll).Methods("GET")
	router.HandleFunc("/votes/{id}", blogHandler.GetById).Methods("GET")
}

func initializeBlogRecommendationRoutes(router *mux.Router, blogRecommendationHandler *handler.BlogRecommendationHandler) {
	router.HandleFunc("/blog-recommendations", blogRecommendationHandler.Create).Methods("POST")
	router.HandleFunc("/blog-recommendations", blogRecommendationHandler.GetAll).Methods("GET")
	router.HandleFunc("/blog-recommendations/{id}", blogRecommendationHandler.GetById).Methods("GET")
}

func main() {
	database := initDB()
	if database == nil {
		println("FAILED TO CONNECT TO DB")
		return
	}
	blogRepository := &repo.BlogRepository{DatabaseConnection: database}
	commentRepository := &repo.CommentRepository{DatabaseConnection: database}
	voteRepository := &repo.VoteRepository{DatabaseConnection: database}
	blogRecommendationRepository := &repo.BlogRecommendationRepository{DatabaseConnection: database}

	blogService := &service.BlogService{BlogRepository: blogRepository}
	commentService := &service.CommentService{CommentRepository: commentRepository}
	voteService := &service.VoteService{VoteRepository: voteRepository}
	blogRecommendationService := &service.BlogRecommendationService{BlogRecommendationRepository: blogRecommendationRepository}

	blogHandler := &handler.BlogHandler{BlogService: blogService}
	commentHandler := &handler.CommentHandler{CommentService: commentService}
	voteHandler := &handler.VoteHandler{VoteService: voteService}
	blogRecommendationHandler := &handler.BlogRecommendationHandler{BlogRecommendationService: blogRecommendationService}

	startServer(blogHandler, commentHandler, voteHandler, blogRecommendationHandler)
}
