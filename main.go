package main

import (
	"fmt"
	"log"

	"password-manager/src/config/database"
	"password-manager/src/controller"
	"password-manager/src/middleware"
	"password-manager/src/model"

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
	database.Database.AutoMigrate(&model.AccountPassword{})
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	proctectedRoutes := router.Group("/api/v1")
	proctectedRoutes.Use(middleware.JWTAuthentication())
	proctectedRoutes.GET("/")

	router.Run("localhost:1111")
}

func main() {
	fmt.Println("Start bitch")
	loadEnv()
	loadDatabase()
	serveApplication()
}
