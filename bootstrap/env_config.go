package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv        string `mapstructure:"AppEnv"`
	ServerAddress string `mapstructure:"ServerAddress"`
	DbHost        string `mapstructure:"DbHost"`
	DbPort        string `mapstructure:"DbPort"`
	DbUser        string `mapstructure:"DbUser"`
	DbPass        string `mapstructure:"DbPass"`
	DbName        string `mapstructure:"DbName"`
	SecretKey     string `mapstructure:"SecretKey"`
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
