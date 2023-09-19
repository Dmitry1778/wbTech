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
	pool *pgx.Conn
}

var order domain.Order

func GetStorageJSON() (s []byte, err error) {
	data, err := json.MarshalIndent(order, "", " ") // data go publisher
	if err != nil {
		log.Fatalf("Fail marshaling JSON: %v\n", err)
	}
	fmt.Printf("Send massage:%s\n", string(data))
	return data, nil
}

// NewDB NewDB NewCreateDb CreateNewDb Полученные данные записывать в БД
func NewDB(pool *pgx.Conn) *DB {
	return &DB{pool: pool}
}

// AddStorage Сохранение Storage в БД
func (db *DB) AddStorage(o domain.Order, d domain.Delivery, p domain.Payment, item domain.Items) (err error) {
	tx, err := db.pool.Begin(context.Background())
	if err != nil {
		return
	}
	defer tx.Rollback(context.Background())

	//Order
	_, err = tx.Exec(context.Background(), `INSERT INTO orders (OrderUID, TrackNumber, Entry, Locale, InternalSignature, 
		CustomerId, DeliveryService, Shardkey, SmId, DateCreated, OffShard) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id`,
		o.OrderUid, o.TrackNumber, o.Entry, o.Locale, o.InternalSignature,
		o.CustomerId, o.DeliveryService, o.Shardkey, o.SmId, o.DateCreated, o.OffShard)
	if err != nil {
		log.Printf("Unable to insert data (orders): %v\n", err)
		return err
	}

	//Delivery
	_, err = tx.Exec(context.Background(), `INSERT INTO delivery (Name, Phone, Zip, City, Addres, 
		Region, Email) values ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`, d.Name, d.Phone, d.Zip, d.City, d.Address, d.Region, d.Email)
	if err != nil {
		log.Printf("Unable to insert data (delivery): %v\n", err)
		return err
	}
	//Payment
	_, err = tx.Exec(context.Background(), `INSERT INTO payments (Transaction, Currency, Provider, 
		Amount, PaymentDt, Bank, DeliveryCost, GoodsTotal, CustomFee) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`, p.Transaction, p.Currency, p.Provider,
		p.Amount, p.PaymentDt, p.Bank, p.DeliveryCost, p.GoodsTotal, p.CustomFee)
	if err != nil {
		log.Printf("Unable to insert data (payment): %v\n", err)
		return err
	}

	//Items
	_, err = tx.Exec(context.Background(), `INSERT INTO items (ChrtId, Price, Rid, Name, Sale, 
			Size, TotalPrice, NmId, Brand, Status) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id`, item.ChrtId, item.Price, item.Rid, item.Name, item.Sale,
		item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status)
	if err != nil {
		log.Printf("Unable to insert data (items): %v\n", err)
		return err
	}
	return nil
}

//Реализовать кэширование полученных данных в сервисе (сохранять in memory)
