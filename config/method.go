package config

import (
	"github.com/dean0731/dean-tool/log"
	"github.com/spf13/viper"
)

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func SetKeyValue(key string, value any) {
	viper.Set(key, value)
}

func SetKeyValueIfBlank(key string, value any) {
	if GetString(key) == "" {
		SetKeyValue(key, value)
	}
	printConfig(key)
}

func GetStringWithDefault(key string, defaultValue string) string {
	value := GetString(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func printConfig(k string) {
	if k == "" {
		return
	}
	if value := GetString(k); value != "" {
		log.Debugf("%s: %v", k, value)
	}
}
