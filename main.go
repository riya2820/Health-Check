package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// global variable, track duration
var startTime = time.Now()

type ComponentStatus struct {
	Status          string     `json:"status"`
	LatestHeartbeat *time.Time `json:"latest heartbeat"`
}

type HealthCheck struct {
	DAGProcessor ComponentStatus `json:"dag_processor"`
	Metadatabase ComponentStatus `json:"metadatabase"`
	Scheduler    ComponentStatus `json:"scheduler"`
	Triggerer    ComponentStatus `json:"triggered"`
	StartTime    ComponentStatus `json:"starttime"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		uptime := time.Since(startTime).String()
		response := map[string]string{
			"uptime": uptime,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// func handler
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		health := HealthCheck{
			DAGProcessor: ComponentStatus{
				LatestHeartbeat: nil,
				Status:          "unavailable",
			},
			Metadatabase: ComponentStatus{
				Status: "healthy",
			},
			Scheduler: ComponentStatus{
				LatestHeartbeat: &now,
				Status:          "healthy",
			},
			Triggerer: ComponentStatus{
				LatestHeartbeat: &now,
				Status:          "healthy",
			},
			StartTime: ComponentStatus{
				LatestHeartbeat: &now,
				Status:          "started",
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(health)
	})

	// Liveness probe
	http.HandleFunc("/live", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Alive!\n"))
	})

	// Readiness probe
	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ready to serve!\n"))
	})

	log.Println("Server running at :8080")
	http.ListenAndServe(":8080", nil)

	server := &http.Server{Addr: ":8080"}

	go func() {
		log.Println("Server running on :8080")
		if err := server.ListenAndServe(); err != nil {
			log.Println("Server stopped:", err)
		}
	}()

	// Wait for keyboard interrupt (Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down server...")
	server.Close()
}
