package config

import (
	"os"

	"github.com/spf13/viper"
)

var App *Config

const (
	environment = "ENVIRONMENT"
)

func Init() {
	if os.Getenv(environment) != "test" {
		viper.SetConfigName("application")
	} else {
		viper.SetConfigName("test")
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	App = &Config{}
}
