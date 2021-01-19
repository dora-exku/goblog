package config

import (
	"goblog/pkg/logger"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var Viper *viper.Viper

type MapStr map[string]interface{}

func init() {

	// 初始化
	Viper = viper.New()
	// 设置文件名
	Viper.SetConfigName(".env")
	// 设置文件类型
	Viper.SetConfigType("env")
	// 目录
	Viper.AddConfigPath(".")

	err := Viper.ReadInConfig()
	logger.LogError(err)

	Viper.SetEnvPrefix("appenv")

	Viper.AutomaticEnv()
}

func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}
	return envName
}

func Add(name string, configuration map[string]interface{}) {
	Viper.Set(name, configuration)
}

func Get(path string, defaultValue ...interface{}) interface{} {

	if !Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return Viper.Get(path)
}

func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(path, defaultValue...))
}

func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(Get(path, defaultValue...))
}

func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}
