package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Environment string `required:"true" split_words:"true"`

	SessionKey    string `required:"true" split_words:"true"`
	RunServerPort string `required:"true" split_words:"true"`

	DBUser     string `required:"true" split_words:"true"`
	DBPassword string `required:"true" split_words:"true"`
	DBHost     string `required:"true" split_words:"true"`
	DBPort     string `required:"true" split_words:"true"`
	DBName     string `required:"true" split_words:"true"`

	AwsDefaultRegion   string `required:"true" split_words:"true"`
	AwsCognitoPoolID   string `required:"true" split_words:"true"`
	AwsCognitoClientID string `required:"true" split_words:"true"`
}

func LoadConfig() Config {
	var con Config
	if err := envconfig.Process("", &con); err != nil {
		log.Fatal(err)
	}

	return con
}
