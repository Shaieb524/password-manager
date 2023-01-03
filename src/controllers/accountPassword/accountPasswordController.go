package accountPassword

import (
	"net/http"
	"password-manager/src/models/database/accountPassword"
	services "password-manager/src/services/accountPassword"

	"github.com/gin-gonic/gin"
)

type AccountPasswordController struct {
	service services.AccountPasswordService
}

type accountPasswordController interface {
	CreateAccountPassword(*gin.Context)
	GetAppPasswordById(*gin.Context)
}

func (apC *AccountPasswordController) CreatAccountPassword(ctx *gin.Context) {
	var accPassword *accountPassword.AccountPasswordModel

	err := ctx.BindJSON(&accPassword)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not parse request : " + err.Error(),
		})
	}

	if accPassword, err := apC.service.CreateAccountPassword(*accPassword); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating item : " + err.Error(),
		})
	} else {
		ctx.JSON(http.StatusCreated, accPassword)
	}
}

func (apC *AccountPasswordController) GetAppPasswordById(ctx *gin.Context) {
	appPasswordId := ctx.Param("id")

	if appPassword, err := apC.service.GetAppPasswordById(appPasswordId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, appPassword)
	}
}

func (apC *AccountPasswordController) RegisterRoutes(router *gin.RouterGroup) {
	apC.registerAccoutPasswordsRoutes(router)
}

func (apC *AccountPasswordController) registerAccoutPasswordsRoutes(router *gin.RouterGroup) {
	router.GET("/accountPassword", apC.CreatAccountPassword)
	router.POST("/accountPassword", apC.GetAppPasswordById)
}
