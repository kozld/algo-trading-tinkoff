package trader

import (
	"context"
	"fmt"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"log"
	"strconv"
	"time"
	"tinkoff-trade-bot/trader/config"
	pb "tinkoff-trade-bot/trader/proto"
)

type TraderService struct {
	client *sdk.RestClient
	config *config.TraderConfig
}

func NewTraderService(cfg *config.TraderConfig) *TraderService {
	client := sdk.NewRestClient(cfg.TinkoffToken)

	return &TraderService{
		client: client,
		config: cfg,
	}
}

func (ts *TraderService) MarketOrder(ctx context.Context, req *pb.MarketOrderRequest) (*pb.MarketOrderResponse, error) {
	fmt.Println("[REQUEST]", req)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf("Получение инструменов по тикеру %s", req.Ticker)
	instruments, err := ts.client.InstrumentByTicker(ctx, req.Ticker)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", instruments)

	if len(instruments) > 0 {
		figi := instruments[0].FIGI
		actionType := sdk.BUY
		if req.Type == "SELL" {
			actionType = sdk.SELL
		}

		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		log.Printf("[%s] Выставление рыночной заявки  %s (%s)", req.Type, figi, req.Ticker)
		// Выставление рыночной заявки для счета по-умолчанию
		placedOrder, err := ts.client.MarketOrder(ctx, sdk.DefaultAccount, figi, int(req.Qty), actionType)
		if err != nil {
			log.Println(err)
		}
		log.Printf("[RESULT] %+v\n", placedOrder)

		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		log.Printf("[SELL] Выставление лимитной заявки %s (%s)", figi, req.Ticker)

		// Рассчитываем будущую цену продажи
		price := float64(req.Price)
		price += price / 100 * ts.config.TakeProfit
		formattedPrice := fmt.Sprintf("%.2f", price)
		roundedPrice, _ := strconv.ParseFloat(formattedPrice, 64)

		// Выставление лимитной заявки для счета по-умолчанию
		placedOrder, err = ts.client.LimitOrder(ctx, sdk.DefaultAccount, figi, int(req.Qty), sdk.SELL, roundedPrice)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("[RESULT] %+v\n", placedOrder)
	}

	return &pb.MarketOrderResponse{}, nil
}