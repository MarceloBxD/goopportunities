package handler

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func ListOppeningsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET oppenings",
	})
}

func GetOppeningByIdHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET oppening",
	})
}

func CreateOppeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST oppening",
	})
}

func UpdateOppeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT oppening",
	})
}

func DeleteOppeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE oppening",
	})
}
