package fetch_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/fetch"
)

func TestCanceledContext(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "late")
	}))
	t.Cleanup(srv.Close)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := fetch.Do(ctx, http.MethodGet, srv.URL, nil, nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v", err)
	}
}

func TestNilContext(t *testing.T) {
	t.Parallel()

	_, err := fetch.Do(nil, http.MethodGet, "https://example.com", nil, nil)
	if err == nil || !strings.Contains(err.Error(), "nil context") {
		t.Fatalf("err = %v", err)
	}
}
