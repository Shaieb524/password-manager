package authentication

import (
	"net/http"
	"password-manager/src/models"
	"password-manager/src/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input models.AuthenticationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error})
		return
	}

	user := models.User{
		Username: input.UserName,
		Password: input.Password,
	}

	savedUser, err := user.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(c *gin.Context) {
	var input models.AuthenticationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.FindUserByUsername(input.UserName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := utils.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

func RegisterRoutes(router *gin.RouterGroup) {
	registerAuthenticationRoutes(router)
}

func registerAuthenticationRoutes(router *gin.RouterGroup) {
	router.POST("/register", Register)
	router.POST("/login", Login)
}
