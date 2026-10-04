package page_test

import (
	"context"
	"os"
	"testing"
)

func TestPollReloadDeletedFileKeepsPicture(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	path := writeSource(t, "index.html", `<p id="p">one</p>`)
	p := watchedPage(t, path)
	gen := p.Generation()

	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}

	changed, err := p.PollReload(ctx)
	if changed || err == nil {
		t.Fatalf("first poll = %v, %v, want a stat error", changed, err)
	}

	if p.Generation() != gen || p.Display() == nil {
		t.Fatal("the last good picture was dropped")
	}

	changed, err = p.PollReload(ctx)
	if changed || err != nil {
		t.Fatalf("second poll = %v, %v, want the error reported once", changed, err)
	}

	rewriteSource(t, path, `<p id="p">back, longer</p>`)

	changed, err = p.PollReload(ctx)
	if !changed || err != nil {
		t.Fatalf("recreated poll = %v, %v", changed, err)
	}

	if p.HTML() != `<p id="p">back, longer</p>` {
		t.Fatalf("HTML = %q", p.HTML())
	}
}
