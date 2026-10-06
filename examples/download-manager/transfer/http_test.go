package transfer

import (
	"context"
	"errors"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/fixture"
)

// TestHTTPOK downloads the stable route and a redirect to it.
func TestHTTPOK(t *testing.T) {
	srv := fixture.NewServer()
	defer srv.Close()

	tr := newHTTP(t)
	body, out := fetchFull(t, tr, srv.URL+fixture.PathOK)

	if len(body) != 64<<10 {
		t.Fatalf("body %d bytes", len(body))
	}

	if out.Total != 64<<10 {
		t.Errorf("total %d", out.Total)
	}

	if out.Validators.ETag != `"fixture-v1"` {
		t.Errorf("etag %q", out.Validators.ETag)
	}

	again, _ := fetchFull(t, tr, srv.URL+fixture.PathOK)
	if string(again) != string(body) {
		t.Error("stable route changed between hits")
	}

	redirected, _ := fetchFull(t, tr, srv.URL+fixture.PathRedirect)
	if string(redirected) != string(body) {
		t.Error("redirect body differs")
	}
}

// TestHTTPStatusError maps an HTTP error to ErrBadStatus.
func TestHTTPStatusError(t *testing.T) {
	srv := fixture.NewServer()
	defer srv.Close()

	dest, partial := httpPaths(t)

	_, err := newHTTP(t).Download(context.Background(), Request{
		URL: srv.URL + "/status?code=503", Dest: dest, Partial: partial,
	}, nil)
	if !errors.Is(err, ErrBadStatus) {
		t.Fatalf("want ErrBadStatus, got %v", err)
	}
}
