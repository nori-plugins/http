package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v4"
)

type Server struct {
	router chi.Router
	server *http.Server
}

// todo: implement start and graceful shutdown methods
// todo: implement wrapper around go-chi methods. map http interface to go-chi router interface

func (s *Server) Start(port string) error {
	s.server.Addr = port
	s.server.Handler = s.router
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
