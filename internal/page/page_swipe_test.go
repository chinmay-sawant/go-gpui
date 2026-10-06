package page

import (
	"context"
	"testing"
)

func TestSwipeHandlerGetsDelta(t *testing.T) {
	t.Parallel()

	p := bedPage(t, `<p>hi</p>`, nil)

	var gotX, gotY float64
	p.Handle(Handlers{Swipe: func(_ context.Context, dx, dy float64) error {
		gotX, gotY = dx, dy

		return nil
	}})

	if err := p.Swipe(context.Background(), 3, -40); err != nil {
		t.Fatal(err)
	}

	if gotX != 3 || gotY != -40 {
		t.Fatalf("swipe = %v, %v", gotX, gotY)
	}
}

func TestSwipeWithoutHandlerDrawsNothing(t *testing.T) {
	t.Parallel()

	p := bedPage(t, `<p>hi</p>`, nil)
	p.TakeDirty()
	before := p.Stats().Redraws

	if err := p.Swipe(context.Background(), 0, -60); err != nil {
		t.Fatal(err)
	}

	if after := p.Stats().Redraws; after != before {
		t.Fatalf("redraws = %d, want %d", after, before)
	}

	if _, ok := p.TakeDirty(); ok {
		t.Fatal("a swipe with no handler dirtied the page")
	}
}
