package accountPassword

import (
	"net/http"
	"password-manager/src/models/database/accountPassword"
	services "password-manager/src/services/accountPassword"

	"github.com/gin-gonic/gin"
)

type AccountPasswordController struct {
	service *services.AccountPasswordService
}

type accountPasswordController interface {
	CreateAccountPassword(*gin.Context)
	GetAllAccountsPasswords(*gin.Context)
	GetAppPasswordById(*gin.Context)
	GetAppPasswordByServiceName(ctx *gin.Context)
	EditServicePassword(ctx *gin.Context)
	DeleteServicePassword(ctx *gin.Context)
	ProvideModuleforDI(service *services.AccountPasswordService) *AccountPasswordController
}

func (apC *AccountPasswordController) CreatAccountPassword(ctx *gin.Context) {
	var accPassword *accountPassword.AccountPassword

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

func (apC *AccountPasswordController) GetAllAccountsPasswords(ctx *gin.Context) {

	if accPasswords, err := apC.service.GetAllAccountsPasswords(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating item : " + err.Error(),
		})
	} else {
		ctx.JSON(http.StatusCreated, accPasswords)
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

func (apC *AccountPasswordController) GetAppPasswordByServiceName(ctx *gin.Context) {
	accountName := ctx.Param("name")

	if appPassword, err := apC.service.GetAppPasswordByServiceName(accountName); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, appPassword)
	}
}

func (apC *AccountPasswordController) EditAccountPassword(ctx *gin.Context) {
	var accPassword *accountPassword.AccountPassword

	err := ctx.BindJSON(&accPassword)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not parse request : " + err.Error(),
		})
	}

	if accPassword, err := apC.service.EditServicePassword(*accPassword); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating item : " + err.Error(),
		})
	} else {
		ctx.JSON(http.StatusCreated, accPassword)
	}
}

func (apC *AccountPasswordController) DeleteServicePassword(ctx *gin.Context) {
	accountName := ctx.Param("name")

	if err := apC.service.DeleteServicePassword(accountName); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusAccepted, "done")
	}
}

// Routing
func (apC *AccountPasswordController) RegisterAccountPasswordRoutes(router *gin.RouterGroup) {
	apC.registerAccoutPasswordsRoutes(router)
}

func (apC *AccountPasswordController) registerAccoutPasswordsRoutes(router *gin.RouterGroup) {
	router.POST("/accountPassword", apC.CreatAccountPassword)
	router.GET("/accountPassword", apC.GetAllAccountsPasswords)
	// router.GET("/accountPassword/:id", apC.GetAppPasswordById)
	router.GET("/accountPassword/:name", apC.GetAppPasswordByServiceName)
	router.PATCH("/accountPassword", apC.EditAccountPassword)
	router.DELETE("/accountPassword/:name", apC.DeleteServicePassword)
}

//DI
func NewAccPasswordControllerModule(service *services.AccountPasswordService) *AccountPasswordController {
	return &AccountPasswordController{service: service}
}
