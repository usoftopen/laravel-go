package config

import (
	"laravel-go/pkg/orm/config"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Default     string
	Connections map[string]config.ConnParam
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Default: viper.GetString("DB_CONNECTION"),
		Connections: map[string]config.ConnParam{
			"mysql": {
				Driver:   viper.GetString("DB_DRIVER"),
				Host:     viper.GetString("DB_HOST"),
				Port:     viper.GetString("DB_PORT"),
				Username: viper.GetString("DB_USERNAME"),
				Password: viper.GetString("DB_PASSWORD"),
				Database: viper.GetString("DB_DATABASE"),
			},
			"mysql-other": {
				Driver:   viper.GetString("DB_OTHER_DRIVER"),
				Host:     viper.GetString("DB_OTHER_HOST"),
				Port:     viper.GetString("DB_OTHER_PORT"),
				Username: viper.GetString("DB_OTHER_USERNAME"),
				Password: viper.GetString("DB_OTHER_PASSWORD"),
				Database: viper.GetString("DB_OTHER_DATABASE"),
			},
		},
	}
}
