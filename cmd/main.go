package main

import (
	"context"
	"wbTech/cmd/config"
	"wbTech/internal/db"
	"wbTech/internal/domain"
	"wbTech/internal/test_db"
)

type Storage struct {
	o domain.Order
	d domain.Delivery
	p domain.Payment
	i domain.Items
}

func main() {

	cfg := config.GetConfig(&config.Config{})
	//
	////Building a project
	//err := server.NatsConnectMethod()
	//if err != nil {
	//	panic(err)
	//}

	dbConn, err := db.PostgresConn(context.Background(), 3, cfg.Storage)
	if err != nil {
		panic(err)
	}
	database := test_db.NewDB(dbConn)
	err = database.PutOrder(context.Background(), domain.NewOrder{OrderUid: "sukablet", TrackNumber: "suka blet"})
	if err != nil {
		panic(err)
	}
	_, err = database.GetOrder(context.Background(), "311123123123")
	if err != nil {
		panic(err)
	}

	//newDb := db.NewDB(dbConn)
	//err = newDb.AddStorage(domain.Order{}, domain.Delivery{}, domain.Payment{}, domain.Items{})
	//if err != nil {
	//	panic(err)
	//}
}
