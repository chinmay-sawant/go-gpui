package web

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/host"
	"github.com/chinmay-sawant/ownframe/internal/page"
)

type fakePDF struct {
	host.Screen
	data []byte
}

func (f *fakePDF) PDF(context.Context, page.PDFOptions) ([]byte, error) {
	return f.data, nil
}

func TestPDFHandler(t *testing.T) {
	srv := &server{app: &fakePDF{data: []byte("%PDF-test")}}

	rec := httptest.NewRecorder()
	srv.pdf(rec, httptest.NewRequest("GET", "/pdf", nil))

	if got := rec.Header().Get("Content-Type"); got != "application/pdf" {
		t.Fatalf("Content-Type = %q", got)
	}

	if got := rec.Body.String(); got != "%PDF-test" {
		t.Fatalf("body = %q", got)
	}
}
