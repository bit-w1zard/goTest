package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "time"
    "github.com/google/uuid"
    "sort"
)

var events = make(map[string]Event)

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.Header.Get("UserID")
    var event Event
    if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if event.Type != "shipping" && event.Type != "receiving" {
        http.Error(w, "Invalid event type", http.StatusBadRequest)
        return
    }

    event.ID = uuid.New().String()
    event.CreatedAt = time.Now().UTC()
    event.CreatedBy = userID
    event.IsDeleted = false

    events[event.ID] = event

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(event)
}

func GetEventHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.Header.Get("UserID")
    vars := mux.Vars(r)
    id := vars["id"]

    event, exists := events[id]
    if !exists || event.CreatedBy != userID || event.IsDeleted {
        http.Error(w, "Event not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(event)
}

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.Header.Get("UserID")
    vars := mux.Vars(r)
    id := vars["id"]

    event, exists := events[id]
    if !exists || event.CreatedBy != userID {
        http.Error(w, "Event not found", http.StatusNotFound)
        return
    }

    event.IsDeleted = true
    events[id] = event

    w.WriteHeader(http.StatusOK)
}

func ListEventsHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.Header.Get("UserID")
    userEvents := []Event{}

    for _, event := range events {
        if event.CreatedBy == userID && !event.IsDeleted {
            userEvents = append(userEvents, event)
        }
    }


    sort.Slice(userEvents, func(i, j int) bool {
        return userEvents[i].CreatedAt.After(userEvents[j].CreatedAt)
    })

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(userEvents)
}
