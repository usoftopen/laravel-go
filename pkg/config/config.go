package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// 加载配置文件
func LoadConfig(configName string, configType string) {

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("加载配置文件失败: %s \n", err))
	}
}
