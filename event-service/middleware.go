package main

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        log.Printf("Authorization Header: %s", authHeader)

        if userID, ok := authenticate(r); ok {
            r.Header.Set("UserID", userID)
            next.ServeHTTP(w, r)
        } else {
            log.Printf("Forbidden: Invalid token")
            http.Error(w, "Forbidden", http.StatusForbidden)
        }
    })
}
