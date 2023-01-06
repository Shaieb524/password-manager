package routes

import (
	"password-manager/src/controllers/accountPassword"
	"password-manager/src/controllers/authentication"
	"password-manager/src/middlewares"
	"password-manager/src/utils/env"

	"github.com/gin-gonic/gin"
)

func SetupRoutesAndRun(apC *accountPassword.AccountPasswordController, authC *authentication.AuthenticationController, globalEnv *env.Env) *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I'm alive bitch",
		})
	})

	publicRoutes := router.Group("/auth")
	authC.RegisterAuthenticationRoutes(publicRoutes)

	apiV1 := router.Group("/api/v1")
	apiV1.Use(middlewares.JWTAuthentication())
	apC.RegisterAccountPasswordRoutes(apiV1)

	serverUrl := globalEnv.ServerHost + ":" + globalEnv.ServerPort
	router.Run(serverUrl)
	return router
}
