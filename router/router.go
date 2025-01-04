package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Set the router as the default one shipped with Gin
	r := gin.Default()

	// Initialize routes
	InitializeRoutes(r)

	// Start serving the application on deafult port 8080
	r.Run()
}
