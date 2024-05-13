package main

import (
	"jwt/handler"
	jwtPb "jwt/proto/jwt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}

	logger := log.New(os.Stdout, "[jwt-api] ", log.LstdFlags)

	// Initialize any dependencies your JWT service needs
	// For example, if you need a connection to a data store:
	// store, err := repo.NewStore(logger)
	// if err != nil {
	//     logger.Fatal("Failed to create store:", err)
	// }
	// defer store.Close() // Make sure to close any resources when the application exits

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	// Assuming JwtHandler is correctly implemented in the handler package
	jwtHandler := handler.NewJwtHandler() // Replace with actual initialization if needed

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// Register your service with the gRPC server
	jwtPb.RegisterJwtServiceServer(grpcServer, jwtHandler)

	go func() {
		logger.Printf("Server starting at port %s", port)
		if err := grpcServer.Serve(listener); err != nil {
			logger.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Graceful Shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigCh
	logger.Printf("Received terminate, graceful shutdown %s", sig)
	grpcServer.GracefulStop()
}
