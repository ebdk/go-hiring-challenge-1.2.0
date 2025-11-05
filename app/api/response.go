package api

import (
    "encoding/json"
    "net/http"
)

func OKResponse(w http.ResponseWriter, data any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _ = json.NewEncoder(w).Encode(data)
}

func ErrorResponse(w http.ResponseWriter, status int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    _ = json.NewEncoder(w).Encode(struct {
        Error string `json:"error"`
    }{Error: message})
}

// CreatedResponse writes a 201 response with JSON data.
func CreatedResponse(w http.ResponseWriter, data any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    _ = json.NewEncoder(w).Encode(data)
}
