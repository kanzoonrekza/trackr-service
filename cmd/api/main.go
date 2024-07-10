package main

import (
	"log"
	"net/http"
	"trackr-service/internal/controllers"
	"trackr-service/internal/initialize"
	"trackr-service/internal/utils"
)

func init() {
	initialize.EnvironmentVariables()
	initialize.ConnectDatabase()
}

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Server listening on port 8080")

	mux.HandleFunc("POST /api/user/login", controllers.UserLogin)
	mux.HandleFunc("POST /api/user/register", controllers.UserRegister)

	mux.HandleFunc("GET /api/trackr", utils.Authenticate(controllers.TrackrGetAll))
	mux.HandleFunc("POST /api/trackr", utils.Authenticate(controllers.TrackrCreate))
	mux.HandleFunc("PATCH /api/trackr/{id}/episode", utils.Authenticate(controllers.TrackrAddCurrentEpisode))
	mux.HandleFunc("PATCH /api/trackr/{id}", utils.Authenticate(controllers.TrackrUpdate))

	server.ListenAndServe()
}
