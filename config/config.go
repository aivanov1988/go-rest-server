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

type DatabaseConfig struct {
	Database string
	User     string
	Password string
	Addr     string
}

type AppConfig struct {
	Port   int
	DbConf DatabaseConfig
}

var Config AppConfig

func InitConfig() error {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&privConfig); err != nil {
		return err
	}

	Config = AppConfig{
		Port: privConfig.PORT,
		DbConf: DatabaseConfig{
			Database: privConfig.POSTGRES_DB,
			User:     privConfig.POSTGRES_USER,
			Password: privConfig.POSTGRES_PASSWORD,
			Addr:     privConfig.POSTGRES_HOST,
		},
	}

	return nil
}
