package handlers

import (
	"net/http"
	"time"

	"github.com/codescalersinternships/secret-note-api-spa-amr/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// NoteResponse represents the response structure for a note
type NoteResponse struct {
	Data model.Note `json:"data"`
	URL  string     `json:"url"`
}

// ErrorResponse represents the response structure for errors
type ErrorResponse struct {
	Error string `json:"error"`
}

// createNote creates a new note with a unique url
// @Summary Create a new note
// @Description Create a new note with a unique URL
// @Tags notes
// @Accept json
// @Produce json
// @Param note body model.Note true "Note to create"
// @Success 200 {object} NoteResponse "Note created successfully"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Router / [post]
func CreateNote(c *gin.Context, db *gorm.DB, port string) {
	var input model.Note
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
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

	db.Create(&note)

	fullURL := "http://localhost" + port + "/" + note.URL
	c.JSON(http.StatusOK, NoteResponse{Data: note, URL: fullURL})
}

// viewNote views a note content
// @Summary View a note
// @Description View a note content
// @Tags notes
// @Accept json
// @Produce json
// @Param url path string true "Note URL"
// @Success 200 {object} NoteResponse "Note retrieved successfully"
// @Failure 400 {object} ErrorResponse "Record not found"
// @Failure 410 {object} ErrorResponse "Note is expired"
// @Router /{url} [get]
func ViewNote(c *gin.Context, db *gorm.DB) {
	var note model.Note

	if err := db.Where("url = ?", c.Param("url")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "record not found!"})
		return
	}

	if note.NoteIsExpired() {
		db.Delete(&note)
		c.JSON(http.StatusGone, ErrorResponse{Error: "note is expired"})
		return
	}

	note.Views++
	db.Save(&note)

	c.JSON(http.StatusOK, NoteResponse{Data: note})
}
