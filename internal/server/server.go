package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router chi.Router
	server *http.Server
}

func (s *Server) Start() error {
	s.server.Handler = s.router
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) Handle(pattern string, h http.Handler) {
	s.router.Handle(pattern, h)
}

func (s *Server) HandleFunc(pattern string, h http.HandlerFunc) {
	s.router.HandleFunc(pattern, h)
}

func (s *Server) Method(method, pattern string, h http.Handler) {
	s.router.Method(method, pattern, h)
}

func (s *Server) MethodFunc(method, pattern string, h http.HandlerFunc) {
	s.router.MethodFunc(method, pattern, h)
}

func (s *Server) Connect(pattern string, h http.HandlerFunc) {
	s.router.Connect(pattern, h)
}

func (s *Server) Delete(pattern string, h http.HandlerFunc) {
	s.router.Delete(pattern, h)
}

func (s *Server) Get(pattern string, h http.HandlerFunc) {
	s.router.Get(pattern, h)
}

func (s *Server) Head(pattern string, h http.HandlerFunc) {
	s.router.Head(pattern, h)
}

func (s *Server) Options(pattern string, h http.HandlerFunc) {
	s.router.Options(pattern, h)
}

func (s *Server) Patch(pattern string, h http.HandlerFunc) {
	s.router.Patch(pattern, h)
}

func (s *Server) Post(pattern string, h http.HandlerFunc) {
	s.router.Post(pattern, h)
}

func (s *Server) Put(pattern string, h http.HandlerFunc) {
	s.router.Put(pattern, h)
}

func (s *Server) Trace(pattern string, h http.HandlerFunc) {
	s.router.Trace(pattern, h)
}

func (s *Server) Use(middlewares ...func(http.Handler) http.Handler) {
	s.router.Use(middlewares...)
}

func (s *Server) URLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}
