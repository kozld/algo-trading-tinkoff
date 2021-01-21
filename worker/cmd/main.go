package main

import (
	"fmt"
	"google.golang.org/grpc"
	"tinkoff-trade-bot/worker"
	"tinkoff-trade-bot/worker/config"
	pb "tinkoff-trade-bot/worker/trader/proto"
)

func main() {

	cfg := config.GetConfig()

	conn, _ := grpc.Dial(fmt.Sprintf("%s:%d", cfg.TraderHost, cfg.TraderPort), grpc.WithInsecure())
	trader := pb.NewTraderClient(conn)

	service := worker.NewWorkerService(trader, &cfg)
	service.Monitor()
}