package insights

import (
	"context"
	"testing"
)

// TestShareAndDownloadClick checks the share badge and the mobile download
// carry actions and leave a note, so every card control responds.
func TestShareAndDownloadClick(t *testing.T) {
	app, err := New()
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	ctx := context.Background()

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	share, ok := findBox(app.Boxes(), "share")
	if !ok {
		t.Fatal("no share box")
	}

	if share.Action != "share" {
		t.Fatalf("share action = %q", share.Action)
	}

	clickBox(t, app, share)

	if app.View().Note != "Share link copied" {
		t.Fatalf("note after share = %q", app.View().Note)
	}

	download, ok := findBox(app.Boxes(), "download")
	if !ok {
		t.Fatal("no download box")
	}

	if download.Action != "download" {
		t.Fatalf("download action = %q", download.Action)
	}

	clickBox(t, app, download)

	if app.View().Note != "Download link sent to your phone" {
		t.Fatalf("note after download = %q", app.View().Note)
	}
}
