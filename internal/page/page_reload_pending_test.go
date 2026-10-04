package page_test

import (
	"context"
	"testing"
)

func TestPollReloadSecondEditWins(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	path := writeSource(t, "index.html", `<p id="p">one</p>`)
	p := watchedPage(t, path)
	gen := p.Generation()

	rewriteSource(t, path, `<p id="p">two</p>`)
	rewriteSource(t, path, `<p id="p">three, longer</p>`)

	changed, err := p.PollReload(ctx)
	if !changed || err != nil {
		t.Fatalf("poll = %v, %v", changed, err)
	}

	if p.HTML() != `<p id="p">three, longer</p>` {
		t.Fatalf("HTML = %q", p.HTML())
	}

	if p.Generation() != gen+1 {
		t.Fatalf("generation moved by %d, want one redraw", p.Generation()-gen)
	}
}

func TestPollReloadPendingRetry(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	good := `<p id="p">good</p>`
	path := writeSource(t, "index.html", good)
	p := watchedPage(t, path)
	gen := p.Generation()

	rewriteSource(t, path, `<p id="p">{{.Broken</p>`)

	changed, err := p.PollReload(ctx)
	if changed || err == nil {
		t.Fatalf("bad poll = %v, %v, want a parse error", changed, err)
	}

	if p.HTML() != good {
		t.Fatalf("HTML = %q, want the last good source", p.HTML())
	}

	if p.Generation() != gen {
		t.Fatal("a parse error redrew the page")
	}

	changed, err = p.PollReload(ctx)
	if changed || err == nil {
		t.Fatalf("retry poll = %v, %v, want the same parse error", changed, err)
	}

	rewriteSource(t, path, `<p id="p">fixed</p>`)

	changed, err = p.PollReload(ctx)
	if !changed || err != nil {
		t.Fatalf("fixed poll = %v, %v", changed, err)
	}

	if p.HTML() != `<p id="p">fixed</p>` {
		t.Fatalf("HTML = %q", p.HTML())
	}

	if p.Generation() != gen+1 {
		t.Fatal("the fix did not redraw exactly once")
	}
}
