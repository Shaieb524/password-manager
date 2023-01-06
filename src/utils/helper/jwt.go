package helper

import (
	"password-manager/src/models/database/user"
	"password-manager/src/utils/env"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// TODO Check how to pass this env
var globalEnv = env.NewEnv()
var privateKey = []byte(globalEnv.JwtPrivateKey)

func GenerateJWT(user user.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(globalEnv.TokenTil)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}
