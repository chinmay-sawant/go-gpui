package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/page"
)

// The changed text sits below the window frame. TakeDirty must keep its
// box, since the replay buffer covers the content past the canvas.
func TestDirtyRectBelowFrameIsKept(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen, err := page.New(page.Config{
		HTML:   `<body style="margin:0"><p id="t" style="margin-top:300px">{{.V}}</p></body>`,
		Width:  320,
		Height: 100,
	})
	if err != nil {
		t.Fatal(err)
	}

	screen.SetData(struct{ V string }{V: "a"})
	if err = screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}
	screen.TakeDirty()

	screen.SetData(struct{ V string }{V: "b"})
	if err = screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	rect, ok := screen.TakeDirty()
	if !ok || rect.Empty() {
		t.Fatalf("rect %v ok %v", rect, ok)
	}

	if rect.Min.Y <= 100 {
		t.Fatalf("rect %v collapsed to the frame", rect)
	}

	b := boxFor(t, screen, "t")
	if rect.Min.Y > int(b.Y+b.H+0.5) || rect.Max.Y < int(b.Y) {
		t.Fatalf("rect %v misses the changed text at %v", rect, b)
	}
}
