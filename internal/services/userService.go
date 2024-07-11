package services

import (
	"encoding/json"
	"net/http"
	"trackr-service/internal/initialize"
	"trackr-service/internal/models"
	"trackr-service/internal/utils"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(r.PostForm.Get("password"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser := models.User{
		Username: r.PostForm.Get("username"),
		Email:    r.PostForm.Get("email"),
		Password: hashedPassword,
	}

	result := initialize.DB.Create(&newUser)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	userResponse := models.UserResponse{
		Username: newUser.Username,
		Email:    newUser.Email,
	}

	response := map[string]interface{}{
		"message": "User created successfully",
		"data":    userResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	var user models.User

	if err := initialize.DB.Where("username = ? OR email = ?", username, username).First(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	isPasswordValid, err := utils.VerifyPassword(password, user.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !isPasswordValid {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	tokenString, err := utils.GenerateJWT(username, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userResponse := models.UserResponse{
		Username: user.Username,
		Email:    user.Email,
	}

	response := map[string]interface{}{
		"message": "User logged in successfully",
		"data":    userResponse,
		"token":   tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
