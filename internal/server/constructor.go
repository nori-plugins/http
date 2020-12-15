package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

func New(addr string) *Server {

	return &Server{
		server: &http.Server{Addr: addr},
		router: chi.NewRouter(),
	}
}
