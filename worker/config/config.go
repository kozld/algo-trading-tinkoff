package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type WorkerConfig struct {
	PgHost string `env:"PG_HOST" envDefault:"localhost"`
	PgUser string `env:"PG_USER,required"`
	PgPassword string `env:"PG_PASSWORD,required"`
	PgName string `env:"PG_NAME,required"`
	TraderHost string `env:"TRADER_HOST,required"`
	TraderPort int32  `env:"TRADER_PORT" envDefault:"5300"`
}

func GetConfig() WorkerConfig {
	cfg := &WorkerConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal("Cannot parse initial ENV vars: ", err)
	}
	return *cfg
}