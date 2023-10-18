package server

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"wbTech/internal/db"
	"wbTech/internal/domain"
)

func Publisher(s stan.Conn) (err error) {
	item := domain.NewItem{ChrtId: 2, TrackNumber: "#433", Price: 9999, Rid: "rid 1", Name: "Jeans", Sale: 9, Size: "M", TotalPrice: 15, NmId: 1, Brand: "Adidas", Status: 0}
	payment := domain.NewPayment{Transaction: "tran 1", Currency: "Rub", Provider: "Provider 1", Amount: 47, PaymentDt: 2, Bank: "VTB", DeliveryCost: 7, GoodsTotal: 3}
	order := domain.NewOrder{OrderUid: "6", Entry: "2", InternalSignature: "IS 2", Payment: payment, Items: item, Locale: "Ru", CustomerId: "2", TrackNumber: "2", DeliveryService: "DS 2", Shardkey: "SK 2", SmId: 2}
	orderData, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("Error marshal: %v\n", err)
	}
	orderData, err = db.TestSubscribeOnline(orderData)
	if err != nil {
		log.Fatal("Message don't to send", err)
	}

	ackHandler := func(ackedNuid string, err error) {
		if err != nil {
			log.Printf("%s: error publishing msg id %s:\n", ackedNuid, err.Error())
		} else {
			log.Printf("%s: received ack for msg id:\n", ackedNuid)
		}
	}

	nuid, err := s.PublishAsync("MessengerFoo", orderData, ackHandler)
	if err != nil {
		log.Printf("%s: error publishing msg %s:\n", nuid, err.Error())
	}

	return nil
}
