package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

// TestClickReachesControlThroughChild checks a click that lands on an
// anonymous child falls back to the innermost id-bearing element, so a
// child icon does not swallow the control's action.
func TestClickReachesControlThroughChild(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen, err := page.New(page.Config{
		HTML: `<body style="margin:0"><div id="go" data-action="go" style="width:100px;height:40px">` +
			`<div style="width:20px;height:20px">hit</div></div></body>`,
		Width:  320,
		Height: 400,
	})
	if err != nil {
		t.Fatal(err)
	}

	got := ""
	screen.Handle(page.Handlers{
		Click: func(_ context.Context, box page.Box) error {
			got = box.ID + ":" + box.Action

			return nil
		},
	})

	if err := screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	child := page.Box{}

	for _, box := range screen.Boxes() {
		if box.Tag == "div" && box.ID == "" && box.W > 0 {
			child = box
		}
	}

	if child.W <= 0 {
		t.Fatal("no child div box")
	}

	if err := screen.Click(ctx, child.X+child.W/2, child.Y+child.H/2); err != nil {
		t.Fatal(err)
	}

	if got != "go:go" {
		t.Fatalf("click reached %q, want the control go:go", got)
	}
}
