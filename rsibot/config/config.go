package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type RsiBotConfig struct {
	Ticker              string  `env:"TICKER,required"`
	RsiPeriod           int     `env:"RSI_PERIOD" envDefault:"9"`
	RsiOverbought       float64 `env:"RSI_OVERBOUGHT" envDefault:"70"`
	RsiOversold         float64 `env:"RSI_OVERSOLD" envDefault:"10"`
	TelegramBotToken    string  `env:"TELEGRAM_BOT_TOKEN,required"`
	TelegramChannelName string  `env:"TELEGRAM_CHANNEL_NAME,required"`
	TinkoffToken        string  `env:"TINKOFF_TOKEN,required"`
}

func GetConfig() RsiBotConfig {
	cfg := &RsiBotConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal("Cannot parse initial ENV vars: ", err)
	}
	return *cfg
}
