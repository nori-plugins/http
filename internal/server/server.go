package server

import "github.com/go-chi/chi/v4"

type Server struct {
	server chi.Router
}

// todo: implement start and graceful shutdown methods
// todo: implement wrapper around go-chi methods. map http interface to go-chi router interface

func (s *Server) Start(port uint) error {
	// todo: start http Server
	panic("not implemented")
}

func (s *Server) Shutdown() error {
	// todo: stop http Server
	panic("not implemented")
}