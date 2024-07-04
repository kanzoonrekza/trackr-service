package initialize

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvironmentVariables() {
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		// We're not on Railway, so try to load from .env file
		err := godotenv.Load()
		log.Println("Using environment variables from .env file")
		if err != nil {
			log.Println("Warning: Error loading .env file. Using existing environment variables.")
		} else {
			log.Println("Loaded .env file")
		}
	} else {
		log.Println("Running on Railway. Using provided environment variables.")
	}
}
