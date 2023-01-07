package authentication

import (
	"fmt"
	"net/http"
	"password-manager/src/models/database/user"
	services "password-manager/src/services/authentcation"
	"password-manager/src/utils/helper"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	services *services.AuthenticationService
}

type authenticationController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

// @Tags Authentication
// @Summary register a new user
// @Description register a new user
// @Accept json
// @Produce json
// @Param data body user.AuthenticationInput true "payload"
// @Success 200 {object} user.User
// Failure 400 {object} models.ErrorResponse
// Failure 404 {object} models.ErrorResponse
// Failure 500 {object} models.ErrorResponse
// @Router /auth/register [post]
func (authC *AuthenticationController) Register(c *gin.Context) {
	var input user.AuthenticationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error})
		return
	}

	user := user.User{
		Username: input.UserName,
		Password: input.Password,
	}

	savedUser, err := authC.services.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

// @Tags Authentication
// @Summary login a user
// @Description login a user
// @Accept json
// @Produce json
// @Param data body user.AuthenticationInput true "payload"
// Success 200 {object}
// Failure 400 {object} models.ErrorResponse
// Failure 404 {object} models.ErrorResponse
// Failure 500 {object} models.ErrorResponse
// @Router /auth/login [post]
func (authC *AuthenticationController) Login(c *gin.Context) {
	var input user.AuthenticationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parsing": err.Error()})
		return
	}

	user, err := authC.services.FindUserByName(input.UserName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error finding user": err.Error()})
		return
	}

	fmt.Println(user.Username)
	fmt.Println(user.Password)
	err = authC.services.ValidatePassword(user, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error validating pass": err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(*user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error generating jwt": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

// Routing
func (authC *AuthenticationController) RegisterAuthenticationRoutes(router *gin.RouterGroup) {
	authC.registerAuthenticationRoutes(router)
}

func (authC *AuthenticationController) registerAuthenticationRoutes(router *gin.RouterGroup) {
	router.POST("/register", authC.Register)
	router.POST("/login", authC.Login)
}

// DI
func NewAuthenticationControllerModule(service *services.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{services: service}
}
