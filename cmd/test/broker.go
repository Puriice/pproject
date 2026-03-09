package main

import (
	"log"

	"github.com/puriice/golibs/pkg/env"
	"github.com/puriice/golibs/pkg/messaging"
	"github.com/puriice/pProject/pkg/model"
	"github.com/puriice/pProject/pkg/sdk"
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
	broker, err := messaging.NewRabbitMQ(env.Get("amqp_url", "amqp://guest:guest@localhost/"), sdk.ExchangeName)

	if err != nil {
		log.Fatal(err)
	}

	projectService := sdk.NewService("", broker)

	listener, err := projectService.NewListener("projects.test")

	if err != nil {
		log.Fatal(err)
	}

	listener.OnCreate(onCreate)
	listener.OnUpdate(onUpdate)
	listener.OnDelete(onDelete)
	listener.OnError(onError)

	forever := make(chan struct{})

	listener.Subscribe()

	<-forever
}
