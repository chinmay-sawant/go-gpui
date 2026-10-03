package player

import (
	"io"
	"net/http"
	"strings"
	"sync"
)

// testAPI serves canned iTunes JSON and tiny artwork bytes, and records the
// paths it answered so a test can prove the covers were fetched.
type testAPI struct {
	mu    sync.Mutex
	paths []string
}

func (s *testAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	s.paths = append(s.paths, r.URL.RequestURI())
	s.mu.Unlock()

	base := "http://" + r.Host

	switch {
	case r.URL.Path == "/search":
		switch r.URL.Query().Get("entity") {
		case "song":
			io.WriteString(w, songsBody(base))
		case "album":
			io.WriteString(w, albumsBody(base))
		default:
			http.NotFound(w, r)
		}
	case strings.Contains(r.URL.Path, "song-"):
		w.Write(tinyPNG())
	case strings.Contains(r.URL.Path, "album-"):
		w.Write(tinyJPEG())
	default:
		http.NotFound(w, r)
	}
}

// saw reports whether any answered path contains fragment.
func (s *testAPI) saw(fragment string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, path := range s.paths {
		if strings.Contains(path, fragment) {
			return true
		}
	}

	return false
}
