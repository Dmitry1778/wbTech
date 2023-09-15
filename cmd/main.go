package main

import (
	"wbTech/internal/db"
	"wbTech/server"
)

func main() {

	//Building a project
	server.NatsConnectMethod()
	db.CreateNewDb()
}
