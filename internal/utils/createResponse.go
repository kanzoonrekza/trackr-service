package utils

import (
	"encoding/json"
	"net/http"
	"trackr-service/internal/types"
)

func CreateResponse(w http.ResponseWriter, message string, payload types.JSON, statusCode int) {
	response := types.JSON{
		"statusCode": statusCode,
		"message":    message,
		"payload":    payload,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func CreateErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	CreateResponse(w, message, nil, statusCode)
}
