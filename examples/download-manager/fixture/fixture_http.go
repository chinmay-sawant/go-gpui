package fixture

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// serveOK sends a complete response with a length and a validator.
func (s *Server) serveOK(w http.ResponseWriter, r *http.Request) {
	body := s.fill(requestSeed(r))
	w.Header().Set("ETag", etagOK)
	w.Header().Set("Last-Modified", "Wed, 01 Oct 2025 12:00:00 GMT")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	_, _ = w.Write(body)
}

// serveSlow trickles the body in small chunks.
func (s *Server) serveSlow(w http.ResponseWriter, r *http.Request) {
	body := s.fill(requestSeed(r))
	w.Header().Set("ETag", `"fixture-slow"`)
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))

	for off := 0; off < len(body); off += 1024 {
		end := off + 1024
		if end > len(body) {
			end = len(body)
		}

		_, _ = w.Write(body[off:end])
		w.(http.Flusher).Flush()

		if s.SlowPause > 0 {
			time.Sleep(s.SlowPause)
		}
	}
}

// serveChunked hides the length by using chunked transfer encoding.
func (s *Server) serveChunked(w http.ResponseWriter, r *http.Request) {
	body := s.fill(requestSeed(r))
	w.Header().Set("ETag", `"fixture-chunked"`)

	for off := 0; off < len(body); off += 4096 {
		end := off + 4096
		if end > len(body) {
			end = len(body)
		}

		_, _ = w.Write(body[off:end])
		w.(http.Flusher).Flush()
	}
}

// serveRange answers Range and If-Range; a changed validator gives 200.
func (s *Server) serveRange(w http.ResponseWriter, r *http.Request) {
	body := s.fill(requestSeed(r))
	w.Header().Set("ETag", etagOK)
	w.Header().Set("Accept-Ranges", "bytes")

	ifRange := r.Header.Get("If-Range")
	if ifRange != "" && ifRange != etagOK {
		writeFull(w, body)

		return
	}

	start, end, ok := parseRange(r.Header.Get("Range"), int64(len(body)))
	if !ok {
		w.Header().Set("Content-Range", fmt.Sprintf("bytes */%d", len(body)))
		w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)

		return
	}

	w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, len(body)))
	w.Header().Set("Content-Length", strconv.FormatInt(end-start+1, 10))
	w.WriteHeader(http.StatusPartialContent)
	_, _ = w.Write(body[start : end+1])
}

// serveChanging bumps its ETag and body on every hit.
func (s *Server) serveChanging(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	s.round++
	round := s.round
	s.mu.Unlock()

	body := s.fill(round)
	w.Header().Set("ETag", fmt.Sprintf(`"changing-%d"`, round))
	writeFull(w, body)
}

// serveRedirect sends a 302 to the stable route.
func (s *Server) serveRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, PathOK, http.StatusFound)
}

// serveInterrupt writes half the announced body and drops the connection.
func (s *Server) serveInterrupt(w http.ResponseWriter, r *http.Request) {
	body := s.fill(requestSeed(r))
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	_, _ = w.Write(body[:len(body)/2])
	w.(http.Flusher).Flush()

	panic(http.ErrAbortHandler)
}

// serveStatus answers with a bare status code from ?code=.
func (s *Server) serveStatus(w http.ResponseWriter, r *http.Request) {
	code, err := strconv.Atoi(r.URL.Query().Get("code"))
	if err != nil || code < 100 || code > 599 {
		code = http.StatusTeapot
	}

	w.WriteHeader(code)
}

// writeFull sends the body with a length and a validator.
func writeFull(w http.ResponseWriter, body []byte) {
	w.Header().Set("ETag", etagOK)
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	_, _ = w.Write(body)
}

// parseRange turns "bytes=5-" or "bytes=5-9" into an inclusive span. It
// returns ok=false for anything unsatisfiable.
func parseRange(value string, size int64) (start, end int64, ok bool) {
	value = strings.TrimPrefix(value, "bytes=")
	if value == "" || size == 0 {
		return 0, 0, false
	}

	left, right, found := strings.Cut(value, "-")
	if !found {
		return 0, 0, false
	}

	start, err := strconv.ParseInt(left, 10, 64)
	if err != nil || start < 0 || start >= size {
		return 0, 0, false
	}

	end = size - 1
	if right != "" {
		end, err = strconv.ParseInt(right, 10, 64)
		if err != nil || end < start {
			return 0, 0, false
		}

		if end >= size {
			end = size - 1
		}
	}

	return start, end, true
}
