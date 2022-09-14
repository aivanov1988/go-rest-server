package config

import (
	"github.com/spf13/viper"
)

type configType struct {
	PORT int
}

var Config configType

func InitConfig() error {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return err
	}

	return nil
}
