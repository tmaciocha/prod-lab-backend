package main

// It creates a simple HTTP server that listens on port 8080 and responds to requests at the /health endpoint with a JSON object indicating the status of the server. 
// The response is in the format {"status": "ok"}. 
// The server uses the net/http package to handle HTTP requests and responses, and the encoding/json package to encode the response as JSON.

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HealthResponse{Status: "ok"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Backend listening on :8080")
	log.Fatal(server.ListenAndServe())
}
