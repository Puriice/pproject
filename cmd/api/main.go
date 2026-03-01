package main

import (
	"log"

	"github.com/puriice/httplibs/pkg/db"
	"github.com/puriice/httplibs/pkg/server"
	"github.com/puriice/pProject/internal/env"
	"github.com/puriice/pProject/pkg/project"
)

func main() {
	env.InitEnv()

	host := env.GetEnv("HOST", "localhost")
	port := env.GetEnv("PORT", "8080")
	database, err := db.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	serv := server.NewServer(host, port, database)

	project.Register(serv)

	serv.Start()
}
