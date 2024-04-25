package main

import (
	"blogs/handler"
	"blogs/model"
	"blogs/repo"
	"blogs/service"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "host=blogsdb user=postgres password=super dbname=soa-blogs port=5432 sslmode=disable"
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

func initMongoDB() *mongo.Client {
	// Set up MongoDB connection options
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx , options.Client().ApplyURI("mongodb://blogs-mongodb:27017"))
	if err != nil {
		return nil
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	println("Connected to MongoDB!")

	return client
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
	router.HandleFunc("/blogs/authors/{authorIds}", blogHandler.GetByAuthorIds).Methods("GET")
	router.HandleFunc("/blogs/search/{name}", blogHandler.SearchByName).Methods("GET")
	router.HandleFunc("/blogs/publish/{id}", blogHandler.Publish).Methods("PATCH")
	router.HandleFunc("/blogs/{id}", blogHandler.Delete).Methods("DELETE")
}

func initializeCommentRoutes(router *mux.Router, commentHandler *handler.CommentHandler) {
	router.HandleFunc("/comments", commentHandler.GetAll).Methods("GET")
	router.HandleFunc("/comments/{id}", commentHandler.GetById).Methods("GET")
	router.HandleFunc("/comments/byBlog/{id}", commentHandler.GetByBlogId).Methods("GET")
	router.HandleFunc("/comments", commentHandler.Create).Methods("POST")
	router.HandleFunc("/comments/{id}", commentHandler.Delete).Methods("DELETE")
	router.HandleFunc("/comments", commentHandler.Update).Methods("PUT")

}

func initializeVoteRoutes(router *mux.Router, blogHandler *handler.VoteHandler) {
	router.HandleFunc("/votes", blogHandler.GetAll).Methods("GET")
	router.HandleFunc("/votes/{id}", blogHandler.GetById).Methods("GET")
}

func initializeBlogRecommendationRoutes(router *mux.Router, blogRecommendationHandler *handler.BlogRecommendationHandler) {
	router.HandleFunc("/blog/recommendations", blogRecommendationHandler.Create).Methods("POST")
	router.HandleFunc("/blog/recommendations", blogRecommendationHandler.GetAll).Methods("GET")
	router.HandleFunc("/blog/recommendations/{id}", blogRecommendationHandler.GetById).Methods("GET")
	router.HandleFunc("/blog/recommendations/by-receiver/{receiver}", blogRecommendationHandler.GetByReceiverId).Methods("GET")
}

func main() {
	database := initDB()
	blogsMongoDB := initMongoDB();
	if database == nil {
		println("FAILED TO CONNECT TO DB")
		return
	}
	blogRepository := &repo.BlogRepository{DatabaseConnection: blogsMongoDB}
	commentRepository := &repo.CommentRepository{MongoConnection: blogsMongoDB}
	voteRepository := &repo.VoteRepository{DatabaseConnection: database}
	blogRecommendationRepository := &repo.BlogRecommendationRepository{DatabaseConnection: blogsMongoDB}

	blogService := &service.BlogService{BlogRepository: blogRepository, BlogRecommendationRepository: blogRecommendationRepository}
	commentService := &service.CommentService{CommentRepository: commentRepository}
	voteService := &service.VoteService{VoteRepository: voteRepository}
	blogRecommendationService := &service.BlogRecommendationService{BlogRecommendationRepository: blogRecommendationRepository, BlogRepository: blogRepository}

	blogHandler := &handler.BlogHandler{BlogService: blogService}
	commentHandler := &handler.CommentHandler{CommentService: commentService}
	voteHandler := &handler.VoteHandler{VoteService: voteService}
	blogRecommendationHandler := &handler.BlogRecommendationHandler{BlogRecommendationService: blogRecommendationService}

	startServer(blogHandler, commentHandler, voteHandler, blogRecommendationHandler)
}
