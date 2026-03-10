package routing

import (
	"net/http"

	"github.com/puriice/golibs/pkg/messaging"
	"github.com/puriice/golibs/pkg/middleware"
	"github.com/puriice/golibs/pkg/middleware/cors"
	"github.com/puriice/golibs/pkg/server"
	"github.com/puriice/pproject/internal/hander/project"
	"github.com/puriice/pproject/internal/repository"
)

func Register(s *server.Server, broker *messaging.RabbitBroker) {
	router := http.NewServeMux()

	projectModel := repository.NewPostgresProjectRepository(s.Database)
	projectHandler := project.NewHandler(projectModel, broker)

	projectHandler.RegisterRoute(router)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	corsOption := cors.NewCorsOptions()
	corsOption.AllowOrigins = cors.Wildcard()
	corsOption.AllowNoOrigin = true
	corsOption.AllowCredentials = true

	pipeline := middleware.Pipe(
		middleware.Logger,
		middleware.Cors(*corsOption),
	)

	s.Handler = pipeline(mux)
}
