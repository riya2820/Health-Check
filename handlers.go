package main

import (
	"net/http"
	"time"
)

func handleUptime(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime).String()
	response := map[string]string{"uptime": uptime}
	writeJSON(w, response)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	health := HealthCheck{
		DAGProcessor: ComponentStatus{Status: "unavailable"},
		Metadatabase: ComponentStatus{Status: "healthy"},
		Scheduler:    ComponentStatus{Status: "healthy", LatestHeartbeat: &now},
		Triggerer:    ComponentStatus{Status: "healthy", LatestHeartbeat: &now},
		StartTime:    ComponentStatus{Status: "started", LatestHeartbeat: &now},
	}
	writeJSON(w, health)
}

func handleLive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Alive!\n"))
}

func handleReady(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ready to serve!\n"))
}
