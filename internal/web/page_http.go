package web

import (
	"bytes"
	"log"
	"net/http"
)

func (s *server) page(w http.ResponseWriter, _ *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data := shellData{
		Areas: areas(s.app),
	}

	if display := s.app.Display(); display != nil {
		data.Width, data.Height = display.Width, display.Height
	} else if img := s.app.Image(); img != nil {
		b := img.Bounds()
		data.Width = b.Dx()
		data.Height = b.Dy()
	}

	var buf bytes.Buffer
	if err := s.shell.Execute(&buf, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Printf("write page: %v", err)
	}
}

func (s *server) frame(w http.ResponseWriter, _ *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	pngBytes := s.app.PNG()
	if len(pngBytes) == 0 {
		http.Error(w, "no frame", http.StatusNotFound)

		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "no-store")

	if _, err := w.Write(pngBytes); err != nil {
		log.Printf("write png: %v", err)
	}
}
