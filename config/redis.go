package config

import "github.com/spf13/viper"

type RedisConfig struct {
	Host     string
	Port     string
	Database string
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		Host:     viper.GetString("REDIS_HOST"),
		Port:     viper.GetString("REDIS_PORT"),
		Database: viper.GetString("REDIS_DATABASE"),
	}
}
