package music

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync/atomic"
	"testing"
)

// ovServer serves an Openverse-shaped search plus two track bodies.
func ovServer(t *testing.T, downloads *atomic.Int32) *httptest.Server {
	t.Helper()

	var srv *httptest.Server

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/audio/", func(w http.ResponseWriter, r *http.Request) {
		// Openverse answers 401 when an anonymous request exceeds this.
		if size, _ := strconv.Atoi(r.URL.Query().Get("page_size")); size > 20 {
			w.WriteHeader(http.StatusUnauthorized)

			return
		}

		fmt.Fprintf(w, `{"results":[
			{"title":"First","creator":"Ann","license":"by-sa","license_version":"3.0","url":"%s/a.mp3","filetype":"mp32"},
			{"title":"Second","creator":"Bo","license":"cc0","license_version":"1.0","url":"%s/b.mp3","filetype":"mp3"},
			{"title":"Skipped","creator":"Cy","license":"by","license_version":"4.0","url":"%s/c.ogg","filetype":"ogg"}
		]}`, srv.URL, srv.URL, srv.URL)
	})
	mux.HandleFunc("/a.mp3", func(w http.ResponseWriter, _ *http.Request) {
		downloads.Add(1)
		fmt.Fprint(w, "ID3first-body")
	})
	mux.HandleFunc("/b.mp3", func(w http.ResponseWriter, _ *http.Request) {
		downloads.Add(1)
		fmt.Fprint(w, "ID3second-body")
	})

	srv = httptest.NewServer(mux)

	return srv
}
