package worker

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
	"tinkoff-trade-bot/worker/config"
	pb "tinkoff-trade-bot/worker/trader/proto"
)

type WorkerService struct {
	db     *sql.DB
	trader pb.TraderClient
	config *config.WorkerConfig
}

func NewWorkerService(db *sql.DB, trader pb.TraderClient, cfg *config.WorkerConfig) *WorkerService {
	return &WorkerService{
		db:     db,
		trader: trader,
		config: cfg,
	}
}

type Operation struct {
	id     int
	opType string
	ticker string
	price  float64
	qty    int
}

func (ws *WorkerService) Start() {

	for {
		rows, err := ws.db.Query("SELECT id, type, ticker, price, qty FROM actions WHERE closed = FALSE LIMIT 10")
		if err != nil {
			log.Println(err)
			continue
		}

		defer rows.Close()

		var operations []Operation
		for rows.Next() {
			op := Operation{}
			err := rows.Scan(&op.id, &op.opType, &op.ticker, &op.price, &op.qty)
			if err != nil {
				log.Println(err)
				continue
			}
			operations = append(operations, op)
		}
		for _, a := range operations {

			log.Printf("Отправляю в Trader операцию: %s #%s $%.2f x %d", a.opType, a.ticker, a.price, a.qty)

			//qty := a.qty / 20
			//if qty == 0 {
			//	qty = 1
			//}

			placedOrder, err := ws.trader.CreateMarketOrder(context.Background(),
				&pb.CreateMarketOrderRequest{
					OpType: a.opType,
					Ticker: a.ticker,
					Qty:    int32(a.qty),
				})

			if err != nil {
				log.Println(err)
				continue
			}

			log.Printf("Ответ: %+v\n", placedOrder)

			_, err = ws.db.Exec("UPDATE actions SET closed = TRUE WHERE id = $1", a.id)
			if err != nil {
				panic(err)
			}
		}

		time.Sleep(2 * time.Second)
	}
}
