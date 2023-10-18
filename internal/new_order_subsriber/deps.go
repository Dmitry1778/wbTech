package new_order_subsriber

import (
	"context"
	"wbTech/internal/domain"
)

type Database interface {
	PutOrder(ctx context.Context, order domain.NewOrder) error
	GetOrder(ctx context.Context, id string) (*domain.NewOrder, error)
}

type Cache interface {
	Set(key string, value *domain.NewOrder)
}
