package main

import (
	config "playnex-api/configs"
	"playnex-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to database
	config.ConnectDB()

	// Auto-migrate the Users table
	//config.DB.AutoMigrate(&models.Users{})

	// Initialize router
	router := gin.Default()

	// CORS configuration - allow localhost:3001
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// routes
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	router.Run() // listens on 0.0.0.0:8080 by default
}
