package main

import (
	"encoding/json"
	"log"
	"net/http"
	"trackr-service/internal/initialize"
	"trackr-service/internal/services"
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

	mux.HandleFunc("GET /api/hello", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"message": "Hello from Trackr-Service!",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	mux.HandleFunc("POST /api/user/login", services.UserLogin)
	mux.HandleFunc("POST /api/user/register", services.UserRegister)

	mux.HandleFunc("GET /api/trackr", utils.Authenticate(services.TrackrGetAll))
	mux.HandleFunc("POST /api/trackr", utils.Authenticate(services.TrackrCreate))
	mux.HandleFunc("GET /api/trackr/{id}", utils.Authenticate(services.TrackrGetById))
	mux.HandleFunc("PATCH /api/trackr/{id}/episode", utils.Authenticate(services.TrackrAddCurrentEpisode))
	mux.HandleFunc("PATCH /api/trackr/{id}", utils.Authenticate(services.TrackrUpdate))
	mux.HandleFunc("DELETE /api/trackr/{id}", utils.Authenticate(services.TrackrDelete))

	server.ListenAndServe()
}
