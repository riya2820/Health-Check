package main

import "time"

type ComponentStatus struct {
	Status          string     `json:"status"`
	LatestHeartbeat *time.Time `json:"latest_heartbeat,omitempty"`
}

type HealthCheck struct {
	DAGProcessor ComponentStatus `json:"dag_processor"`
	Metadatabase ComponentStatus `json:"metadatabase"`
	Scheduler    ComponentStatus `json:"scheduler"`
	Triggerer    ComponentStatus `json:"triggerer"`
	StartTime    ComponentStatus `json:"starttime"`
}
