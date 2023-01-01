package main

import (
	"fmt"
	"log"

	"password-manager/src/config/database"
	"password-manager/src/controllers"
	"password-manager/src/middlewares"
	"password-manager/src/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.AccountPassword{})
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)

	proctectedRoutes := router.Group("/api/v1")
	proctectedRoutes.Use(middlewares.JWTAuthentication())
	proctectedRoutes.GET("/")

	router.Run("localhost:1111")
}

func main() {
	fmt.Println("Start bitch")
	loadEnv()
	loadDatabase()
	serveApplication()
}
