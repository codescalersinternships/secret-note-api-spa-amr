package handlers_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/codescalersinternships/secret-note-api-spa-amr/handlers"
	"github.com/codescalersinternships/secret-note-api-spa-amr/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNote(t *testing.T) {
	db := setupTestDB()
	defer cleanupDB(db)

	router := gin.Default()
	router.POST("/create", func(c *gin.Context) { handlers.CreateNote(c, db, "8000") })

	note := model.Note{
		Content:     "test note",
		ExpireAfter: 24,
		MaxViews:    10,
	}
	payload, _ := json.Marshal(note)

	req, err := http.NewRequest("POST", "/create", strings.NewReader(string(payload)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response handlers.NoteResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, response.URL)
	assert.Equal(t, note.Content, response.Data.Content)
}

func TestViewNote(t *testing.T) {
	db := setupTestDB()
	defer cleanupDB(db)

	note := model.Note{
		URL:         "test-url",
		Content:     "test note",
		PublishDate: time.Now(),
		ExpireAfter: 24,
		Views:       0,
		MaxViews:    10,
	}
	db.Create(&note)

	router := gin.Default()
	router.GET("/:url", func(c *gin.Context) { handlers.ViewNote(c, db) })

	req, err := http.NewRequest("GET", "/"+note.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response handlers.NoteResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, note.Content, response.Data.Content)

	var updatedNote model.Note
	if err := db.Where("url = ?", note.URL).First(&updatedNote).Error; err != nil {
		t.Fatalf("failed to retrieve updated note from database: %v", err)
	}
	assert.Equal(t, 1, updatedNote.Views)
}

func TestDatabaseIntegration(t *testing.T) {
	db := setupTestDB()
	defer cleanupDB(db)

	user := model.User{
		Name:     "testuser",
		Password: "testpassword",
	}
	db.Create(&user)

	note := model.Note{
		URL:         "test-url",
		Content:     "This is a test note",
		PublishDate: time.Now(),
		ExpireAfter: 24,
		MaxViews:    10,
		UserID:      user.ID,
	}
	db.Create(&note)

	var retrievedNote model.Note
	result := db.First(&retrievedNote, "url = ?", note.URL)
	assert.NoError(t, result.Error)
	assert.Equal(t, note.Content, retrievedNote.Content)
}

// setupTestDB initializes a new test database
func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect test database: %v", err)
	}
	err = db.AutoMigrate(&model.Note{}, &model.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	return db
}

// cleanupDB drops the tables after test execution
func cleanupDB(db *gorm.DB) {
	err := db.Migrator().DropTable(&model.Note{}, &model.User{})
	if err != nil {
		log.Fatalf("failed to drop tables: %v", err)
	}
	err = os.Remove("test.db")
	if err != nil {
		log.Fatalf("failed to delete test database file: %v", err)
	}
}
