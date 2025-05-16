package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// service health status
type HealthCheck struct {
	Service string    `json:"service"`
	Status  string    `json:"status"`
	Time    time.Time `json:"time"`
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		status := HealthCheck{
			Service: "basic-service",
			Status:  "OK",
			Time:    time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(status)
	})

	log.Println("Server running at :8080")
	http.ListenAndServe(":8080", nil)
}

