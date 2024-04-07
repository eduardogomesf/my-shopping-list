package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	APPPort string `mapstructure:"APP_PORT"`
	PGHost  string `mapstructure:"PG_HOST"`
	PGPort  string `mapstructure:"PG_PORT"`
	PGUser  string `mapstructure:"PG_USER"`
	PGPass  string `mapstructure:"PG_PASS"`
	PGDB    string `mapstructure:"PG_DB"`
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
