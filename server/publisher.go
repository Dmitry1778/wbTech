package server

import (
	"log"
	"wbTech/internal/db"

	"github.com/nats-io/stan.go"
)

//type Publisher struct {
//	pc *stan.Conn
//}

func Publish(p stan.Conn) (err error) {
	data, err := db.GetStorageJSON()
	if err != nil {
		log.Fatalf("Fail get storage fail on publish: %v\n", err)
	}
	err = p.Publish("MessengerFoo", []byte(data))
	if err != nil {
		log.Fatal("Message don't to send", err)
	}
	return nil
}
