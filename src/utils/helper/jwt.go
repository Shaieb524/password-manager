package helper

import (
	"fmt"
	"os"
	"password-manager/src/models/database/user"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user user.User) (string, error) {
	fmt.Println("privateKey : ", privateKey)

	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	fmt.Println("tokenTTL : ", tokenTTL)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}
