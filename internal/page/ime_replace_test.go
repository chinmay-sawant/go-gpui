package page

import (
	"context"
	"testing"
)

func TestIMEReplaceInsertsAndReplaces(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := imeField(t, "")

	if err := p.IMEReplace(ctx, 0, 0, "ab", 2); err != nil {
		t.Fatal(err)
	}

	_, caret, _, _, _ := p.IMEContext()
	if caret != 2 {
		t.Fatalf("caret after insert = %d", caret)
	}

	if err := p.IMEReplace(ctx, 0, 1, "X", 1); err != nil {
		t.Fatal(err)
	}

	if got := p.FormValue("e"); got != "Xb" {
		t.Fatalf("value = %q", got)
	}
}

func TestIMEReplaceSnapsToRunes(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := imeField(t, "héllo")

	// Byte 2 sits inside the two-byte é, so it snaps back to its start.
	if err := p.IMEReplace(ctx, 2, 3, "", 0); err != nil {
		t.Fatal(err)
	}

	if got := p.FormValue("e"); got != "hllo" {
		t.Fatalf("value = %q", got)
	}

	// An empty replacement of an empty range deletes nothing.
	if err := p.IMEReplace(ctx, 3, 3, "", 0); err != nil {
		t.Fatal(err)
	}

	if got := p.FormValue("e"); got != "hllo" {
		t.Fatalf("value after no-op = %q", got)
	}
}
