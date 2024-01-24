package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB Sqlite3
}

type Sqlite3 struct {
	Path string
}

func NewConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{
		DB: Sqlite3{
			Path: viper.GetString("DB_PATH"),
		},
	}

	return cfg, nil
}
