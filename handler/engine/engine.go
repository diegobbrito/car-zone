package engine

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/diegobbrito/car-zone/models"
	"github.com/diegobbrito/car-zone/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type EngineHandler struct {
	service service.EngineServiceInterface
}

func NewEngineHandler(service service.EngineServiceInterface) *EngineHandler {
	return &EngineHandler{
		service: service,
	}
}

func (e *EngineHandler) GetEngineByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	resp, err := e.service.GetEngineByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error getting engine by ID: %v", err)
		return
	}
	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshalling response: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}

func (e *EngineHandler) CreateEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error reading request body: %v", err)
		return
	}
	var engineRequest models.EngineRequest
	err = json.Unmarshal(body, &engineRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error unmarshalling request body: %v", err)
		return
	}
	createdEngine, err := e.service.CreateEngine(ctx, &engineRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error creating engine: %v", err)
		return
	}
	responseBody, err := json.Marshal(createdEngine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshalling response: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(responseBody)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}

func (e *EngineHandler) UpdateEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error reading request body: %v", err)
		return
	}
	var engineRequest models.EngineRequest
	err = json.Unmarshal(body, &engineRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error unmarshalling request body: %v", err)
		return
	}

	updatedEngine, err := e.service.UpdateEngine(ctx, id, &engineRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error updating engine: %v", err)
		return
	}
	responseBody, err := json.Marshal(updatedEngine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshalling response: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func (e *EngineHandler) DeleteEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]
	deletedEngine, err := e.service.DeleteEngine(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error deleting engine: %v", err)
		response := map[string]string{"error": "Failed to delete engine"}
		responseBody, _ := json.Marshal(response)
		w.Write(responseBody)
		return
	}

	if deletedEngine.EngineID == uuid.Nil {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{"error": "Engine not found"}
		responseBody, _ := json.Marshal(response)
		w.Write(responseBody)
		return
	}

	responseBody, err := json.Marshal(deletedEngine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshalling response: %v", err)
		response := map[string]string{"error": "Internal server error"}
		responseBody, _ := json.Marshal(response)
		w.Write(responseBody)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
