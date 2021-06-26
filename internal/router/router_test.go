package router_test

import (
	"context"
	"fmt"
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

	const (
		key1               = "key1"
		key1Value          = "1"
		key2               = "key2"
		key2Value          = "2"
		key1subrouter      = "key1subrouter"
		key1subrouterValue = "1subrouter"
		key2subrouter      = "key2subrouter"
		key2subrouterValue = "2subrouter"
	)

	middlewareAddKey1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, key1, key1Value)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	middlewareAddKey2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, key2, key2Value)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	middlewareAddKey1Subrouter := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, key1subrouter, key1subrouterValue)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	middlewareAddKey2Subrouter := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, key2subrouter, key2subrouterValue)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	//handler for route in testRequest(t, ts, http.MethodGet, "/profiles", nil)
	getAllProfiles := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxValue1 := r.Context().Value(key1)
		if ctxValue1 == nil {
			t.Fatal("ctxValue1=nil")
		}
		assert.Equal(t, key1Value, fmt.Sprint(ctxValue1))

		ctxValue2 := r.Context().Value(key2)
		if ctxValue2 == nil {
			t.Fatal("ctxValue2=nil")
		}
		assert.Equal(t, key2Value, ctxValue2)

		ctxValue1Subrouter := r.Context().Value(key1subrouter)
		assert.Nil(t, ctxValue1Subrouter)

		ctxValue2Subrouter := r.Context().Value(key2subrouter)
		assert.Nil(t, ctxValue2Subrouter)
	})

	//handler for route in testRequest(t, ts, http.MethodGet, "/profiles/1", nil)
	getOneProfile := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxValue1 := r.Context().Value(key1)
		if ctxValue1 == nil {
			t.Fatal("ctxValue1=nil")
		}
		assert.Equal(t, key1Value, fmt.Sprint(ctxValue1))

		ctxValue2 := r.Context().Value(key2)
		assert.Nil(t, ctxValue2)

		ctxValue1Subrouter := r.Context().Value(key1subrouter)
		if ctxValue1Subrouter == nil {
			t.Fatal("ctxValue1Subrouter=nil")
		}
		assert.Equal(t, key1subrouterValue, fmt.Sprint(ctxValue1Subrouter))

		ctxValue2Subrouter := r.Context().Value(key2subrouter)
		if ctxValue2Subrouter == nil {
			t.Fatal("ctxValue1Subrouter=nil")
		}
		assert.Equal(t, key2subrouterValue, fmt.Sprint(ctxValue2Subrouter))
	})

	router.Route("/profiles", func(r http2.Router) {
		r.Use(middlewareAddKey1)
		r.With(middlewareAddKey2).Get("/", getAllProfiles)
		// Subrouters:
		r.Route("/{profileID}", func(r http2.Router) {
			r.Use(middlewareAddKey1Subrouter)
			r.With(middlewareAddKey2Subrouter).Get("/", getOneProfile) // GET /profiles/1
		})
	})

	ts := httptest.NewServer(router)
	defer ts.Close()

	testRequest(t, ts, http.MethodGet, "/profiles", nil)
	testRequest(t, ts, http.MethodGet, "/profiles/1", nil)
}
