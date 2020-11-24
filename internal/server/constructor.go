package server

import (
	"net/http"

	"github.com/go-chi/chi/v4"
)

func New(port string) *Server {

	return &Server{
		server: &http.Server{Addr: port},
		router: chi.NewRouter(),
	}
}
