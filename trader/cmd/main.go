package main

import (
	"tinkoff-trade-bot/trader"
	"tinkoff-trade-bot/trader/config"
)

func main() {

	cfg := config.GetConfig()
	service := trader.NewTraderService(&cfg)

	gRPCServer := trader.NewGRPCServer(service, &cfg)
	gRPCServer.Start()
}