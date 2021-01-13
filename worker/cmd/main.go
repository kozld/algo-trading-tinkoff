package main

import (
	"database/sql"
	"fmt"
	"google.golang.org/grpc"
	"tinkoff-trade-bot/worker"
	"tinkoff-trade-bot/worker/config"
	pb "tinkoff-trade-bot/worker/trader/proto"
)

func main() {

	cfg := config.GetConfig()

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", cfg.PgHost, cfg.PgUser, cfg.PgPassword, cfg.PgName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	conn, _ := grpc.Dial(fmt.Sprintf("%s:%d", cfg.TraderHost, cfg.TraderPort), grpc.WithInsecure())
	trader := pb.NewTraderClient(conn)

	worker := worker.NewWorkerService(db, trader, &cfg)
	worker.Start()
}