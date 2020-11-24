package server

import (
	"net/http"

	"github.com/go-chi/chi/v4"
)

func New() *Server {
	return &Server{
		server: &http.Server{},
		router: chi.NewRouter(),
	}
}
