package server

import (
	"net/http"
)

func NewServer(addr string) *Server {
	return &Server{
		server: &http.Server{Addr: addr},
		router: NewRouter(),
	}
}
