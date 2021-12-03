package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBUser     string `required:"true" split_words:"true"`
	DBPassword string `required:"true" split_words:"true"`
	DBHost     string `required:"true" split_words:"true"`
	DBPort     string `required:"true" split_words:"true"`
	DBName     string `required:"true" split_words:"true"`

	FreeGroupName string
	AllChatName   string
}

func LoadConfig() Config {
	var con Config
	if err := envconfig.Process("", &con); err != nil {
		log.Fatal(err)
	}

	return con
}
