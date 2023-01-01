package middlewares

import (
	"net/http"
	"password-manager/src/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.ValidateJWT(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
