package storage

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

// TestImportAtomicOnAbort interrupts a 400-row import at the second
// progress batch, halfway through the transaction, and checks the stored
// workbook is untouched.
func TestImportAtomicOnAbort(t *testing.T) {
	st := openMemory(t)
	ctx := context.Background()

	w, s := newBook(t, st)
	edit(t, st, w, s, workbook.Pos{Row: 399, Col: 9}, "keep me")

	var csv strings.Builder

	for i := 0; i < 400; i++ {
		csv.WriteString("1\n")
	}

	tab, err := workbook.ParseCSV([]byte(csv.String()), workbook.DefaultCSVOptions())
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

			if done >= 256 {
				return errors.New("user cancelled halfway")
			}

			return nil
		},
	}

	if _, err := st.ImportCells(ctx, w.ID(), s.ID(), cmd.Edits, w.SavedRev(), opt); err == nil {
		t.Fatal("aborted import succeeded")
	}

	if calls < 2 {
		t.Fatalf("progress calls = %d, want the abort halfway", calls)
	}

	got, err := st.LoadWorkbook(ctx, w.ID())
	if err != nil {
		t.Fatal(err)
	}

	if v := got.Sheets()[0].Display(workbook.Pos{Row: 399, Col: 9}); v != "keep me" {
		t.Fatalf("J400 = %q, import was not rolled back", v)
	}

	if v := got.Sheets()[0].Display(workbook.Pos{Row: 0, Col: 0}); v != "" {
		t.Fatalf("A1 = %q, half the import leaked", v)
	}

	if got.Rev() != 1 {
		t.Fatalf("rev = %d, want 1", got.Rev())
	}
}
