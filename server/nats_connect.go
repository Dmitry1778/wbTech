package server

import (
	"github.com/nats-io/stan.go"
	"log"
)

func NatsConnectMethod() (err error) {
	sc, err := stan.Connect("test-cluster", "test-id", stan.NatsURL(":4444"))
	if err != nil {
		log.Fatal("fail to connect:", err)
	}
	defer sc.Close()

	err = Publish(sc)
	if err != nil {
		panic("Publish method fail:")
	}

	err = Subscribe(sc)
	if err != nil {
		panic("Subscribe method fail:")
	}

	return nil
}
