package config

import (
	"log"

	"github.com/spf13/viper"
)

var config Config

type Config struct {
	Port int `mapstructure:"PORT"`
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Error occurs when loading environment variables", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln("Error occurs when unmarshaling environment variables", err)
	}

	setDefault()
}

func setDefault() {
	if config.Port == 0 {
		config.Port = 3000
	}
}

func GetPort() int {
	return config.Port
}
