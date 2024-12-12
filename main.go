package main

import (
	"os"
	"log"

	routes "github.com/lucapierini/UserAuthentication/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// controllers "github.com/lucapierini/UserAuthentication/controllers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
		// port = "9000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRouters(router)
	routes.UserRouters(router)
	// router.POST("/signup", controllers.Signup())
	// router.POST("/login", controllers.Login())

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "Access granted for API 1",
		})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "Access granted for API 2",
		})
	})

	router.Run(":" + port)

}
