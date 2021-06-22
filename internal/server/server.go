package server

import (
	"context"
	"net/http"
)

type Server struct {
	router *Router
	server *http.Server
}

func (s *Server) Start() error {
	s.server.Handler = s.router
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
