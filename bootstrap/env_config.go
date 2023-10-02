package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	DbHost                 string `mapstructure:"DB_HOST"`
	DbPort                 string `mapstructure:"DB_PORT"`
	DbUser                 string `mapstructure:"DB_USER"`
	DbPass                 string `mapstructure:"DB_PASS"`
	DbName                 string `mapstructure:"DB_NAME"`
	SecretKey              string `mapstructure:"SecretKey"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {
	var err error

	env := &Env{}

	viper.SetConfigFile(".env")

	err = viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Can't find the file .env: %v ", err.Error())
	}

	err = viper.Unmarshal(env)

	if err != nil {
		log.Fatalf("Environment can't be loaded: %v", err.Error())
	}

	return env
}
