package web

import (
	"context"
	"log"
	"net/http"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

// pdfScreen is a screen that can render its page as PDF bytes. A screen
// without PDF gets no /pdf route.
type pdfScreen interface {
	PDF(ctx context.Context, opts page.PDFOptions) ([]byte, error)
}

func (s *server) pdf(w http.ResponseWriter, r *http.Request) {
	exporter := s.app.(pdfScreen)

	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := exporter.PDF(r.Context(), page.PDFOptions{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Cache-Control", "no-store")

	if _, err := w.Write(data); err != nil {
		log.Printf("write pdf: %v", err)
	}
}
