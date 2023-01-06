package env

import (
	"fmt"

	"github.com/spf13/viper"
)

type Env struct {
	ServerHost string `mapstructure:"SERVER_HOST"`
	ServerPort string `mapstructure:"SERVER_PORT"`

	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`

	JwtPrivateKey string `mapstructure:"JWT_PRIVATE_KEY"`
	TokenTil      string `mapstructure:"TOKEN_TTL"`
}

var globalEnv = Env{}

func GetEnv() Env {
	// TODO check when we should return this or *
	return globalEnv
}

func NewEnv() *Env {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Print("cannot read cofiguration", err)
	}

	viper.SetDefault("TIMEZONE", "UTC")

	err = viper.Unmarshal(&globalEnv)
	if err != nil {
		fmt.Print("environment cant be loaded:", err)
	}

	return &globalEnv
}
