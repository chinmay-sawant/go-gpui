package page_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

// rewriteSource replaces a watched file and bumps the mtime when the stat
// did not change.
func rewriteSource(t *testing.T, path, body string) {
	t.Helper()

	before, err := os.Stat(path)
	if err != nil {
		before = nil
	}

	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}

	after, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	if before != nil && before.ModTime().Equal(after.ModTime()) && before.Size() == after.Size() {
		stamp := after.ModTime().Add(time.Second)
		if err := os.Chtimes(path, stamp, stamp); err != nil {
			t.Fatal(err)
		}
	}
}

// watchedPage opens a file-backed page and draws it once.
func watchedPage(t *testing.T, path string) *page.Page {
	t.Helper()

	p := newFilePage(t, page.Config{File: path, Width: 320, Height: 200})
	if err := p.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	return p
}

func TestPollReloadNoChangeDoesNotRedraw(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := watchedPage(t, writeSource(t, "index.html", `<p id="p">one</p>`))
	gen := p.Generation()

	for i := 0; i < 2; i++ {
		changed, err := p.PollReload(ctx)
		if changed || err != nil {
			t.Fatalf("poll = %v, %v", changed, err)
		}
	}

	if p.Generation() != gen {
		t.Fatalf("generation = %d, want %d", p.Generation(), gen)
	}
}

func TestPollReloadOneChangedByteRedrawsOnce(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	path := writeSource(t, "index.html", `<p id="p">one</p>`)
	p := watchedPage(t, path)
	gen := p.Generation()

	rewriteSource(t, path, `<p id="p">ono</p>`)

	changed, err := p.PollReload(ctx)
	if !changed || err != nil {
		t.Fatalf("poll = %v, %v", changed, err)
	}

	if p.HTML() != `<p id="p">ono</p>` {
		t.Fatalf("HTML = %q", p.HTML())
	}

	if p.Generation() != gen+1 {
		t.Fatalf("generation = %d, want one redraw", p.Generation()-gen)
	}
}
