package routes

import (
	"log"
	"os"
	"password-manager/src/controllers/accountPassword"
	authentication "password-manager/src/controllers/auth"
	"password-manager/src/middlewares"
	"password-manager/src/providers/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
}

func LoadDatabase() {
	database.Connect()
	// database.Database.AutoMigrate(&models.AccountPassword{})
}

func SetupRoutesAndRun(apC accountPassword.AccountPasswordController) *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I'm alive bitch",
		})
	})

	publicRoutes := router.Group("/auth")
	authentication.RegisterRoutes(publicRoutes)

	apiV1 := router.Group("/api/v1")
	apiV1.Use(middlewares.JWTAuthentication())
	apC.RegisterRoutes(apiV1)

	serverUrl := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")
	router.Run(serverUrl)
	return router
}
