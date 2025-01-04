package router

import (
	"github.com/MarceloBxD/goopportunities/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// API V1 group
	v1 := router.Group("/api/v1")

	// Ping test
	v1.GET("/oppening", handler.GetOppeningByIdHandler)

	v1.GET("/oppenings", handler.ListOppeningsHandler)

	v1.POST("/oppening", handler.CreateOppeningHandler)

	v1.PUT("/oppening", handler.UpdateOppeningHandler)

	v1.DELETE("/oppening", handler.DeleteOppeningHandler)
}
