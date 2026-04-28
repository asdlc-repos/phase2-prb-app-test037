package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type timeResponse struct {
	Timestamp string `json:"timestamp"`
}

type healthResponse struct {
	Status string `json:"status"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	now := time.Now().UTC().Format(time.RFC3339)
	resp := timeResponse{Timestamp: now}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse{Error: "Internal server error"})
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := healthResponse{Status: "ok"}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse{Error: "Internal server error"})
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", timeHandler)
	mux.HandleFunc("/health", healthHandler)

	addr := ":9090"
	log.Printf("Starting time-service on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
