package authentication

import (
	"net/http"
	"password-manager/src/models/database/user"
	services "password-manager/src/services/authentcation"
	"password-manager/src/utils/helper"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthenticationController struct {
	logger   *zap.Logger
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
		authC.logger.Error("Couldn't parse auth input " + c.Request.Header.Get("TransactionId"))
		c.JSON(http.StatusOK, gin.H{"error": err.Error})
		return
	}

	user := user.User{
		Username: input.UserName,
		Password: input.Password,
	}

	savedUser, err := authC.services.RegisterUser(&user)
	if err != nil {
		authC.logger.Error("Couldn't register user" + c.Request.Header.Get("TransactionId"))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authC.logger.Info("User " + savedUser.Username + " was registered successfully!! TransactionId : " + c.Request.Header.Get("TransactionId"))
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
		authC.logger.Error("Couldn't parse auth input" + c.Request.Header.Get("TransactionId"))
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	user, err := authC.services.FindUserByName(input.UserName)
	if err != nil {
		authC.logger.Error("Couldn't find the user" + c.Request.Header.Get("TransactionId"))
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	err = authC.services.ValidatePassword(user, input.Password)
	if err != nil {
		authC.logger.Error("Couldn't find the user" + c.Request.Header.Get("TransactionId"))
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(*user)
	if err != nil {
		authC.logger.Error("Couldn't generate jwt" + c.Request.Header.Get("TransactionId"))
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	authC.logger.Info("User " + input.UserName + " logged in successfully!" + c.Request.Header.Get("TransactionId"))
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
func NewAuthenticationControllerModule(logger *zap.Logger, service *services.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		logger:   logger,
		services: service,
	}
}
