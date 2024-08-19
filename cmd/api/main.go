package main

import (
	"log"
	"net/http"
	"os"

	"trackr-service/docs"
	"trackr-service/internal/initialize"
	"trackr-service/internal/services"
	"trackr-service/internal/utils"

	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {
	initialize.EnvironmentVariables()
	initialize.ConnectDatabase()
}

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: utils.CorsMiddleware(mux),
	}

	host := os.Getenv("TRACKR_HOST")
	if host == "" {
		host = "localhost:8080"
	}

	docs.SwaggerInfo.Title = "Trackr"
	docs.SwaggerInfo.Description = "API for Trackr app. Feel free to try it out 🚀"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.BasePath = "/api"

	log.Println("Server listening on port 8080")

	mux.HandleFunc("/api/status", services.Status)

	mux.HandleFunc("POST /api/user/login", services.UserLogin)
	mux.HandleFunc("POST /api/user/register", services.UserRegister)

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	server.ListenAndServe()
}
