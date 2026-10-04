package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func TestReloadKeepsHistory(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	path := writeSource(t, "index.html", `<p id="p">one</p>`)
	p := watchedPage(t, path)

	if err := p.Load(ctx, `<p id="p">two</p>`); err != nil {
		t.Fatal(err)
	}

	rewriteSource(t, path, `<p id="p">three</p>`)

	changed, err := p.PollReload(ctx)
	if !changed || err != nil {
		t.Fatalf("poll = %v, %v", changed, err)
	}

	if got := p.HTML(); got != `<p id="p">three</p>` {
		t.Fatalf("html = %q", got)
	}

	if err := p.Back(ctx); err != nil {
		t.Fatal(err)
	}

	if got := p.HTML(); got != `<p id="p">one</p>` {
		t.Fatalf("back = %q", got)
	}
}

func TestDisableHotReloadConfig(t *testing.T) {
	t.Parallel()

	path := writeSource(t, "index.html", `<p>one</p>`)
	p := newFilePage(t, page.Config{File: path, Width: 320, Height: 200, DisableHotReload: true})

	if p.Watching() {
		t.Fatal("watch is on with DisableHotReload")
	}
}

func TestReloadStatsCount(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	path := writeSource(t, "index.html", `<p id="p">one</p>`)
	p := watchedPage(t, path)

	if got := p.Stats().Reloads; got != 0 {
		t.Fatalf("reloads = %d, want 0", got)
	}

	rewriteSource(t, path, `<p id="p">two</p>`)

	if changed, err := p.PollReload(ctx); !changed || err != nil {
		t.Fatalf("poll = %v, %v", changed, err)
	}

	st := p.Stats()
	if st.Reloads != 1 || st.LastReloadError != "" {
		t.Fatalf("stats = %d reloads, error %q", st.Reloads, st.LastReloadError)
	}

	rewriteSource(t, path, `{{`)

	if _, err := p.PollReload(ctx); err == nil {
		t.Fatal("a broken template did not report an error")
	}

	if p.Stats().LastReloadError == "" {
		t.Fatal("reload error not recorded")
	}
}
