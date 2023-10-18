package test_db

//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/jackc/pgx/v5"
//	"log"
//	"wbTech/internal/domain"
//)
//
//var order domain.NewOrder
//
//type DB struct {
//	pool *pgx.Conn
//}
//
//// NewDB NewCreateDb CreateNewDb Полученные данные записывать в БД
//func NewDB(pool *pgx.Conn) *DB {
//	return &DB{pool: pool}
//}
//
//func GetStorageJSON() (s []byte, err error) {
//	data, err := json.MarshalIndent(order, "", " ") // data go publisher
//	if err != nil {
//		log.Fatalf("Fail marshaling JSON: %v\n", err)
//	}
//	fmt.Printf("Send massage:%s\n", string(data))
//	return data, nil
//}
//
//// AddStorage Сохранение Storage в БД
//func (db *DB) AddStorage(o domain.NewOrder) (err error) {
//
//	//Order
//	_, err = db.pool.Exec(context.Background(), `INSERT INTO orders (OrderUID, TrackNumber, Entry, Locale, InternalSignature,
//		CustomerId, DeliveryService, Shardkey, SmId, DateCreated, OffShard) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
//		o.OrderUid, o.TrackNumber, o.Entry, o.Locale, o.InternalSignature,
//		o.CustomerId, o.DeliveryService, o.Shardkey, o.SmId, o.DateCreated, o.OofShard)
//	if err != nil {
//		log.Printf("Unable to insert data (orders): %v\n", err)
//		return err
//	}
//
//	//Delivery
//	_, err = db.pool.Exec(context.Background(), `INSERT INTO delivery (Name, Phone, Zip, City, Addres,
//		Region, Email) values ($1, $2, $3, $4, $5, $6, $7)`, o.Delivery.Name, o.Delivery.Phone, o.Delivery.Zip, o.Delivery.City, o.Delivery.Address, o.Delivery.Region, o.Delivery.Email)
//	if err != nil {
//		log.Printf("Unable to insert data (delivery): %v\n", err)
//		return err
//	}
//	//Payment
//	_, err = db.pool.Exec(context.Background(), `INSERT INTO payments (Transaction, Currency, Provider,
//		Amount, PaymentDt, Bank, DeliveryCost, GoodsTotal, CustomFee) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, o.Payment.Transaction, o.Payment.Currency, o.Payment.Provider,
//		o.Payment.Amount, o.Payment.PaymentDt, o.Payment.Bank, o.Payment.DeliveryCost, o.Payment.GoodsTotal, o.Payment.CustomFee)
//	if err != nil {
//		log.Printf("Unable to insert data (payment): %v\n", err)
//		return err
//	}
//
//	//Items
//	//for _, item := range o.Items {
//	//	_, err = db.pool.Exec(context.Background(), `INSERT INTO items (ChrtId, Price, Rid, Name, Sale,
//	//		Size, TotalPrice, NmId, Brand, Status) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, item.ChrtId, item.Price, item.Rid, item.Name, item.Sale,
//	//		item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status)
//	//}
//	//if err != nil {
//	//	log.Printf("Unable to insert data (items): %v\n", err)
//	//	return err
//	//}
//	return nil
//}

//Реализовать кэширование полученных данных в сервисе (сохранять in memory)
