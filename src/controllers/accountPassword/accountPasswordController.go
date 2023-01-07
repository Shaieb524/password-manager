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
}

// @Tags AccountPassword
// @Summary create a new account password
// @Description create a new account password
// @Accept json
// @Produce json
// @param Authorization header string true "Bearer Auth, pls add bearer before"
// @Param data body accountPassword.AccountPasswordInputDto true "payload"
// @Success 200 {object} accountPassword.AccountPasswordInputDto
// Failure 400 {object} models.ErrorResponse
// Failure 404 {object} models.ErrorResponse
// Failure 500 {object} models.ErrorResponse
// @Router /accountPassword [post]
func (apC *AccountPasswordController) CreatAccountPassword(ctx *gin.Context) {
	var accPassword *accountPassword.AccountPasswordInputDto

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

// @Tags AccountPassword
// @Summary get all accounts passwords
// @Description get all accounts passwords
// @Accept json
// @Produce json
// @param Authorization header string true "Bearer Auth, pls add bearer before"
// @Success 200 {object} []accountPassword.AccountPassword
// Failure 400 {object} models.ErrorResponse
// Failure 404 {object} models.ErrorResponse
// Failure 500 {object} models.ErrorResponse
// @Router /accountPassword [get]
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

// @Tags AccountPassword
// @Summary get account password by service name
// @Description account password by service name
// @Accept json
// @Produce json
// @param Authorization header string true "Bearer Auth, pls add bearer before"
// @Param name path string true "name"
// @Success 200 {object} accountPassword.AccountPassword
// Failure 400 {object} models.ErrorResponse
// Failure 404 {object} models.ErrorResponse
// Failure 500 {object} models.ErrorResponse
// @Router /accountPassword/{serviceName} [get]
func (apC *AccountPasswordController) GetAppPasswordByServiceName(ctx *gin.Context) {
	accountName := ctx.Param("serviceName")

	if appPassword, err := apC.service.GetAppPasswordByServiceName(accountName); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, appPassword)
	}
}

// @Tags AccountPassword
// @Summary edit account password
// @Description get edit account password
// @Accept json
// @Produce json
// @param Authorization header string true "Bearer Auth, pls add bearer before"
// @Param data body accountPassword.AccountPasswordInputDto true "payload"
// @Success 200 {object} accountPassword.AccountPasswordInputDto
// Failure 400 {object} models.ErrorResponse
// Failure 404 {object} models.ErrorResponse
// Failure 500 {object} models.ErrorResponse
// @Router /accountPassword [patch]
func (apC *AccountPasswordController) EditAccountPassword(ctx *gin.Context) {
	var accPassword *accountPassword.AccountPasswordInputDto

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

// @Tags AccountPassword
// @Summary delete account password
// @Description delete account password
// @Accept json
// @Produce json
// @param Authorization header string true "Bearer Auth pls add bearer before"
// @Param name path string true "name"
// @Success 200 {object} accountPassword.AccountPassword
// Failure 400 {object} models.ErrorResponse
// Failure 404 {object} models.ErrorResponse
// Failure 500 {object} models.ErrorResponse
// @Router /accountPassword/{serviceName} [delete]
func (apC *AccountPasswordController) DeleteServicePassword(ctx *gin.Context) {
	accountName := ctx.Param("serviceName")

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
	router.GET("/accountPassword/:serviceName", apC.GetAppPasswordByServiceName)
	router.PATCH("/accountPassword", apC.EditAccountPassword)
	router.DELETE("/accountPassword/:serviceName", apC.DeleteServicePassword)
}

//DI
func NewAccPasswordControllerModule(service *services.AccountPasswordService) *AccountPasswordController {
	return &AccountPasswordController{service: service}
}
