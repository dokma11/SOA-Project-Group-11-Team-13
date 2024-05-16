package main

import (
	"blogs/handler"
	"blogs/model"
	"blogs/proto/blog_recommendations"
	"blogs/proto/blogs"
	"blogs/proto/comments"
	"blogs/proto/votes"
	"blogs/repo"
	"blogs/service"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"time"
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
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://blogs-mongodb:27017"))
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

func main() {
	database := initDB()
	blogsMongoDB := initMongoDB()
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

	listener, err := net.Listen("tcp", "blogs:8082")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	blogs.RegisterBlogsServiceServer(grpcServer, blogHandler)
	comments.RegisterCommentsServiceServer(grpcServer, commentHandler)
	votes.RegisterVotesServiceServer(grpcServer, voteHandler)
	blog_recommendations.RegisterBlogRecommendationServiceServer(grpcServer, blogRecommendationHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("server error: ", err)
	}
}
