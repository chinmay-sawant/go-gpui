package player

import (
	"bytes"
	"context"
	"image/color"
	"image/png"
	"strings"
	"testing"
)

// TestLoadLiveFromServer serves canned JSON and artwork from one server.
func TestLoadLiveFromServer(t *testing.T) {
	srv, artwork := cannedServer(t)

	app, err := NewAt(srv.URL)
	if err != nil {
		t.Fatalf("NewAt: %v", err)
	}

	ctx := context.Background()
	if err := app.Load(ctx, "canned"); err != nil {
		t.Fatalf("Load: %v", err)
	}

	v := app.View()
	if len(v.Queue) != 2 || v.Queue[0].Title != "Canned One" || v.Queue[1].Title != "Canned Two" {
		t.Fatalf("Queue = %+v", v.Queue)
	}

	if !strings.Contains(v.Status, "Live") {
		t.Fatalf("Status = %q", v.Status)
	}

	if artwork.Load() != 2 {
		t.Fatalf("artwork requests = %d, want 2", artwork.Load())
	}

	if err := app.Redraw(ctx); err != nil {
		t.Fatalf("Redraw: %v", err)
	}

	assertReplayable(t, app)

	img, err := png.Decode(bytes.NewReader(app.PNG()))
	if err != nil {
		t.Fatalf("decode: %v", err)
	}

	if !hasColor(img, color.RGBA{R: 255, A: 255}, 40) {
		t.Fatal("no fetched red PNG pixels in the frame")
	}

	if !hasColor(img, color.RGBA{B: 255, A: 255}, 40) {
		t.Fatal("no fetched blue JPEG pixels in the frame")
	}
}
