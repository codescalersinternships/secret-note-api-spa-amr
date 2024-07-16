package db

import (
	"log"

	"github.com/codescalersinternships/secret-note-api-spa-amr/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDatabase initialize a new sqlite database
func InitDatabase() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("note.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB.AutoMigrate(&model.Note{}, &model.User{})
	return DB
}
