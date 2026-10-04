package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func TestSetHotReloadOff(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	path := writeSource(t, "index.html", `<p id="p">one</p>`)
	p := watchedPage(t, path)

	if !p.Watching() {
		t.Fatal("file page does not watch")
	}

	p.SetHotReload(false)
	if p.Watching() {
		t.Fatal("watch stayed on")
	}

	rewriteSource(t, path, `<p id="p">two</p>`)

	changed, err := p.PollReload(ctx)
	if changed || err != nil {
		t.Fatalf("off poll = %v, %v", changed, err)
	}

	p.SetHotReload(true)

	changed, err = p.PollReload(ctx)
	if !changed || err != nil {
		t.Fatalf("on poll = %v, %v", changed, err)
	}
}

func TestStringPageDoesNotWatch(t *testing.T) {
	t.Parallel()

	p := newFilePage(t, page.Config{HTML: `<p id="p">one</p>`, Width: 320, Height: 200})

	if p.Watching() {
		t.Fatal("string page watches")
	}

	changed, err := p.PollReload(context.Background())
	if changed || err != nil {
		t.Fatalf("poll = %v, %v", changed, err)
	}
}
