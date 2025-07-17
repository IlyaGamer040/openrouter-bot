package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ApiToken string `env:"API_TOKEN" env-default:"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"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Println(err)
	}

	return &cfg, nil
}
