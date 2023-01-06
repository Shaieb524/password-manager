package middlewares

import (
	"net/http"
	"password-manager/src/utils/helper"

	"github.com/gin-gonic/gin"
)

func JWTAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helper.ValidateJWT(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
