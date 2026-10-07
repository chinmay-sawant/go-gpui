package storage

import (
	"context"
	"errors"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

func TestCreateSaveLoad(t *testing.T) {
	st := openMemory(t)
	ctx := context.Background()

	w, s := newBook(t, st)

	if w.ID() == 0 || s.ID() == 0 {
		t.Fatalf("ids = %d/%d, want assigned", w.ID(), s.ID())
	}

	rev := edit(t, st, w, s, workbook.Pos{Row: 0, Col: 0}, "42")
	if rev != 1 {
		t.Fatalf("rev = %d, want 1", rev)
	}

	edit(t, st, w, s, workbook.Pos{Row: 0, Col: 1}, "=A1*2")

	got, err := st.LoadWorkbook(ctx, w.ID())
	if err != nil {
		t.Fatal(err)
	}

	if got.Rev() != 2 || got.SavedRev() != 2 {
		t.Fatalf("rev = %d saved = %d, want 2/2", got.Rev(), got.SavedRev())
	}

	if v := got.Sheets()[0].Display(workbook.Pos{Row: 0, Col: 0}); v != "42" {
		t.Fatalf("A1 = %q", v)
	}

	if v := got.Sheets()[0].Display(workbook.Pos{Row: 0, Col: 1}); v != "84" {
		t.Fatalf("B1 = %q, want recalculated 84", v)
	}

	if got.Dirty() {
		t.Fatal("loaded workbook reports dirty")
	}
}

func TestBlankDeletePersists(t *testing.T) {
	st := openMemory(t)
	w, s := newBook(t, st)

	edit(t, st, w, s, workbook.Pos{Row: 0, Col: 0}, "1")
	edit(t, st, w, s, workbook.Pos{Row: 0, Col: 0}, "")

	got, err := st.LoadWorkbook(context.Background(), w.ID())
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := got.Sheets()[0].Cell(workbook.Pos{Row: 0, Col: 0}); ok {
		t.Fatal("cleared cell came back")
	}
}

func TestLoadMissing(t *testing.T) {
	st := openMemory(t)

	_, err := st.LoadWorkbook(context.Background(), 999)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}

	if err := st.DeleteWorkbook(context.Background(), 999); !errors.Is(err, ErrNotFound) {
		t.Fatalf("delete err = %v", err)
	}
}
