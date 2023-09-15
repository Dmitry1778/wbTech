package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"wbTech/cmd/config"
	"wbTech/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
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

// Полученные данные записывать в БД
func CreateNewDb() *DB {
	cfg := config.GetConfig(&config.Config{})
	db := DB{}
	PostgresConn(context.TODO(), 5, cfg.Storage)
	return &db
}

// Сохранение Storage в БД
func (db *DB) AddStorage(o domain.Order, d domain.Delivery, p domain.Payment, item domain.Items) (err error) {
	fmt.Println("Это AddStorage method")
	tx, err := db.pool.Begin(context.Background())
	if err != nil {
		return
	}
	defer tx.Rollback(context.Background())

	//Order
	err = tx.QueryRow(context.Background(), `INSERT INTO order (OrderUID, TrackNumber, Entry, Locale, InternalSignature, 
		CustomerId, DeliveryService, Shardkey, SmId, DateCreated, OofShard) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id`,
		o.OrderUid, o.TrackNumber, o.Entry, o.Locale, o.InternalSignature,
		o.CustomerId, o.DeliveryService, o.Shardkey, o.SmId, o.DateCreated, o.OofShard).Scan()
	if err != nil {
		log.Printf("Unable to insert data (orders): %v\n", err)
		return err
	}

	//Delivery
	err = tx.QueryRow(context.Background(), `INSERT INTO delivery (Name, Phone, Zip, City, Address, 
		Region, Email) values ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`, d.Name, d.Phone, d.Zip, d.City, d.Address, d.Region, d.Email).Scan()
	if err != nil {
		log.Printf("Unable to insert data (delivery): %v\n", err)
		return err
	}
	//Payment
	err = tx.QueryRow(context.Background(), `INSERT INTO payment (Transaction, RequestId, Currency, Provider, 
		Amount, PaymentDt, Bank, DeliveryCost, GoodsTotal, CustomFee) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id`, p.Transaction, p.RequestId, p.Currency, p.Provider,
		p.Amount, p.PaymentDt, p.Bank, p.Bank, p.DeliveryCost, p.GoodsTotal, p.CustomFee).Scan()
	if err != nil {
		log.Printf("Unable to insert data (payment): %v\n", err)
		return err
	}

	//Items
	err = tx.QueryRow(context.Background(), `INSERT INTO items (ChrtId, TrackNumber, Price, Rid, Name, Sale, 
			Size, TotalPrice, NmId, Brand, Status) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id`, item.ChrtId, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale,
		item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status).Scan()
	if err != nil {
		log.Printf("Unable to insert data (items): %v\n", err)
		return err
	}

	return nil
}

//Реализовать кэширование полученных данных в сервисе (сохранять in memory)
