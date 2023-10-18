package server

import (
	"context"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

func Subscribe(ctx context.Context, s stan.Conn, p processor) (err error) {
	sub, err := s.Subscribe("MessengerFoo", func(msg *stan.Msg) {
		err = p.Process(msg.Data)
		if err != nil {
			panic(err.Error())
		}
	}, stan.DeliverAllAvailable())
	if err != nil {
		log.Fatal("message not received:", err)
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("context done, exit...")
			err = sub.Unsubscribe()
			if err != nil {
				return err
			}
			return ctx.Err() // context canceled
		default:
			fmt.Println("123")
		}
	}
}
