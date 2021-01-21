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
	client          *sdk.RestClient
	config          *config.TraderConfig
}

func NewTraderService(cfg *config.TraderConfig) *TraderService {
	client := sdk.NewRestClient(cfg.TinkoffToken)

	return &TraderService{
		client:          client,
		config:          cfg,
	}
}

func (ts *TraderService) getFigiByTicker(ticker string) (string, error) {
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

	name, err := ts.getFigiByTicker(req.Ticker)
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
		placedOrder, err := ts.createMarketOrder(opType, name, qty)

		if err != nil {
			log.Println(err)
		} else {
			log.Printf("Успех: %+v\n", placedOrder)
		}
	}

	return &pb.CreateMarketOrderResponse{}, nil
}

func (ts *TraderService) createMarketOrder(opType sdk.OperationType, figi string, qty int) (*sdk.PlacedOrder, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	placedOrder, err := ts.client.MarketOrder(ctx, sdk.DefaultAccount, figi, qty, opType)
	if err != nil {
		log.Println(err)
	}

	return &placedOrder, nil
}

func (ts *TraderService) CreateLimitOrder(ctx context.Context, req *pb.CreateLimitOrderRequest) (*pb.CreateLimitOrderResponse, error) {

	name, err := ts.getFigiByTicker(req.Ticker)
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

		placedOrder, err := ts.createLimitOrder(opType, name, price, qty)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("Успех: %+v\n", placedOrder)
		}
	}

	return &pb.CreateLimitOrderResponse{}, nil
}

func (ts *TraderService) createLimitOrder(opType sdk.OperationType, figi string, price float64, qty int) (*sdk.PlacedOrder, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	formattedPrice := fmt.Sprintf("%.2f", price)
	roundedPrice, _ := strconv.ParseFloat(formattedPrice, 64)

	placedOrder, err := ts.client.LimitOrder(ctx, sdk.DefaultAccount, figi, qty, opType, roundedPrice)
	if err != nil {
		return &sdk.PlacedOrder{}, err
	}

	return &placedOrder, nil
}

