package rsibot

import (
	"context"
	"fmt"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/markcheno/go-talib"
	"log"
	"time"
	"tinkoff-trade-bot/rsibot/config"
)

var figi string
var inPosition bool
var closes = make([]float64, 0)

type RsiBotService struct {
	client *sdk.RestClient
	config *config.RsiBotConfig
}

func NewRsiBotService(cfg *config.RsiBotConfig) (*RsiBotService, error) {
	client := sdk.NewRestClient(cfg.TinkoffToken)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	instruments, err := client.InstrumentByTicker(ctx, cfg.Ticker)
	if err != nil {
		return nil, err
	}

	if len(instruments) > 0 {
		figi = instruments[0].FIGI
	} else {
		return nil, fmt.Errorf("instruments by ticker not found")
	}

	return &RsiBotService{
		client: client,
		config: cfg,
	}, nil
}

func (rbs *RsiBotService) Start() {

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		candles, err := rbs.client.Candles(ctx, time.Now().AddDate(0, 0, -1), time.Now(), sdk.CandleInterval1Min, figi)
		if err != nil {
			log.Printf("error getting candles: %v", err)
			continue
		}

		for _, c := range candles {
			closes = append(closes, c.ClosePrice)
		}

		if len(closes) > rbs.config.RsiPeriod {
			rsi := talib.Rsi(closes, rbs.config.RsiPeriod)
			lastRsi := rsi[len(rsi)-1]
			log.Println(lastRsi)

			if lastRsi > rbs.config.RsiOverbought {
				inPosition = false
			}

			if lastRsi < rbs.config.RsiOversold {
				if inPosition {
					log.Println("It is oversold, but we already in position, nothing to do.")
				} else {
					log.Println("Oversold! Buy! Buy! Buy!")
					msg := fmt.Sprintf("#%s: $%.2f Oversold! Buy! Buy! Buy!", rbs.config.Ticker, closes[len(closes)-1])
					err = rbs.sendMessage(msg)
					if err != nil {
						log.Printf("error: %v", err)
					}
					inPosition = true
				}
			}
		}

		time.Sleep(60 * time.Second)
	}
}
