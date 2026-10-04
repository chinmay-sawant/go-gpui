package page

import (
	"context"
	"testing"
)

func TestContextMenuRows(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := bftPage(t, `<input id="t" type="text" value="ab"`+bedWide+`>`)
	bftDraw(t, p, nil)

	rows := p.ContextMenu()
	want := []string{"cut", "copy", "paste", "select-all", "undo", "redo"}
	if len(rows) != len(want) {
		t.Fatalf("rows %d", len(rows))
	}

	for i, id := range want {
		if rows[i].ID != id || rows[i].Label == "" {
			t.Fatalf("row %d = %+v", i, rows[i])
		}

		if rows[i].Enabled {
			t.Fatalf("row %s enabled with no focus", id)
		}
	}

	bftClick(t, p, "t")
	if err := p.Type(ctx, "c"); err != nil {
		t.Fatal(err)
	}

	rows = p.ContextMenu()
	if rows[0].Enabled || rows[1].Enabled || !rows[2].Enabled || !rows[3].Enabled {
		t.Fatalf("no selection rows %+v", rows)
	}

	if err := p.SelectAll(ctx); err != nil {
		t.Fatal(err)
	}

	rows = p.ContextMenu()
	if !rows[0].Enabled || !rows[1].Enabled {
		t.Fatalf("selected rows %+v", rows)
	}

	p.Handle(Handlers{
		Undo: func(context.Context) error { return nil },
		Redo: func(context.Context) error { return nil },
	})

	rows = p.ContextMenu()
	if !rows[4].Enabled || !rows[5].Enabled {
		t.Fatalf("handler rows %+v", rows)
	}
}
