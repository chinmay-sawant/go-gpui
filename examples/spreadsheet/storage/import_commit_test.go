package storage

import (
	"context"
	"errors"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

func TestImportReplaceCommits(t *testing.T) {
	st := openMemory(t)
	ctx := context.Background()

	w, s := newBook(t, st)
	edit(t, st, w, s, workbook.Pos{Row: 9, Col: 9}, "old")

	tab, err := workbook.ParseCSV([]byte("1,2\n"), workbook.DefaultCSVOptions())
	if err != nil {
		t.Fatal(err)
	}

	cmd := tab.Command(s.ID(), workbook.Pos{})

	if _, err := w.Apply(cmd); err != nil {
		t.Fatal(err)
	}

	rev, err := st.ImportCells(ctx, w.ID(), s.ID(), cmd.Edits, w.SavedRev(), ImportOptions{Replace: true})
	if err != nil {
		t.Fatal(err)
	}

	if rev != 2 {
		t.Fatalf("rev = %d, want 2", rev)
	}

	got, err := st.LoadWorkbook(ctx, w.ID())
	if err != nil {
		t.Fatal(err)
	}

	if v := got.Sheets()[0].Display(workbook.Pos{Row: 9, Col: 9}); v != "" {
		t.Fatalf("old cell survived replace: %q", v)
	}

	if v := got.Sheets()[0].Display(workbook.Pos{Row: 0, Col: 1}); v != "2" {
		t.Fatalf("B1 = %q", v)
	}
}

func TestDeleteDuringPendingSave(t *testing.T) {
	st := openMemory(t)
	ctx := context.Background()

	w, s := newBook(t, st)

	if err := st.DeleteWorkbook(ctx, w.ID()); err != nil {
		t.Fatal(err)
	}

	cmd := []workbook.CellEdit{{Pos: workbook.Pos{Row: 0, Col: 0}, Cell: workbook.ParseInput("late")}}

	if _, err := st.SaveCells(ctx, w.ID(), s.ID(), cmd, 0, "pending"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("save after delete = %v, want ErrNotFound", err)
	}

	page, err := st.ListWorkbooks(ctx, 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	if len(page.Items) != 0 {
		t.Fatalf("workbooks = %d after delete", len(page.Items))
	}
}
