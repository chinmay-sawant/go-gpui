package page

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestCaretBlinkHidesAndShows(t *testing.T) {
	ctx := context.Background()
	p, err := New(Config{HTML: `<input id="e" type="text" value="ab">`, Width: 300, Height: 100})
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if err := p.Focus(ctx, "e"); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(p.source, `data-gpui-caret="1"`) {
		t.Fatalf("caret missing after focus: %s", p.source)
	}

	p.form.blinkAt = time.Now().Add(-caretBlinkInterval)
	if err := p.Tick(ctx); err != nil {
		t.Fatal(err)
	}

	if strings.Contains(p.source, `data-gpui-caret="1"`) {
		t.Fatalf("caret still painted after a blink: %s", p.source)
	}

	if err := p.Tick(ctx); err != nil {
		t.Fatal(err)
	}

	if strings.Contains(p.source, `data-gpui-caret="1"`) {
		t.Fatalf("caret returned before the interval: %s", p.source)
	}

	p.form.blinkAt = time.Now().Add(-caretBlinkInterval)
	if err := p.Tick(ctx); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(p.source, `data-gpui-caret="1"`) {
		t.Fatalf("caret did not return: %s", p.source)
	}
}

func TestEditResetsBlink(t *testing.T) {
	ctx := context.Background()
	p, err := New(Config{HTML: `<input id="e" type="text" value="ab">`, Width: 300, Height: 100})
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if err := p.Focus(ctx, "e"); err != nil {
		t.Fatal(err)
	}

	p.form.blinkAt = time.Now().Add(-caretBlinkInterval)
	if err := p.Tick(ctx); err != nil {
		t.Fatal(err)
	}

	if p.form.blinkOn {
		t.Fatal("blink did not turn off")
	}

	if err := p.Type(ctx, "c"); err != nil {
		t.Fatal(err)
	}

	if !p.form.blinkOn {
		t.Fatal("typing did not reset the blink")
	}
}
