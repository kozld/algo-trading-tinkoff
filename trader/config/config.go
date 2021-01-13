package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type TraderConfig struct {
	RPCListenPort int32  `env:"RPC_LISTEN_PORT" envDefault:"5300"`
	TinkoffToken string `env:"TINKOFF_TOKEN,required"`
	TakeProfit float64 `env:"TAKE_PROFIT" envDefault:"2.5"`
}

func GetConfig() TraderConfig {
	cfg := &TraderConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal("Cannot parse initial ENV vars: ", err)
	}
	return *cfg
}