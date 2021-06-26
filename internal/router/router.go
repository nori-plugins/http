package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	nori_http "github.com/nori-io/interfaces/nori/http/v2"
)

type Router struct {
	router chi.Router
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt.router.ServeHTTP(w, r)
}

func (rt *Router) Handle(pattern string, h http.Handler) {
	rt.router.Handle(pattern, h)
}

func (rt *Router) HandleFunc(pattern string, h http.HandlerFunc) {
	rt.router.HandleFunc(pattern, h)
}

func (rt *Router) Method(method, pattern string, h http.Handler) {
	rt.router.Method(method, pattern, h)
}

func (rt *Router) MethodFunc(method, pattern string, h http.HandlerFunc) {
	rt.router.MethodFunc(method, pattern, h)
}

func (rt *Router) Connect(pattern string, h http.HandlerFunc) {
	rt.router.Connect(pattern, h)
}

func (rt *Router) Delete(pattern string, h http.HandlerFunc) {
	rt.router.Delete(pattern, h)
}

func (rt *Router) Get(pattern string, h http.HandlerFunc) {
	rt.router.Get(pattern, h)
}

func (rt *Router) Head(pattern string, h http.HandlerFunc) {
	rt.router.Head(pattern, h)
}

func (rt *Router) Options(pattern string, h http.HandlerFunc) {
	rt.router.Options(pattern, h)
}

func (rt *Router) Patch(pattern string, h http.HandlerFunc) {
	rt.router.Patch(pattern, h)
}

func (rt *Router) Post(pattern string, h http.HandlerFunc) {
	rt.router.Post(pattern, h)
}

func (rt *Router) Put(pattern string, h http.HandlerFunc) {
	rt.router.Put(pattern, h)
}

func (rt *Router) Trace(pattern string, h http.HandlerFunc) {
	rt.router.Trace(pattern, h)
}

func (rt *Router) Use(middlewares ...func(http.Handler) http.Handler) {
	rt.router.Use(middlewares...)
}

func (rt *Router) URLParam(r *http.Request, key string) string {
	name := chi.URLParam(r, key)
	return name
}

func (rt *Router) Route(pattern string, fn func(r nori_http.Router)) nori_http.Router {
	if fn == nil {
		panic(fmt.Sprintf("attempting to Route() a nil subrouter on '%s'", pattern))
	}
	subRouter := Router{router: chi.NewRouter()}
	fn(&subRouter)
	rt.router.Mount(pattern, &subRouter)
	return &subRouter

}
func (rt *Router) Mount(pattern string, h http.Handler) {
	rt.router.Mount(pattern, h)
}

func (rt *Router) NotFound(h http.HandlerFunc) {
	rt.router.NotFound(h)
}

func (rt *Router) MethodNotAllowed(h http.HandlerFunc) {
	rt.router.MethodNotAllowed(h)
}

func (rt *Router) With(middlewares ...func(http.Handler) http.Handler) nori_http.Router {
	mux := rt.router.With(middlewares...)
	return &Router{
		router: mux,
	}
}
