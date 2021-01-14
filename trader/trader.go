package trader

import (
	"context"
	"fmt"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"log"
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

func (ts *TraderService) getFIGIbyTicker(ticker string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	instruments, err := ts.client.InstrumentByTicker(ctx, ticker)
	if err != nil {
		log.Println(err)
	}

	if len(instruments) > 0 {
		return instruments[0].FIGI, nil
	}

	return "", fmt.Errorf("инструмент не найден")
}

func (ts *TraderService) CreateMarketOrder(ctx context.Context, req *pb.CreateMarketOrderRequest) (*pb.CreateMarketOrderResponse, error) {

	name, err := ts.getFIGIbyTicker(req.Ticker)
	if err != nil {
		log.Println(err)
	}

	if name != "" {
		opType := sdk.BUY
		if req.OpType == "SELL" {
			opType = sdk.SELL
		}

		qty := int(req.Qty)
		log.Printf("Выставление рыночной заявки: %s #%s x %d", opType, req.Ticker, qty)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		placedOrder, err := ts.client.MarketOrder(ctx, sdk.DefaultAccount, name, qty, opType)
		if err != nil {
			log.Println(err)
		}

		log.Printf("Результат: %+v\n", placedOrder)
	}

	return &pb.CreateMarketOrderResponse{}, nil
}

func (ts *TraderService) CreateLimitOrder(ctx context.Context, req *pb.CreateLimitOrderRequest) (*pb.CreateLimitOrderResponse, error) {

	name, err := ts.getFIGIbyTicker(req.Ticker)
	if err != nil {
		log.Println(err)
	}

	if name != "" {
		opType := sdk.BUY
		if req.OpType == "SELL" {
			opType = sdk.SELL
		}

		price := float64(req.Price)
		qty := int(req.Qty)
		log.Printf("Выставление лимитной заявки: %s #%s $%.2f x %d", opType, req.Ticker, price, qty)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		placedOrder, err := ts.client.LimitOrder(ctx, sdk.DefaultAccount, name, qty, opType, price)
		if err != nil {
			log.Println(err)
		}

		log.Printf("Успех: %+v\n", placedOrder)
	}

	return &pb.CreateLimitOrderResponse{}, nil
}

func (ts *TraderService) TakeProfit(ctx context.Context, req *pb.TakeProfitRequest) (*pb.TakeProfitResponse, error) {

	return &pb.TakeProfitResponse{}, nil
}

func (ts *TraderService) StopLoss(ctx context.Context, req *pb.StopLossRequest) (*pb.StopLossResponse, error) {

	return &pb.StopLossResponse{}, nil
}