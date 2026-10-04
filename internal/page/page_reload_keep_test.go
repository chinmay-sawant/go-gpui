package page

import (
	"context"
	"testing"
)

func TestReloadKeepsTickAndImages(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	path := writeWatchFile(t, `<p id="p">one</p>`)
	p, err := New(Config{File: path, Width: 320, Height: 200})
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	ticks := 0
	p.SetTick(func(context.Context) error {
		ticks++

		return nil
	})
	p.SetImage("logo", []byte("bytes"))

	rewriteWatchFile(t, path, `<p id="p">two, longer</p>`)

	if _, err := p.PollReload(ctx); err != nil {
		t.Fatal(err)
	}

	if err := p.Tick(ctx); err != nil {
		t.Fatal(err)
	}

	if ticks != 1 {
		t.Fatalf("ticks = %d, want 1", ticks)
	}

	data, err := p.imageBytes("logo")
	if err != nil || string(data) != "bytes" {
		t.Fatalf("image = %q, %v", data, err)
	}
}
