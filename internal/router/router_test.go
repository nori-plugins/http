package router_test

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	http2 "github.com/nori-io/interfaces/nori/http/v2"

	"github.com/nori-plugins/http/internal/router"
)

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer resp.Body.Close()

	return resp, string(respBody)
}

func TestRouter_Use_With(t *testing.T) {
	router := router.NewRouter()

	addKey1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "key1", "1")
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	addKey2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "key2", "2")
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	addKey1Subrouter := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "key1subrouter", "1subrouter")
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	addKey2Subrouter := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "key2subrouter", "2subrouter")
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	get := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxValue1 := r.Context().Value("key1").(string)
		assert.Equal(t, "1", ctxValue1)
		ctxValue2 := r.Context().Value("key2").(string)
		assert.Equal(t, "2", ctxValue2)

		w.Write([]byte(ctxValue1))
	})
	getSub := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxValue1 := r.Context().Value("key1").(string)
		assert.Equal(t, "1", ctxValue1)
		ctxValue1Subrouter := r.Context().Value("key1subrouter").(string)
		assert.Equal(t, "1subrouter", ctxValue1Subrouter)
		ctxValue2Subrouter := r.Context().Value("key2subrouter").(string)
		assert.Equal(t, "2subrouter", ctxValue2Subrouter)
	})

	router.Route("/articles", func(r http2.Router) {
		r.Use(addKey1)
		r.With(addKey2).Get("/", get)
		// Subrouters:
		r.Route("/{articleID}", func(r http2.Router) {
			r.Use(addKey1Subrouter)
			r.With(addKey2Subrouter).Get("/", getSub) // GET /articles/123
		})
	})

	ts := httptest.NewServer(router)
	defer ts.Close()

	// Without the fix this test was committed with, this causes a panic.
	testRequest(t, ts, http.MethodGet, "/articles", nil)
	testRequest(t, ts, http.MethodGet, "/articles/1", nil)

}
