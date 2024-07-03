package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"trackr-service/internal/initialize"
	"trackr-service/internal/models"
)

// Get all trackrs
func TrackrGetAll(w http.ResponseWriter, r *http.Request) {
	var trackrs []models.Trackr

	if err := initialize.DB.Order("id asc").Limit(10).Find(&trackrs).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(trackrs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// Create a new trackr
func TrackrCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	totalEpisode, err1 := strconv.Atoi(r.PostForm.Get("totalEpisode"))
	currentEpisode, err2 := strconv.Atoi(r.PostForm.Get("currentEpisode"))

	if title == "" || err1 != nil || err2 != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	trackr := models.Trackr{Title: title, TotalEpisode: totalEpisode, CurrentEpisode: currentEpisode}

	result := initialize.DB.Create(&trackr)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Trackr created successfully",
		"data":    trackr,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
