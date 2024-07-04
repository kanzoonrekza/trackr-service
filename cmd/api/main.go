package main

import (
	"log"
	"net/http"
	"trackr-service/internal/controllers"
	"trackr-service/internal/initialize"
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

	mux.HandleFunc("GET /api/trackr", controllers.TrackrGetAll)
	mux.HandleFunc("POST /api/trackr", controllers.TrackrCreate)

	server.ListenAndServe()
}
