package main

import (
	"log"
	"os"

	"github.com/codescalersinternships/secret-note-api-spa-amr/db"
	"github.com/codescalersinternships/secret-note-api-spa-amr/server"
	"github.com/joho/godotenv"
)

// @title Secret Note API
// @version 1.0
// @description This is a simple API for sharing secret notes.
// @host localhost:8000
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	DB := db.InitDatabase()
	r := server.InitServer(DB, os.Getenv("PORT"))

	if err := r.Run(os.Getenv("PORT")); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
