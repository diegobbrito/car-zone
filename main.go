package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/diegobbrito/car-zone/driver"
	carHandler "github.com/diegobbrito/car-zone/handler/car"
	engineHandler "github.com/diegobbrito/car-zone/handler/engine"
	"github.com/diegobbrito/car-zone/middleware"
	carService "github.com/diegobbrito/car-zone/service/car"
	engineService "github.com/diegobbrito/car-zone/service/engine"
	carStore "github.com/diegobbrito/car-zone/store/car"
	engineStore "github.com/diegobbrito/car-zone/store/engine"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

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

	router.Use(middleware.AuthMiddleware)

	schemaFile := "store/schema.sql"
	if err := executeSchema(db, schemaFile); err != nil {
		log.Fatalf("Error executing schema: %v", err)
	}

	router.HandleFunc("/cars/{id}", carHandler.GetCarById).Methods("GET")
	router.HandleFunc("/cars", carHandler.GetCarByBrand).Methods("GET")
	router.HandleFunc("/cars", carHandler.CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	router.HandleFunc("/engines/{id}", engineHandler.GetEngineByID).Methods("GET")
	router.HandleFunc("/engines", engineHandler.CreateEngine).Methods("POST")
	router.HandleFunc("/engines/{id}", engineHandler.UpdateEngine).Methods("PUT")
	router.HandleFunc("/engines/{id}", engineHandler.DeleteEngine).Methods("DELETE")

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
