package services

import (
	"encoding/json"
	"net/http"
)

// @Summary Status
// @Description Check if the service is up and running
// @Tags Status
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/status [get]
func Status(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status": "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
