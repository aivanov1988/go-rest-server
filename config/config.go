package config

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

	if err := readEnv(); err != nil {
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
