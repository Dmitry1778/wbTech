package server

import (
	"context"
	"github.com/nats-io/stan.go"
	"log"
)

type processor interface {
	Process([]byte) error
}

func NatsConnectMethod(ctx context.Context, processor processor) (err error) {
	sc, err := stan.Connect("test-cluster", "test-id", stan.NatsURL(":4444"))
	if err != nil {
		log.Fatal("fail to connect:", err)
	}
	defer sc.Close()

	Publisher(sc)

	return Subscribe(ctx, sc, processor)
}
