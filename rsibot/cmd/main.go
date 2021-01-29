package main

import (
	"log"
	"tinkoff-trade-bot/rsibot"
	"tinkoff-trade-bot/rsibot/config"
)

func main() {

	cfg := config.GetConfig()

	service, err := rsibot.NewRsiBotService(&cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	service.Start()
}
