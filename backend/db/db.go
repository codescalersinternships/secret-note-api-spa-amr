package db

import (
	"log"

	"github.com/codescalersinternships/secret-note-api-spa-amr/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("note.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB.AutoMigrate(&model.Note{})
}
