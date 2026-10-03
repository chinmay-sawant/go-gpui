package player

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestLoadOfflineKeepsSample pins the degradation path.
func TestLoadOfflineKeepsSample(t *testing.T) {
	srv := httptest.NewServer(http.NotFoundHandler())
	srv.Close()

	app, err := NewAt(srv.URL)
	if err != nil {
		t.Fatalf("NewAt: %v", err)
	}

	if err := app.Load(context.Background(), "nothing"); err == nil {
		t.Fatal("Load succeeded against a closed server")
	}

	v := app.View()
	if len(v.Queue) != 6 || v.Status != offlineStatus || v.Query != "nothing" {
		t.Fatalf("Queue = %d, Status = %q, Query = %q", len(v.Queue), v.Status, v.Query)
	}

	if err := app.Redraw(context.Background()); err != nil {
		t.Fatalf("Redraw: %v", err)
	}

	if app.Page().Display() == nil || len(app.PNG()) == 0 {
		t.Fatal("the degraded frame is not replayable")
	}
}
