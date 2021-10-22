package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbUser     string `required:"true" split_words:"true"`
	DbPassword string `required:"true" split_words:"true"`
	DbHost     string `required:"true" split_words:"true"`
	DbPort     string `required:"true" split_words:"true"`
	DbName     string `required:"true" split_words:"true"`
}

func LoadConfig() Config {
	var con Config
	if err := envconfig.Process("", &con); err != nil {
		log.Fatal(err)
	}
	return con
}
