package initialisers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// script to load our environment variables using godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
