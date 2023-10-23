package config

import (
	"log"

	"github.com/num30/config"
)

type Config struct {
	DatabaseUrl       string `default:"postgres://test:test@postgres:5432/test?sslmode=disable" validate:"required"`
	AppEnv            string `default:"development"                                             validate:"required"`
	MigrationsPath    string `default:"./internal/database/migrations"                          validate:"required"`
	HttpServerAddress string `default:"0.0.0.0:8000"                                            validate:"required"`
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
