package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server     Server
		PostgresDB PostgresDB
	}

	Server struct {
		Port string
	}

	PostgresDB struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
)

func LoadConfig() Config {
	var conf Config
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&conf); err != nil {
		panic(err)
	}

	return conf
}