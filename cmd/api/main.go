package main

import (
	"log"

	"github.com/puriice/golibs/pkg/db"
	"github.com/puriice/golibs/pkg/env"
	"github.com/puriice/golibs/pkg/messaging"
	"github.com/puriice/golibs/pkg/server"
	"github.com/puriice/pproject/pkg/routing"
	"github.com/puriice/pproject/pkg/sdk/pproject"
)

func main() {
	env.Init()

	host := env.Get("HOST", "localhost")
	port := env.Get("PORT", "8080")
	database, err := db.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	serv := server.NewServer(host, port, database)
	rabbit, err := messaging.NewRabbitMQ(env.Get("amqp_url", "amqp://guest:guest@localhost/"))

	if err != nil {
		log.Fatal(err)
	}

	broker, err := rabbit.Broker(pproject.ExchangeName)

	if err != nil {
		log.Fatal(err)
	}

	routing.Register(serv, broker)

	serv.Start()
}
