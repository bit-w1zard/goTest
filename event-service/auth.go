package main

import (
	"net/http"
	"strings"
)

var validTokens = map[string]string{
    "74edf612f393b4eb01fbc2c29dd96671": "12345",
    "d88b4b1e77c70ba780b56032db1c259b": "98765",
}

func authenticate(r *http.Request) (string, bool) {
    token := r.Header.Get("Authorization")
    if token == "" {
        return "", false
    }
    parts := strings.Split(token, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
        return "", false
    }

    actualToken := parts[1]

    if userID, ok := validTokens[actualToken]; ok {
        return userID, true
    }

    return "", false
}
