package router

import "github.com/go-chi/chi/v5"

func NewRouter() *Router {
	return &Router{
		router: chi.NewRouter(),
	}
}
