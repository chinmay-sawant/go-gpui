package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func TestRemovedOpDirtiesOldBounds(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen, err := page.New(page.Config{
		HTML:   `<body style="margin:0">{{if .Show}}<div id="gone" style="width:60px;height:30px;background:#f00">x</div>{{end}}<p id="stay">stay</p></body>`,
		Width:  320,
		Height: 200,
	})
	if err != nil {
		t.Fatal(err)
	}

	screen.SetData(struct{ Show bool }{true})
	if err = screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}
	screen.TakeDirty()

	gone := boxFor(t, screen, "gone")
	screen.SetData(struct{ Show bool }{false})

	if err = screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	rect, ok := screen.TakeDirty()
	if !ok {
		t.Fatal("no dirty rect")
	}

	if !boxIn(rect, gone) {
		t.Fatalf("rect %v does not cover the removed box", rect)
	}
}
