package db

import (
	"log"

	"github.com/codescalersinternships/secret-note-api-spa-amr/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB variable of type gorm to control the database
var DB *gorm.DB

// InitDatabase initialize a new sqlite database
func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("note.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB.AutoMigrate(&model.Note{})
}
