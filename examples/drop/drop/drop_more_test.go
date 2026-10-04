package drop_test

import (
	"context"
	"io/fs"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/chinmay-sawant/go-gpui/examples/drop/drop"
)

func TestDroppedTextShowsFirstLines(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	fsys := fstest.MapFS{
		"notes.txt": {Data: []byte("first line\r\nsecond line\r\nthird line")},
	}

	if err := app.Drop(ctx, dropsFrom(t, fsys)); err != nil {
		t.Fatal(err)
	}

	if got := app.View().Status; !strings.Contains(got, "notes.txt") {
		t.Fatalf("Status = %q", got)
	}

	text := boxText(t, app, "text")
	if !strings.Contains(text, "first line") || !strings.Contains(text, "second line") {
		t.Fatalf("text = %q", text)
	}

	if strings.Contains(text, "\r") {
		t.Fatalf("text keeps a carriage return: %q", text)
	}
}

func TestDroppedDirectoryStaysOneEntry(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	fsys := fstest.MapFS{"docs": {Mode: fs.ModeDir | 0o755}}

	if err := app.Drop(ctx, dropsFrom(t, fsys)); err != nil {
		t.Fatal(err)
	}

	if got := app.View().Status; !strings.Contains(got, "docs: directory") {
		t.Fatalf("Status = %q", got)
	}
}

func TestDroppedOtherFileHasNoPreview(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	fsys := fstest.MapFS{"data.bin": {Data: []byte{0, 1, 2}}}

	if err := app.Drop(ctx, dropsFrom(t, fsys)); err != nil {
		t.Fatal(err)
	}

	if got := app.View().Status; !strings.Contains(got, "no preview") {
		t.Fatalf("Status = %q", got)
	}
}

func boxText(t *testing.T, app *drop.App, id string) string {
	t.Helper()

	for _, box := range app.Boxes() {
		if box.ID == id {
			return box.Text
		}
	}

	t.Fatalf("no box id=%q", id)

	return ""
}
