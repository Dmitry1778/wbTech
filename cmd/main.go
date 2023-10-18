package main

import (
	"context"
	"os"
	"os/signal"
	"time"
	"wbTech/api"
	"wbTech/cmd/config"
	"wbTech/internal/db"
)

func main() {

	cfg := config.GetConfig(&config.Config{})
	dbConn, err := db.PostgresConn(context.Background(), 3, cfg.Storage)
	if err != nil {
		panic(err)
	}
	database := db.NewDB(dbConn)

	cache := db.NewInit(database)

	//worker := new_order_subsriber.NewWorker(database, cache)

	ctx, cancel := context.WithCancel(context.Background())

	//go func() {
	//	err = server.NatsConnectMethod(ctx, worker)
	//	if err != nil {
	//		panic(err)
	//	}
	//}()
	go func() {
		api.NewApi(ctx, cache)
	}()

	osExit := make(chan os.Signal, 1)
	exitOnSignal(osExit)
	<-osExit
	cancel()
	time.Sleep(10 * time.Second)

}
func exitOnSignal(osExit chan os.Signal) {
	signal.Notify(osExit, os.Interrupt, os.Kill)
}
