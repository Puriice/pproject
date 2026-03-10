package main

import (
	"context"
	"log"

	"github.com/puriice/golibs/pkg/env"
	"github.com/puriice/golibs/pkg/messaging"
	"github.com/puriice/pproject/pkg/model"
	"github.com/puriice/pproject/pkg/sdk/pproject"
)

func onCreate(project *model.Project) {
	log.Printf("Project Created: id=%s", *project.ID)
}

func onUpdate(project *model.Project) {
	log.Printf("Project Updated: id=%s", *project.ID)
}

func onDelete(id string) {
	log.Printf("Project Deleted:id=%s", id)
}

func onError(err error) {
	log.Print(err)
}

func main() {
	rabbit, err := messaging.NewRabbitMQ(env.Get("amqp_url", "amqp://guest:guest@localhost/"))

	if err != nil {
		log.Fatal(err)
	}

	broker, err := rabbit.Broker(pproject.ExchangeName)

	if err != nil {
		log.Fatal(err)
	}

	projectService := pproject.NewService("", broker)

	listener, err := projectService.NewListener("projects.test")

	if err != nil {
		log.Fatal(err)
	}

	listener.OnCreate(onCreate)
	listener.OnUpdate(onUpdate)
	listener.OnDelete(onDelete)
	listener.OnError(onError)

	forever := make(chan struct{})

	listener.Subscribe(context.Background())

	<-forever
}
