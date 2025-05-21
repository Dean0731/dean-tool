package config

import (
	"errors"
	"github.com/dean0731/dean-tool/exception"
	"github.com/spf13/viper"
)

func LoadConfig(dir string) {
	viper.AddConfigPath(".")
	viper.AddConfigPath(dir)
	viper.SetConfigName(ConfigurationFileName)
	viper.SetConfigType(ConfigurationFileType)

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			panic(exception.FileFormatError.SetMessage(err.Error()))
		}
	}
}
