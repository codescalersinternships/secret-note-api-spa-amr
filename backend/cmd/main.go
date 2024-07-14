package main

import (
	"log"
	"net/http"
	"time"

	"github.com/codescalersinternships/secret-note-api-spa-amr/db"
	"github.com/codescalersinternships/secret-note-api-spa-amr/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	db.InitDatabase()
	r := gin.Default()

	r.POST("/", createNote)
	r.GET("/:url", viewNote)

	if err := r.Run(); err != nil {
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
		ID:          input.ID,
		URL:         uuid.NewString(),
		Content:     input.Content,
		PublishDate: time.Now(),
		ExpireAfter: input.ExpireAfter,
		Views:       input.Views,
		MaxViews:    input.MaxViews,
	}

	db.DB.Create(&note)

	c.JSON(http.StatusOK, gin.H{"data": note})

}

func viewNote(c *gin.Context) {
	var note model.Note

	if err := db.DB.Where("url = ?", c.Param("url")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if note.NoteIsExpired() {
		db.DB.Delete(note)
	}

	note.Views++

	c.JSON(http.StatusOK, gin.H{"data": note})
}
