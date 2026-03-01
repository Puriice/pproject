package project

import (
	"net/http"

	"github.com/puriice/httplibs/pkg/middleware"
	"github.com/puriice/httplibs/pkg/middleware/cors"
	"github.com/puriice/httplibs/pkg/server"
	"github.com/puriice/pProject/internal/hander/project"
	"github.com/puriice/pProject/internal/repository/postgres"
)

func Register(s *server.Server) {
	router := http.NewServeMux()

	projectModel := postgres.NewRepository(s.Database)
	projectHandler := project.NewHandler(projectModel)

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
