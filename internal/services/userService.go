package services

import (
	"net/http"
	"trackr-service/internal/initialize"
	"trackr-service/internal/models"
	"trackr-service/internal/utils"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		utils.CreateErrorResponse(w, "Failed to parse form data", http.StatusBadRequest)
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

	utils.CreateResponse(w, "User signed in successfully", utils.PayloadType{
		"user": userResponse,
	}, http.StatusOK)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		utils.CreateErrorResponse(w, "Failed to parse form data", http.StatusBadRequest)
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

	utils.CreateResponse(w, "User logged in successfully", utils.PayloadType{
		"user": models.UserResponse{
			Username: user.Username,
			Email:    user.Email,
		},
		"token": tokenString,
	}, http.StatusOK)
}
