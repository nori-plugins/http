package server

import (
	"context"
	"net/http"

	"github.com/nori-plugins/http/internal/router"
)

type Server struct {
	router *router.Router
	server *http.Server
}

func (s *Server) Start() error {
	s.server.Handler = s.router
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) Router() *router.Router {
	return s.router
}
