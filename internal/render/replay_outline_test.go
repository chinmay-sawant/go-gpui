package render_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/render"
)

// TestReplayAcceptsOutlines covers the two shapes the engine emits for a CSS
// outline: side lines in the outline layer, and one rounded stroke when the
// box has a border radius. Both used to fall back because the gate rejected
// every op with the outline flag.
func TestReplayAcceptsOutlines(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		source string
		focus  string
	}{
		{
			"focus",
			`<style>#f{width:200px;border:1px solid #ccc}` +
				`#f:focus-visible{outline:2px solid #1a56db}</style>` +
				`<input id="f">`,
			"f",
		},
		{
			"checked",
			`<style>#a{width:18px;height:18px}` +
				`#a:checked{outline:3px solid #176b45}</style>` +
				`<input id="a" type="checkbox" checked>`,
			"",
		},
		{
			"rounded",
			`<style>#r{width:80px;height:40px}` +
				`#r:focus{outline:2px solid #1a56db;border-radius:6px}</style>` +
				`<div id="r">round</div>`,
			"r",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			display, err := render.DisplayListState(
				context.Background(), tc.source, 320, 240, render.State{Focus: tc.focus})
			if err != nil {
				t.Fatal(err)
			}

			outlines := 0

			for i := range display.Ops {
				if display.Ops[i].Outline() {
					outlines++
				}
			}

			if outlines == 0 {
				t.Fatal("no outline op to replay")
			}

			if !render.Replayable(display) {
				t.Fatal("outline page should replay")
			}
		})
	}
}
