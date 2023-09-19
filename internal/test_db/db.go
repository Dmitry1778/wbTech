package test_db

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"wbTech/internal/domain"
)

func (db *DB) GetOrder(ctx context.Context, id string) (*domain.NewOrder, error) {
	row := db.pgClient.QueryRow(ctx, `select payload from ordernew where orderid = $1`, id)
	order := domain.NewOrder{}
	err := row.Scan(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (db *DB) PutOrder(ctx context.Context, order domain.NewOrder) error {
	orderPayload, err := json.Marshal(order)
	if err != nil {
		return err
	}
	_, err = db.pgClient.Exec(ctx, `insert into ordernew (orderid, payload) values ($1, $2)`, order.OrderUid, orderPayload)
	return err
}

type DB struct {
	pgClient *pgx.Conn
}

func NewDB(pgClient *pgx.Conn) *DB {
	return &DB{pgClient: pgClient}
}
