package routes

import (
	"password-manager/src/controllers/accountPassword"
	authentication "password-manager/src/controllers/auth"
	"password-manager/src/middlewares"
	"password-manager/src/providers/database"
	"password-manager/src/utils/env"

	"github.com/gin-gonic/gin"
)

func LoadDatabase() {
	database.Connect()
	// database.Database.AutoMigrate(&models.AccountPassword{})
}

func SetupRoutesAndRun(apC *accountPassword.AccountPasswordController, globalEnv *env.Env) *gin.Engine {
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

	serverUrl := globalEnv.ServerHost + ":" + globalEnv.ServerPort
	router.Run(serverUrl)
	return router
}
