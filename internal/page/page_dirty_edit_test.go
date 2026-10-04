package page_test

import (
	"context"
	"testing"
)

func TestFocusTypingAndBlurDirtyField(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen := formPage(t, `<input id="t" type="text"`+wide+`><div id="pad" style="height:40px">pad</div>`)

	clickID(t, screen, "t")
	rect, ok := screen.TakeDirty()
	if !ok || !boxIn(rect, boxFor(t, screen, "t")) {
		t.Fatalf("focus rect %v ok %v", rect, ok)
	}

	if err := screen.Type(ctx, "abc"); err != nil {
		t.Fatal(err)
	}

	rect, ok = screen.TakeDirty()
	if !ok || !boxIn(rect, boxFor(t, screen, "t")) {
		t.Fatalf("typing rect %v ok %v", rect, ok)
	}

	clickID(t, screen, "pad")

	rect, ok = screen.TakeDirty()
	if !ok || !boxIn(rect, boxFor(t, screen, "t")) {
		t.Fatalf("blur rect %v ok %v", rect, ok)
	}
}
