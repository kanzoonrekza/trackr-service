package utils

import (
	"encoding/json"
	"net/http"
)

type PayloadType map[string]interface{}

func CreateResponse(w http.ResponseWriter, message string, payload PayloadType, statusCode int) {
	response := map[string]interface{}{
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
