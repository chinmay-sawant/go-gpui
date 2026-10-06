package render_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/render"
)

// TestDisplaySkipsClippedOps pins the case that makes OpNoop necessary: a
// positioned child clipped by an overflow box leaves a deactivated operation
// behind rather than removing it, because the box tree stores operation
// indices that must not shift. A caller iterating a real page meets it.
func TestDisplaySkipsClippedOps(t *testing.T) {
	t.Parallel()

	display, err := render.DisplayList(
		context.Background(),
		`<div style="height:40px;overflow:hidden;position:relative">`+
			`<div style="position:absolute;top:200px;width:30px;height:30px;background:#0f0"></div>`+
			`<p>visible</p></div>`,
		320,
		200,
	)
	if err != nil {
		t.Fatal(err)
	}

	var noops, painted int

	for _, op := range display.Ops {
		if op.Kind == render.OpNoop {
			noops++
		} else {
			painted++
		}
	}

	if noops == 0 {
		t.Fatal("clipped positioned child left no deactivated operation behind")
	}

	if painted == 0 {
		t.Fatal("the visible content was deactivated too")
	}
}
