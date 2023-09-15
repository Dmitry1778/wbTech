package server

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

//type Subscriber struct {
//	sc *stan.Conn
//}

func Subscribe(s stan.Conn) (err error) {
	sub, err := s.Subscribe("MessengerFoo", func(msg *stan.Msg) {
		fmt.Printf("Message: %s\n", string(msg.Data))
	}, stan.DeliverAllAvailable())
	if err != nil {
		log.Fatal("message not received:", err)
	}
	defer sub.Unsubscribe()
	time.Sleep(5 * time.Second)
	return nil
}
