package services

import (
	"net/http"
	"trackr-service/internal/initialize"
	"trackr-service/internal/models"
	"trackr-service/internal/types"
	"trackr-service/internal/utils"
)

// @Summary      Register User
// @Description  Register a new user by providing username, email, and password.
// @Tags         User
// @Accept       multipart/form-data
// @Produce      json
// @Param        username  formData  string  true  "Username"
// @Param        email     formData  string  true  "Email"
// @Param        password  formData  string  true  "Password"
// @Success      200       {object}  map[string]interface{}  "Success response with user details"
// @Router       /user/register [post]
func UserRegister(w http.ResponseWriter, r *http.Request) {
	data, err := utils.BodyParseFormData(r)
	if err != nil {
		utils.CreateErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(data["password"])
	if err != nil {
		utils.CreateErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser := models.User{
		Username: data["username"],
		Email:    data["email"],
		Password: hashedPassword,
	}

	result := initialize.DB.Create(&newUser)
	if result.Error != nil {
		utils.CreateErrorResponse(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	utils.CreateResponse(w, "User signed in successfully", types.JSON{
		"user": models.UserResponse{
			Username: newUser.Username,
			Email:    newUser.Email,
		},
	}, http.StatusOK)
}

// @Summary      Login User
// @Description  Login as a user by providing username (or email) and password.
// @Tags         User
// @Accept       multipart/form-data
// @Produce      json
// @Param        username  formData  string  true  "Username or Email"
// @Param        password  formData  string  true  "Password"
// @Success      200       {object}  map[string]interface{}  "Success response with user details and token"
// @Router       /user/login [post]
func UserLogin(w http.ResponseWriter, r *http.Request) {
	data, err := utils.BodyParseFormData(r)
	if err != nil {
		utils.CreateErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := data["username"]
	password := data["password"]

	var user models.User

	if err := initialize.DB.Where("username = ? OR email = ?", username, username).First(&user).Error; err != nil {
		utils.CreateErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	isPasswordValid, err := utils.VerifyPassword(password, user.Password)
	if err != nil {
		utils.CreateErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !isPasswordValid {
		utils.CreateErrorResponse(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	tokenString, err := utils.GenerateJWT(username, user.ID)
	if err != nil {
		utils.CreateErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.CreateResponse(w, "User logged in successfully", types.JSON{
		"user": models.UserResponse{
			Username: user.Username,
			Email:    user.Email,
		},
		"token": tokenString,
	}, http.StatusOK)
}
