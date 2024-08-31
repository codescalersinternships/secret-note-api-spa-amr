package handlers

import (
	"net/http"

	"github.com/codescalersinternships/secret-note-api-spa-amr/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SignUpRequest represents the request structure for signing up
type SignUpRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// SignInRequest represents the request structure for signing in
type SignInRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserResponse represents the response structure for user
type UserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// SignUp creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided name and password
// @Tags users
// @Accept json
// @Produce json
// @Param input body SignUpRequest true "User details"
// @Success 200 {object} UserResponse "User created successfully"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Failed to create user"
// @Router /signup [post]
func SignUp(c *gin.Context, db *gorm.DB) {
	var input SignUpRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	user := model.User{Name: input.Name, Password: input.Password}
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, UserResponse{ID: user.ID, Name: user.Name})
}

// SignIn authenticates a user
// @Summary Authenticate a user
// @Description Authenticate a user with the provided name and password
// @Tags users
// @Accept json
// @Produce json
// @Param input body SignInRequest true "User credentials"
// @Success 200 {object} UserResponse "User authenticated successfully"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 401 {object} ErrorResponse "Invalid credentials"
// @Router /signin [post]
func SignIn(c *gin.Context, db *gorm.DB) {
	var input SignInRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	var user model.User
	if err := db.Where("name = ? AND password = ?", input.Name, input.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, UserResponse{ID: user.ID, Name: user.Name})
}
