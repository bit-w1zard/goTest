package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/events", CreateEventHandler).Methods("POST")
    r.HandleFunc("/events/{id}", GetEventHandler).Methods("GET")
    r.HandleFunc("/events/{id}", DeleteEventHandler).Methods("DELETE")
    r.HandleFunc("/events", ListEventsHandler).Methods("GET")

    http.Handle("/", AuthMiddleware(r))

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
