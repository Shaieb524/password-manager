package routes

import (
	"password-manager/docs"
	"password-manager/src/controllers/accountPassword"
	"password-manager/src/controllers/authentication"
	"password-manager/src/middlewares"
	"password-manager/src/utils/env"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title Go + Gin Todo API
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:1111
// @BasePath /
// @query.collection.format multi
func SetupRoutesAndRun(apC *accountPassword.AccountPasswordController, authC *authentication.AuthenticationController, globalEnv *env.Env) *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api/v1"

	router := gin.Default()
	router.Use(middlewares.TransactionIdGenerator())
	router.Use(middlewares.CORSMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:1111/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	publicRoutes := router.Group("/api/v1/auth")
	authC.RegisterAuthenticationRoutes(publicRoutes)

	privateRoutes := router.Group("/api/v1")
	privateRoutes.Use(middlewares.JWTAuthentication())
	apC.RegisterAccountPasswordRoutes(privateRoutes)

	// Health Check example go doc
	// @BasePath /
	// @Summary Health check endpoint
	// @Description do Health Check
	// @Tags example
	// @Accept json
	// @Produce json
	// @Success 200 {string} I'm alive bitch
	// @Router /api/v1/health [get]
	router.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I'm alive bitch",
		})
	})

	serverUrl := globalEnv.ServerHost + ":" + globalEnv.ServerPort
	router.Run(serverUrl)
	return router
}
