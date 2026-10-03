package layout_test

import (
	"context"
	"strings"
	"testing"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/layout/layout"
)

func TestClickReportsBoxGeometry(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	if len(app.Boxes()) == 0 {
		t.Fatal("no hit-test boxes")
	}

	alpha, ok := boxByID(app.Boxes(), "alpha")
	if !ok {
		t.Fatal("no box id=alpha")
	}

	if alpha.W <= 0 || alpha.H <= 0 {
		t.Fatalf("alpha is %.0fx%.0f, want a non-zero box", alpha.W, alpha.H)
	}

	if alpha.Tag == "" {
		t.Fatal("alpha has no tag")
	}

	if err := app.Click(ctx, alpha.X+alpha.W/2, alpha.Y+alpha.H/2); err != nil {
		t.Fatal(err)
	}

	status := app.View().Status
	if !strings.Contains(status, "alpha") {
		t.Fatalf("status = %q, want the clicked id", status)
	}

	if !strings.Contains(status, alpha.Tag) {
		t.Fatalf("status = %q, want tag %q", status, alpha.Tag)
	}

	if data := app.PNG(); len(data) == 0 {
		t.Fatal("no PNG")
	}
}

func newApp(t *testing.T, ctx context.Context) *layout.App {
	t.Helper()

	app, err := layout.New()
	if err != nil {
		t.Fatal(err)
	}

	app.SetSize(layout.DefaultWidth, layout.DefaultHeight)

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	return app
}

func boxByID(boxes []gpui.Box, id string) (gpui.Box, bool) {
	for _, b := range boxes {
		if b.ID == id {
			return b, true
		}
	}

	return gpui.Box{}, false
}
