package render_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/render"
)

// TestDisplayCarriesTextRuns checks that a text run arrives with the font a
// replay engine needs: a handle whose bytes can be handed to an independent
// shaper, and a positive size to scale it by.
func TestDisplayCarriesTextRuns(t *testing.T) {
	t.Parallel()

	display, err := render.DisplayList(
		context.Background(),
		`<div style="background:#eee;width:120px;height:40px"></div><p>text run</p>`,
		320,
		200,
	)
	if err != nil {
		t.Fatal(err)
	}

	var seen int

	for _, op := range display.Ops {
		if op.Kind != render.OpText {
			continue
		}

		seen++

		if op.Font == nil {
			t.Fatal("text op without a font")
		}

		if len(op.Font.Bytes()) == 0 {
			t.Error("text op with no face bytes")
		}

		if op.Size <= 0 {
			t.Errorf("text size %v", op.Size)
		}
	}

	if seen == 0 {
		t.Fatal("no text ops")
	}
}

// TestDisplayCarriesRectGeometry checks that a fill arrives with a size, which
// is all a vector fill needs beyond the color it already carries.
func TestDisplayCarriesRectGeometry(t *testing.T) {
	t.Parallel()

	display, err := render.DisplayList(
		context.Background(),
		`<div style="background:#336699;width:120px;height:40px"></div>`,
		320,
		200,
	)
	if err != nil {
		t.Fatal(err)
	}

	for _, op := range display.Ops {
		if op.Kind != render.OpFillRect {
			continue
		}

		if op.W <= 0 || op.H <= 0 {
			t.Errorf("fill %vx%v", op.W, op.H)
		}

		return
	}

	t.Fatal("no fill op")
}
