package services

import (
	"net/http"
	"strconv"
	"trackr-service/internal/initialize"
	"trackr-service/internal/models"
	"trackr-service/internal/utils"
)

func TrackrGetAll(w http.ResponseWriter, r *http.Request) {
	claims, ok := utils.GetClaims(r.Context())
	if !ok {
		utils.CreateErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var trackrs []models.Trackr

	if err := initialize.DB.Order("id asc").Limit(10).Where(&models.Trackr{UserID: claims.UserID}).Find(&trackrs).Error; err != nil {
		utils.CreateErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.CreateResponse(w, "Trackrs retrieved successfully", utils.PayloadType{
		"data": trackrs,
	}, http.StatusOK)
}

func TrackrCreate(w http.ResponseWriter, r *http.Request) {
	claims, ok := utils.GetClaims(r.Context())
	if !ok {
		utils.CreateErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		utils.CreateErrorResponse(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	totalEpisode, err1 := strconv.Atoi(r.PostForm.Get("totalEpisode"))
	currentEpisode, err2 := strconv.Atoi(r.PostForm.Get("currentEpisode"))

	if title == "" || err1 != nil || err2 != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	trackr := models.Trackr{Title: title, TotalEpisode: uint16(totalEpisode), CurrentEpisode: uint16(currentEpisode), UserID: uint(claims.UserID)}

	result := initialize.DB.Create(&trackr)
	if result.Error != nil {
		utils.CreateErrorResponse(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	utils.CreateResponse(w, "Trackr created successfully", utils.PayloadType{
		"data": trackr,
	}, http.StatusOK)
}

func TrackrGetById(w http.ResponseWriter, r *http.Request) {
	claims, ok := utils.GetClaims(r.Context())
	if !ok {
		utils.CreateErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	trackrid, _ := strconv.Atoi(r.PathValue("id"))

	var trackr models.Trackr
	if err := initialize.DB.Where(models.Trackr{UserID: claims.UserID}).First(&trackr, trackrid).Error; err != nil {
		utils.CreateErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.CreateResponse(w, "Trackr retrieved successfully", utils.PayloadType{
		"data": trackr,
	}, http.StatusOK)
}

func TrackrAddCurrentEpisode(w http.ResponseWriter, r *http.Request) {
	claims, ok := utils.GetClaims(r.Context())
	if !ok {
		utils.CreateErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	trackrid, _ := strconv.Atoi(r.PathValue("id"))

	var trackr models.Trackr
	if err := initialize.DB.Where(models.Trackr{UserID: claims.UserID}).First(&trackr, trackrid).Error; err != nil {
		utils.CreateErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	eps, _ := strconv.Atoi(r.URL.Query().Get("eps"))

	if (trackr.CurrentEpisode + uint16(eps)) >= trackr.TotalEpisode {
		trackr.CurrentEpisode = trackr.TotalEpisode
		trackr.Completed = true
	} else {
		trackr.CurrentEpisode = trackr.CurrentEpisode + uint16(eps)
	}

	if trackr.CurrentEpisode == trackr.TotalEpisode {
		trackr.Completed = true
	} else {
		trackr.Completed = false
	}

	initialize.DB.Save(&trackr)

	utils.CreateResponse(w, "Trackr updated successfully", utils.PayloadType{
		"data": trackr,
	}, http.StatusOK)
}

func TrackrUpdate(w http.ResponseWriter, r *http.Request) {
	claims, ok := utils.GetClaims(r.Context())
	if !ok {
		utils.CreateErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	trackrid, _ := strconv.Atoi(r.PathValue("id"))

	var trackr models.Trackr
	if err := initialize.DB.Where(models.Trackr{UserID: claims.UserID}).First(&trackr, trackrid).Error; err != nil {
		utils.CreateErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		utils.CreateErrorResponse(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	if title != "" {
		trackr.Title = title
	}

	totalEpisode, err := strconv.Atoi(r.PostForm.Get("totalEpisode"))
	if err == nil {
		if totalEpisode < 0 || totalEpisode > 65535 {
			utils.CreateErrorResponse(w, "Total episode should be between 0 and 65535", http.StatusBadRequest)
			return
		}
		trackr.TotalEpisode = uint16(totalEpisode)
	}

	currentEpisode, err := strconv.Atoi(r.PostForm.Get("currentEpisode"))
	if err == nil {
		if currentEpisode < 0 || currentEpisode > 65535 || uint16(currentEpisode) > trackr.TotalEpisode {
			utils.CreateErrorResponse(w, "Current episode should be between 0 and 65535 and not exceeding total episode", http.StatusBadRequest)
			return
		}
		trackr.CurrentEpisode = uint16(currentEpisode)
	}

	if trackr.CurrentEpisode == trackr.TotalEpisode {
		trackr.Completed = true
	} else {
		trackr.Completed = false
	}

	rate, err := strconv.Atoi(r.PostForm.Get("rate"))
	if err == nil {
		if rate < 0 || rate > 10 {
			utils.CreateErrorResponse(w, "Rate should be between 0 and 10", http.StatusBadRequest)
			return
		}
		trackr.Rate = int8(rate)
	}

	initialize.DB.Save(&trackr)

	utils.CreateResponse(w, "Trackr updated successfully", utils.PayloadType{
		"data": trackr,
	}, http.StatusOK)
}

func TrackrDelete(w http.ResponseWriter, r *http.Request) {
	claims, ok := utils.GetClaims(r.Context())
	if !ok {
		utils.CreateErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	trackrid, _ := strconv.Atoi(r.PathValue("id"))

	var trackr models.Trackr
	if err := initialize.DB.Where(models.Trackr{UserID: claims.UserID}).First(&trackr, trackrid).Error; err != nil {
		utils.CreateErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	initialize.DB.Delete(&trackr)

	utils.CreateResponse(w, "Trackr deleted successfully", utils.PayloadType{
		"data": trackr,
	}, http.StatusOK)
}
