package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	iss "github.com/go-ozzo/ozzo-validation/is"
)

type EnvConfig struct {
	ServerPort                 string `mapstructure:"SERVER_PORT"`
	ServerAddr                 string `mapstructure:"SERVER_ADDR"`
	MongodbUri                 string `mapstructure:"MONGO_URI"`
	MongodbDatabase            string `mapstructure:"MONGO_DATABASE"`
	UseRedis                   bool   `mapstructure:"USE_REDIS"`
	RedisDefaultAddr           string `mapstructure:"REDIS_DEFAULT_ADDR"`
	JWTSecretKey               string `mapstructure:"JWT_SECRET"`
	JWTAccessExpirationMinutes int    `mapstructure:"JWT_ACCESS_EXPIRATION_MINUTES"`
	JWTRefreshExpirationDays   int    `mapstructure:"JWT_REFRESH_EXPIRATION_DAYS"`
	Mode                       string `mapstructure:"MODE"`
}

func (envConfigPointer *EnvConfig) Validate() error {
	return validation.ValidateStruct(envConfigPointer,
		validation.Field(&envConfigPointer.ServerPort, iss.Port),
		validation.Field(&envConfigPointer.ServerAddr, validation.Required),

		validation.Field(&envConfigPointer.MongodbUri, validation.Required),
		validation.Field(&envConfigPointer.MongodbDatabase, validation.Required),

		validation.Field(&envConfigPointer.UseRedis, validation.In(true, false)),
		validation.Field(&envConfigPointer.RedisDefaultAddr),

		validation.Field(&envConfigPointer.JWTSecretKey, validation.Required),
		validation.Field(&envConfigPointer.JWTAccessExpirationMinutes, validation.Required),
		validation.Field(&envConfigPointer.JWTRefreshExpirationDays, validation.Required),

		validation.Field(&envConfigPointer.Mode, validation.In("debug", "release")),
	)
}
