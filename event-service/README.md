This project is a simple REST service built with Go that provides APIs for creating, retrieving, deleting, and listing events. The project includes rudimentary authentication using bearer tokens and can be deployed using Docker.

Prerequisites

    Go 1.16 or later
    Docker
    Docker Compose
    Make

Install dependencies:
go mod tidy

You can run the service using the Makefile:
make run

Run with Docker:
docker build -t event-service .
docker run -p 8080:8080 event-service

docker-compose up

API Endpoints
Authentication

All endpoints require a bearer token for authentication. The following tokens are valid:

    74edf612f393b4eb01fbc2c29dd96671 for user ID 12345
    d88b4b1e77c70ba780b56032db1c259b for user ID 98765

Makefile

The Makefile includes common tasks for building and running the project:

    make run: Run the service
    make build: Build the Go binary
    make docker: Build the Docker image
