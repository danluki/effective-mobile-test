package config

import (
	"log"

	"github.com/num30/config"
)

type Config struct {
	DatabaseUrl       string `default:""             validate:"required"`
	AppEnv            string `default:"development"  validate:"required"`
	HttpServerAddress string `default:"0.0.0.0:8000" validate:"required"`
}

func LoadConfig(path string) (*Config, error) {
	var conf Config

	err := config.NewConfReader(path + "/" + "conf").Read(&conf)
	if err != nil {
		panic(err)
	}

	log.Println("Config ", conf)

	return &conf, err
}
