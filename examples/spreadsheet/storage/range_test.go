package storage

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

func TestLoadRangeBounded(t *testing.T) {
	st := openMemory(t)
	ctx := context.Background()

	if _, err := st.SeedWith(ctx, 1, workbook.SeedStress(2000)); err != nil {
		t.Fatal(err)
	}

	page, err := st.ListWorkbooks(ctx, 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	w, err := st.LoadWorkbook(ctx, page.Items[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	sheet := w.Sheets()[0].ID()

	rect := workbook.Rect{Min: workbook.Pos{Row: 1500, Col: 0}, Max: workbook.Pos{Row: 1509, Col: 3}}

	edits, err := st.LoadRange(ctx, sheet, rect, 0)
	if err != nil {
		t.Fatal(err)
	}

	if len(edits) == 0 || len(edits) > 40 {
		t.Fatalf("range edits = %d, want a bounded set", len(edits))
	}

	for _, e := range edits {
		if !rect.Contains(e.Pos) {
			t.Fatalf("edit outside rect: %+v", e)
		}
	}

	limited, err := st.LoadRange(ctx, sheet, rect, 5)
	if err != nil {
		t.Fatal(err)
	}

	if len(limited) != 5 {
		t.Fatalf("limited load = %d, want 5", len(limited))
	}

	if limited[0].Pos.Row < 1500 || limited[len(limited)-1].Pos.Row > 1509 {
		t.Fatalf("limited rows = %+v", limited)
	}
}
