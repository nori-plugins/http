package router_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nori-plugins/http/internal/router"
)

func TestRouter_With_Use(t *testing.T) {
	router := router.NewRouter()

	with := router.With(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	})

	with.Get("/with_middleware", func(w http.ResponseWriter, r *http.Request) {})

	ts := httptest.NewServer(with)
	defer ts.Close()

	// Without the fix this test was committed with, this causes a panic.
	testRequest(t, ts, http.MethodGet, "/with_middleware", nil)

}

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
