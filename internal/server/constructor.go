package server

import (
	"github.com/go-chi/chi/v4"
)

func New() *Server {
	return &Server{
		server: chi.NewRouter(),
	}
}
