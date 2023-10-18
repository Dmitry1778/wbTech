package new_order_subsriber

import (
	"context"
	"encoding/json"
	"wbTech/internal/domain"
)

type Worker struct {
	database Database
	cache    Cache
}

func NewWorker(database Database, cache Cache) *Worker {
	return &Worker{
		database: database,
		cache:    cache,
	}
}

func (w *Worker) Process(jsonPayload []byte) error {
	order := &domain.NewOrder{}
	err := json.Unmarshal(jsonPayload, &order)
	if err != nil {
		return err
	}
	err = w.database.PutOrder(context.Background(), *order)
	if err != nil {
		return err
	}
	w.cache.Set(order.OrderUid, order)
	return nil
}
