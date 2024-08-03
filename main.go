package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatalf("Unable to open connection for the database")
	}

	API := NewAPIServer(":3000", store)

	API.Run()

}
