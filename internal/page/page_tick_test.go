package page

import (
	"context"
	"errors"
	"testing"
)

func TestTickRunsCallback(t *testing.T) {
	t.Parallel()

	var got context.Context

	ctx := context.Background()
	p := &Page{}
	p.SetTick(func(c context.Context) error {
		got = c

		return nil
	})

	if err := p.Tick(ctx); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	if got != ctx {
		t.Fatal("the callback did not receive the context")
	}
}

func TestTickWithoutCallback(t *testing.T) {
	t.Parallel()

	if err := (&Page{}).Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}
}

func TestTickNilRemovesCallback(t *testing.T) {
	t.Parallel()

	calls := 0
	p := &Page{}
	p.SetTick(func(context.Context) error {
		calls++

		return nil
	})
	p.SetTick(nil)

	if err := p.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	if calls != 0 {
		t.Fatalf("calls = %d", calls)
	}
}

func TestTickReturnsCallbackError(t *testing.T) {
	t.Parallel()

	want := errors.New("stop")
	p := &Page{}
	p.SetTick(func(context.Context) error { return want })

	if err := p.Tick(context.Background()); !errors.Is(err, want) {
		t.Fatalf("err = %v", err)
	}
}
