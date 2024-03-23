package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	APPPort string `mapstructure:"APP_PORT"`
}

func LoadConfig(path string) *AppConfig {
	var conf *AppConfig
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	return conf
}
