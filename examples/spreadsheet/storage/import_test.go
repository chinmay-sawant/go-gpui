package storage

import (
	"context"
	"errors"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

func TestImportAtomicOnAbort(t *testing.T) {
	st := openMemory(t)
	ctx := context.Background()

	w, s := newBook(t, st)
	edit(t, st, w, s, workbook.Pos{Row: 9, Col: 9}, "keep me")

	tab, err := workbook.ParseCSV([]byte("a,b,c\n1,2,3\n"), workbook.DefaultCSVOptions())
	if err != nil {
		t.Fatal(err)
	}

	cmd := tab.Command(s.ID(), workbook.Pos{})

	if _, err := w.Apply(cmd); err != nil {
		t.Fatal(err)
	}

	calls := 0

	opt := ImportOptions{
		Replace: true,
		Origin:  "csv-import",
		Progress: func(done, total int) error {
			calls++

			return errors.New("user cancelled")
		},
	}

	if _, err := st.ImportCells(ctx, w.ID(), s.ID(), cmd.Edits, w.SavedRev(), opt); err == nil {
		t.Fatal("aborted import succeeded")
	}

	if calls == 0 {
		t.Fatal("progress callback never ran")
	}

	got, err := st.LoadWorkbook(ctx, w.ID())
	if err != nil {
		t.Fatal(err)
	}

	if v := got.Sheets()[0].Display(workbook.Pos{Row: 9, Col: 9}); v != "keep me" {
		t.Fatalf("A10 = %q, import was not rolled back", v)
	}

	if got.Rev() != 1 {
		t.Fatalf("rev = %d, want 1", got.Rev())
	}
}
