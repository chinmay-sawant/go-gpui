package page_test

import (
	"context"
	"testing"
)

func TestOneClickOneLayoutOneRect(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	n := 0
	screen := counterPage(t, &n)
	before := screen.Generation()

	x, y := boxCenter(t, screen, "inc")
	if err := screen.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}

	if got := screen.Generation(); got != before+1 {
		t.Fatalf("generation %d, want %d", got, before+1)
	}

	rect, ok := screen.TakeDirty()
	if !ok || rect.Empty() {
		t.Fatalf("rect %v ok %v", rect, ok)
	}

	if _, again := screen.TakeDirty(); again {
		t.Fatal("dirty survived the first take")
	}
}
