package worker

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
	"tinkoff-trade-bot/worker/config"
	pb "tinkoff-trade-bot/worker/trader/proto"
)

type WorkerService struct {
	db *sql.DB
	trader pb.TraderClient
	config *config.WorkerConfig
}

func NewWorkerService(db *sql.DB, trader pb.TraderClient, cfg *config.WorkerConfig) *WorkerService {
	return &WorkerService{
		db: db,
		trader: trader,
		config: cfg,
	}
}

type Action struct {
	id int
	actionType string
	ticker string
	price float64
	qty int
}

func (ws *WorkerService) Start() {

	for {
		rows, err := ws.db.Query("SELECT id, type, ticker, price, qty FROM actions WHERE closed = FALSE LIMIT 10")
		if err != nil {
			log.Printf("error: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}
		defer rows.Close()
		var actions []Action

		for rows.Next() {
			a := Action{}
			err := rows.Scan(&a.id, &a.actionType, &a.ticker, &a.price, &a.qty)
			if err != nil {
				fmt.Println(err)
				continue
			}
			actions = append(actions, a)
		}
		for _, a := range actions {
			fmt.Println(a.id, a.actionType, a.ticker, a.price, a.qty)

			qty := a.qty / 20
			if qty == 0 {
				qty = 1
			}

			_, err := ws.trader.MarketOrder(context.Background(),
				&pb.MarketOrderRequest{
					Type:   a.actionType,
					Ticker: a.ticker,
					Price: float32(a.price),
					Qty:    int32(qty),
				})

			if err != nil {
				log.Printf("error: %v", err)
			} else {
				log.Println("success order")
				_, err := ws.db.Exec("UPDATE actions SET closed = TRUE WHERE id = $1", a.id)
				if err != nil{
					panic(err)
				}
			}
		}

		time.Sleep(2 * time.Second)
	}
}