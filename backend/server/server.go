package server

import (
	"time"

	_ "github.com/codescalersinternships/secret-note-api-spa-amr/docs"
	"github.com/codescalersinternships/secret-note-api-spa-amr/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func InitServer(db *gorm.DB, port string) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost" + port, "*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/", func(c *gin.Context) { handlers.CreateNote(c, db, port) })
	r.GET("/:url", func(c *gin.Context) { handlers.ViewNote(c, db) })

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
