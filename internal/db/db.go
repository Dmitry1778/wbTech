package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"wbTech/internal/domain"
)

type DB struct {
	pgClient *pgx.Conn
	csh      *Cache
}

func NewDB(pgClient *pgx.Conn) *DB {
	return &DB{pgClient: pgClient}
}

func TestSubscribeOnline(order []byte) (s []byte, err error) {
	if err != nil {
		log.Fatalf("Fail marshaling JSON: %v\n", err)
	}
	fmt.Printf("Send massage:%s\n", string(order))
	return order, nil
}

func (db *DB) getCacheFromDatabase() map[string]domain.NewOrder {
	buffer := make(map[string]domain.NewOrder)
	query := fmt.Sprintf("select orderid from ordernew where payload = $1")
	row, err := db.pgClient.Query(context.Background(), query)
	if err != nil {
		log.Printf("%v: unable to get id from database row:\n", err)
	}
	defer row.Close()

	var id string

	for row.Next() {
		if err = row.Scan(&id); err != nil {
			log.Printf("%v: unable to get id from database row:\n", err)
			return buffer
		}
		o, err := db.GetOrder(context.Background(), id)
		if err != nil {
			log.Printf("%v: unable to get order from database:\n", err)
		}
		buffer[id] = *o
	}
	return buffer
}

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
	if err != nil {
		log.Printf("Unable to insert data (payload): %v\n", err)
		return err
	}
	return err
}

func (db *DB) SetCacheInstance(csh *Cache) {
	db.csh = csh
}
