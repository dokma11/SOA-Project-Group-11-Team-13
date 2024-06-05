package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"tours/handler"
	"tours/messaging/nats"
	"tours/model"
	"tours/proto/equipment"
	"tours/proto/facilities"
	"tours/proto/keypoints"
	"tours/proto/reviews"
	"tours/proto/tours"
	"tours/repo"
	"tours/service"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"

	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dbHost := getEnv("DB_HOST", "localhost");
	dbPort := getEnv("DB_PORT", "5433");
	connectionStr := "host=" + dbHost + " user=postgres password=super dbname=soa-gorm port=" + dbPort + " sslmode=disable"
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

	err = database.AutoMigrate(&model.Facility{})
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("facility auto migrations failed"))
		return nil
	}

	return database
}

var tp *trace.TracerProvider

func initTracer() (*trace.TracerProvider, error) {
	url := "http://jaeger:14268/api/traces"
	if len(url) > 0 {
		return initJaegerTracer(url)
	} else {
		return initFileTracer()
	}
}

func initFileTracer() (*trace.TracerProvider, error) {
	log.Println("Initializing tracing to traces.json")
	f, err := os.Create("traces.json")
	if err != nil {
		return nil, err
	}
	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(f),
		stdouttrace.WithPrettyPrint(),
	)
	if err != nil {
		return nil, err
	}

	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.AlwaysSample()),
	), nil
}

func initJaegerTracer(url string) (*trace.TracerProvider, error) {
	log.Printf("Initializing tracing to jaeger at %s\n", url)
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("tours"),
		)),
	), nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		getEnv("NATS_HOST", "localhost"), getEnv("NATS_PORT", "4222"), subject);
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	var err error
	tp, err = initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	tourRepository := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepository: tourRepository}
	commandPublisher := initPublisher("com.tours");
	tourHandler := &handler.TourHandler{TourService: tourService, CommandPublisher: commandPublisher};

	keyPointRepository := &repo.KeyPointRepository{DatabaseConnection: database}
	keyPointService := &service.KeyPointService{KeyPointRepository: keyPointRepository}
	keyPointHandler := &handler.KeyPointHandler{KeyPointService: keyPointService}

	reviewRepository := &repo.ReviewRepository{DatabaseConnection: database}
	reviewService := &service.ReviewService{ReviewRepository: reviewRepository}
	reviewHandler := &handler.ReviewHandler{ReviewService: reviewService}

	equipmentRepository := &repo.EquipmentRepository{DatabaseConnection: database}
	equipmentService := &service.EquipmentService{EquipmentRepository: equipmentRepository}
	equipmentHandler := &handler.EquipmentHandler{EquipmentService: equipmentService}

	facilityRepository := &repo.FacilityRepository{DatabaseConnection: database}
	facilityService := &service.FacilityService{FacilityRepository: facilityRepository}
	facilityHandler := &handler.FacilityHandler{FacilityService: facilityService}

	listener, err := net.Listen("tcp", "tours:8081")
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

	tours.RegisterToursServiceServer(grpcServer, tourHandler)
	equipment.RegisterEquipmentServiceServer(grpcServer, equipmentHandler)
	facilities.RegisterFacilitiesServiceServer(grpcServer, facilityHandler)
	keypoints.RegisterKeyPointsServiceServer(grpcServer, keyPointHandler)
	reviews.RegisterReviewsServiceServer(grpcServer, reviewHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("server error: ", err)
	}
}
