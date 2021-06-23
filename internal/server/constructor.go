package server

import (
	"net/http"

	"github.com/nori-plugins/http/internal/router"
)

func NewServer(addr string) *Server {
	return &Server{
		server: &http.Server{Addr: addr},
		router: router.NewRouter(),
	}
}
