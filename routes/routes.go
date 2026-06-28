package routes

import (
	"playnex-api/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes sets up all authentication-related routes.
func AuthRoutes(router *gin.Engine) {

	auth := router.Group("/api")
	{
		// Public routes
		auth.POST("/login", controllers.Login)
	}
}
