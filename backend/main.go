package main

import (
	"log"
	"net/http"
	"time"

	"github.com/codescalersinternships/secret-note-api-spa-amr/db"
	_ "github.com/codescalersinternships/secret-note-api-spa-amr/docs"
	"github.com/codescalersinternships/secret-note-api-spa-amr/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @title Secret Note API
// @version 1.0
// @description This is a simple API for sharing secret notes.
// @host localhost:8000
// @BasePath /
func main() {
	db.InitDatabase()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8000", "*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/", createNote)
	r.GET("/:url", viewNote)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
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
func createNote(c *gin.Context) {
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

	db.DB.Create(&note)

	fullURL := "http://localhost:8000/" + note.URL
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
func viewNote(c *gin.Context) {
	var note model.Note

	if err := db.DB.Where("url = ?", c.Param("url")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "record not found!"})
		return
	}

	if note.NoteIsExpired() {
		db.DB.Delete(&note)
		c.JSON(http.StatusGone, ErrorResponse{Error: "note is expired"})
		return
	}

	note.Views++
	db.DB.Save(&note)

	c.JSON(http.StatusOK, NoteResponse{Data: note})
}
