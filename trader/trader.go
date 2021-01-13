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
	//client := sdk.NewSandboxRestClient(cfg.TinkoffToken)
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//log.Println("Регистрация обычного счета в песочнице")
	//account, err := client.Register(ctx, sdk.AccountTinkoff)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Printf("%+v\n", account)
	//
	//ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//log.Println("Рисуем себе 100500 рублей в портфеле песочницы")
	//err = client.SetCurrencyBalance(ctx, account.ID, sdk.USD, 1000)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	client := sdk.NewRestClient(cfg.TinkoffToken)

	return &TraderService{
		client: client,
		config: cfg,
	}
}

func (ts *TraderService) MarketOrder(ctx context.Context, req *pb.MarketOrderRequest) (*pb.MarketOrderResponse, error) {
	fmt.Println("Request:", req)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf("Получение инструменов по тикеру %s", req.Ticker)
	// Получение инструмента по тикеру, возвращает массив инструментов потому что тикер уникален только в рамках одной биржи
	// но может совпадать на разных биржах у разных кампаний
	// Например: https://www.moex.com/ru/issue.aspx?code=FIVE и https://www.nasdaq.com/market-activity/stocks/FIVE
	// В этом примере получить нужную компанию можно проверив поле Currency
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

		log.Printf("Выставление рыночной заявки %s %s (%s)", req.Type, figi, req.Ticker)
		// Выставление рыночной заявки для счета по-умолчанию
		placedOrder, err := ts.client.MarketOrder(ctx, sdk.DefaultAccount, figi, int(req.Qty), actionType)
		if err != nil {
			log.Println(err)
		}
		log.Printf("ORDER: %+v\n", placedOrder)

		//operations, err := ts.client.Operations(ctx, sdk.DefaultAccount, time.Now().AddDate(0, 0, -7), time.Now(), figi)
		//if err != nil {
		//	log.Fatalln(err)
		//}
		//log.Printf("OPERATIONS: %+v\n", operations)
	}

	return &pb.MarketOrderResponse{}, nil
}