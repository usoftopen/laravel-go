package config

import "github.com/spf13/viper"

type AppConfig struct {
	AppEnv     string
	AppEnvTest string
	AppEnvProd string

	AppPoliceUrl string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		AppEnv:     viper.GetString("APP_ENV"),
		AppEnvTest: "test",
		AppEnvProd: "prod",

		AppPoliceUrl: viper.GetString("APP_POLICE_URL"),
	}
}
