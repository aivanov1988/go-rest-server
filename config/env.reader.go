package config

import (
	"github.com/spf13/viper"
)

type configType struct {
	PORT              int
	POSTGRES_DB       string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_HOST     string
}

var privConfig configType

func readEnv() error {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&privConfig); err != nil {
		return err
	}
	return nil
}
