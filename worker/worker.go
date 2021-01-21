package worker

import (
	"context"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"log"
	"time"
	"tinkoff-trade-bot/worker/config"
	pb "tinkoff-trade-bot/worker/trader/proto"
)

type WorkerService struct {
	trader pb.TraderClient
	client *sdk.RestClient
	config *config.WorkerConfig
}

func NewWorkerService(trader pb.TraderClient, cfg *config.WorkerConfig) *WorkerService {
	client := sdk.NewRestClient(cfg.TinkoffToken)

	return &WorkerService{
		trader: trader,
		client: client,
		config: cfg,
	}
}

func (ws *WorkerService) Monitor() {

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		log.Println("Получение списка НЕ валютных активов портфеля для счета по-умолчанию")
		positions, err := ws.client.PositionsPortfolio(ctx, sdk.DefaultAccount)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%+v\n", positions)

		for _, position := range positions {
			if position.InstrumentType != "Stock" {
				continue
			}

			avgPrice := position.AveragePositionPrice.Value
			ticker := position.Ticker
			figi := position.FIGI

			ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			log.Println("Получение списка выставленных заявок(ордеров) для счета по-умолчанию")
			orders, err := ws.client.Orders(ctx, sdk.DefaultAccount)
			if err != nil {
				log.Fatalln(err)
			}

			freeLots := position.Lots
			orderExist := false
			for _, order := range orders {
				if order.FIGI == figi {
					freeLots -= order.RequestedLots
					orderExist = true
				}
			}

			if !orderExist || freeLots > 0 {
				sellPrice := avgPrice + avgPrice*0.025 // Take profit == 2.5
				ws.trader.CreateLimitOrder(context.Background(), &pb.CreateLimitOrderRequest{
					OpType: "SELL",
					Ticker: ticker,
					Price:  float32(sellPrice),
					Qty:    int32(freeLots),
				})
			}

		}
		time.Sleep(30 * time.Second)
	}
}
