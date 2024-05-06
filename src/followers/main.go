package main

import (
	"context"
	"followers/handler"
	"followers/proto/followers"
	"followers/repo"
	"followers/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	//port := os.Getenv("PORT")
	//if len(port) == 0 {
	//	port = "8084"
	//}

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

	listener, err := net.Listen("tcp", "followers:8084")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	userService := service.NewUserService(logger, store)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	logger.Println("Server starting with rpc")

	userHandler := handler.NewUserHandler(logger, userService)
	followers.RegisterFollowersServiceServer(grpcServer, userHandler)

	//cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}), gorillaHandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}))
	//server := http.Server{
	//	Addr: ":" + port,
	//	//Handler:      cors(router),
	//	IdleTimeout:  120 * time.Second,
	//	ReadTimeout:  5 * time.Second,
	//	WriteTimeout: 5 * time.Second,
	//}

	//logger.Println("Server starting", port)
	//logger.Println("Server listening on port", port)
	go func() {
		//err := server.ListenAndServe()
		//if err != nil {
		//	logger.Fatal(err)
		//}

		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	//if server.Shutdown(timeoutContext) != nil {
	//	logger.Fatal("Cannot gracefully shutdown...")
	//}
	//logger.Println("Server stopped")
	grpcServer.Stop()

}
