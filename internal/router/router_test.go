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

	hArticlesList := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxValue1 := r.Context().Value("key1").(string)
		assert.Equal(t, "1", ctxValue1)
		ctxValue2 := r.Context().Value("key2").(string)
		assert.Equal(t, "2", ctxValue2)

		w.Write([]byte(ctxValue1))
	})
	addKey2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "key2", "2")
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	router.Route("/articles", func(r http2.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := context.WithValue(r.Context(), "key1", "1")
				next.ServeHTTP(w, r.WithContext(ctx))
			})
		})
		r.With(addKey2).Get("/", hArticlesList)
	})

	ts := httptest.NewServer(router)
	defer ts.Close()

	// Without the fix this test was committed with, this causes a panic.
	testRequest(t, ts, http.MethodGet, "/articles", nil)

}
