package services

import (
	"github.com/ebubekiryigit/golang-mongodb-rest-api-starter/models"
	"github.com/spf13/viper"
)

var Config *models.EnvConfig

func LoadConfig() {
	viper := viper.New()
	viper.AutomaticEnv()
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("MODE", "debug")
	viper.SetConfigType("dotenv")
	viper.SetConfigName(".env")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}

	if err := Config.Validate(); err != nil {
		panic(err)
	}
}
