package helper

import (
	"errors"
	"fmt"
	"strings"

	// "password-manager/src/models/database/user"

	"github.com/gin-gonic/gin"
	// "github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
)

// func CurrentUser(context *gin.Context) (user.User, error) {
// 	err := ValidateJWT(context)
// 	if err != nil {
// 		return user.User{}, err
// 	}
// 	token, _ := getToken(context)
// 	claims, _ := token.Claims.(jwt.MapClaims)
// 	userId := claims["id"].(string)

// 	user, err := user.FindUserById(uuid.Must(uuid.FromString(userId)))
// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func getToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractTokenFromRequest(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func ExtractTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
