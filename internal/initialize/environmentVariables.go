package initialize

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Loaded .env file")
}
