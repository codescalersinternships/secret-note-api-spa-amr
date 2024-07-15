package main

import (
	"log"
	"net/http"
	"time"

	"github.com/codescalersinternships/secret-note-api-spa-amr/db"
	"github.com/codescalersinternships/secret-note-api-spa-amr/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	db.InitDatabase()
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:8080"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/", createNote)
	r.GET("/:url", viewNote)

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func createNote(c *gin.Context) {
	var input model.Note
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note := model.Note{
		URL:         uuid.NewString(),
		Content:     input.Content,
		PublishDate: time.Now(),
		ExpireAfter: input.ExpireAfter,
		Views:       0,
		MaxViews:    input.MaxViews,
	}

	db.DB.Create(&note)

	fullURL := "http://localhost:8000/" + note.URL
	c.JSON(http.StatusOK, gin.H{"data": note, "url": fullURL})
}

func viewNote(c *gin.Context) {
	var note model.Note

	if err := db.DB.Where("url = ?", c.Param("url")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if note.NoteIsExpired() {
		db.DB.Delete(&note)
		c.JSON(http.StatusGone, gin.H{"error": "Note is expired"})
		return
	}

	note.Views++
	db.DB.Save(&note)

	c.JSON(http.StatusOK, gin.H{"data": note})
}
