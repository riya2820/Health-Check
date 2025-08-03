package main

import "net/http"

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handleUptime)
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/live", handleLive)
	mux.HandleFunc("/ready", handleReady)
}
