package main

import (
	"log"
	"net/http"
	"trackr-service/internal/initialize"
	"trackr-service/internal/services"
	"trackr-service/internal/utils"

	_ "trackr-service/docs" // swaggo docs

	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {
	initialize.EnvironmentVariables()
	initialize.ConnectDatabase()
}

// @title Trackr API
// @version 1.0
// @description This is a sample server for a Trackr app.
// @host trackr-service-production.up.railway.app
func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: utils.CorsMiddleware(mux),
	}
	log.Println("Server listening on port 8080")

	mux.HandleFunc("/api/status", services.Status)

	mux.HandleFunc("POST /api/user/login", services.UserLogin)
	mux.HandleFunc("POST /api/user/register", services.UserRegister)

	mux.HandleFunc("GET /api/trackr", utils.Authenticate(services.TrackrGetAll))
	mux.HandleFunc("POST /api/trackr", utils.Authenticate(services.TrackrCreate))
	mux.HandleFunc("GET /api/trackr/{id}", utils.Authenticate(services.TrackrGetById))
	mux.HandleFunc("PATCH /api/trackr/{id}/episode", utils.Authenticate(services.TrackrAddCurrentEpisode))
	mux.HandleFunc("PATCH /api/trackr/{id}", utils.Authenticate(services.TrackrUpdate))
	mux.HandleFunc("DELETE /api/trackr/{id}", utils.Authenticate(services.TrackrDelete))

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	server.ListenAndServe()
}
