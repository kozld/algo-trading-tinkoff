package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type WorkerConfig struct {
	TraderHost string `env:"TRADER_HOST,required"`
	TraderPort int32  `env:"TRADER_PORT" envDefault:"5300"`
	TinkoffToken string `env:"TINKOFF_TOKEN,required"`
}

func GetConfig() WorkerConfig {
	cfg := &WorkerConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal("Cannot parse initial ENV vars: ", err)
	}
	return *cfg
}