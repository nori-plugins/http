package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

func New(port string) *Server {

	return &Server{
		server: &http.Server{Addr: port},
		router: chi.NewRouter(),
	}
}
