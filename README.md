# Health Check API – Go

This project provides a simple HTTP server written in Go that exposes a basic health check endpoint. It's designed to simulate component-level health status for systems such as DAG processors, metadata databases, and schedulers.

## Features

- REST API endpoint (`/`) that returns a JSON-formatted health status
- Simulated system component status
- Easily extendable structure for real-time health monitoring

## Project Structure

## Getting Started

Prerequisites
Go installed (https://go.dev/dl/)
Port 8080 available (or change in code)
Run the server
go run main.go
Then visit: http://localhost:8080

