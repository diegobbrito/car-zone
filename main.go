package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/diegobbrito/car-zone/driver"
	carHandler "github.com/diegobbrito/car-zone/handler/car"
	engineHandler "github.com/diegobbrito/car-zone/handler/engine"
	loginHandler "github.com/diegobbrito/car-zone/handler/login"
	"github.com/diegobbrito/car-zone/middleware"
	carService "github.com/diegobbrito/car-zone/service/car"
	engineService "github.com/diegobbrito/car-zone/service/engine"
	carStore "github.com/diegobbrito/car-zone/store/car"
	engineStore "github.com/diegobbrito/car-zone/store/engine"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	traceProvider, err := startTracing()
	if err != nil {
		log.Fatalf("Error starting tracing: %v", err)
	}

	defer func() {
		if err := traceProvider.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	otel.SetTracerProvider(traceProvider)

	driver.InitDB()
	defer driver.CloseDB()

	db := driver.GetDB()
	carStore := carStore.New(db)
	carService := carService.NewCarService(carStore)
	engineStore := engineStore.New(db)
	engineService := engineService.NewEngineService(engineStore)

	carHandler := carHandler.NewCarHandler(carService)
	engineHandler := engineHandler.NewEngineHandler(engineService)

	router := mux.NewRouter()

	router.Use(otelmux.Middleware("CarZone"))

	schemaFile := "store/schema.sql"
	if err := executeSchema(db, schemaFile); err != nil {
		log.Fatalf("Error executing schema: %v", err)
	}

	router.HandleFunc("/login", loginHandler.LoginHandler).Methods("POST")

	protected := router.PathPrefix("/").Subrouter()

	protected.Use(middleware.AuthMiddleware)

	protected.HandleFunc("/cars/{id}", carHandler.GetCarById).Methods("GET")
	protected.HandleFunc("/cars", carHandler.GetCarByBrand).Methods("GET")
	protected.HandleFunc("/cars", carHandler.CreateCar).Methods("POST")
	protected.HandleFunc("/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	protected.HandleFunc("/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	protected.HandleFunc("/engines/{id}", engineHandler.GetEngineByID).Methods("GET")
	protected.HandleFunc("/engines", engineHandler.CreateEngine).Methods("POST")
	protected.HandleFunc("/engines/{id}", engineHandler.UpdateEngine).Methods("PUT")
	protected.HandleFunc("/engines/{id}", engineHandler.DeleteEngine).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server is running on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))

}

func executeSchema(db *sql.DB, schemaFile string) error {
	sqlFile, err := os.ReadFile(schemaFile)
	if err != nil {
		return fmt.Errorf("error reading schema file: %w", err)
	}
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return fmt.Errorf("error executing schema: %w", err)
	}
	log.Println("Schema executed successfully")
	return nil
}

func startTracing() (*trace.TracerProvider, error) {
	header := map[string]string{
		"Content-Type": "application/json",
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")),
			otlptracehttp.WithHeaders(header),
			otlptracehttp.WithInsecure(),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP exporter: %w", err)
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(
			exporter,
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithBatchTimeout(trace.DefaultScheduleDelay*time.Millisecond),
		),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("CarZone"),
			),
		),
	)

	return tracerProvider, nil
}
