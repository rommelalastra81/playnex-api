package routes

import (
	"playnex-api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	auth := router.Group("/api")
	{
		// Public routes
		auth.POST("/login", controllers.Login)
	}
}

func UserRoutes(router *gin.Engine) {
	user := router.Group("/api/User")
	{
		// Public routes
		user.POST("/registeruser", controllers.RegisterUser)
	}
}
