package main

import (
	"log"
	"net/http"
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

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from the server"))
	})

	server.ListenAndServe()
}
