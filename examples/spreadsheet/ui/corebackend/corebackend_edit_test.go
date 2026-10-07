package corebackend

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/ui"
)

func TestApplyUndoRedoPersist(t *testing.T) {
	dir := t.TempDir()
	b := newBackend(t, dir)
	id := sheetID(t, b, "Numbers")
	area := ui.Area{R0: 0, C0: 10, R1: 0, C1: 10}

	if _, err := b.Apply(id, []ui.Edit{{Row: 0, Col: 10, Raw: "1234"}}); err != nil {
		t.Fatal(err)
	}

	if _, err := b.Apply(id, []ui.Edit{{Row: 0, Col: 10, Raw: "=B1+1"}}); err != nil {
		t.Fatal(err)
	}

	cells, _ := b.Range(id, area)
	if cells[0].Raw != "=B1+1" || cells[0].Display != "1" {
		t.Fatalf("after edits = %+v", cells[0])
	}

	if res, err := b.Undo(); err != nil || !res.OK {
		t.Fatalf("undo = %+v err=%v", res, err)
	}

	if res, err := b.Undo(); err != nil || !res.OK {
		t.Fatalf("undo 2 = %+v err=%v", res, err)
	}

	if res, _ := b.Undo(); res.OK {
		t.Fatal("third undo reported a step")
	}

	if res, err := b.Redo(); err != nil || !res.OK {
		t.Fatalf("redo = %+v err=%v", res, err)
	}

	if cells, _ = b.Range(id, area); cells[0].Raw != "1234" {
		t.Fatalf("after redo = %+v", cells[0])
	}

	if err := b.Close(); err != nil {
		t.Fatal(err)
	}

	again := newBackend(t, dir)
	cells, err := again.Range(sheetID(t, again, "Numbers"), area)
	if err != nil {
		t.Fatal(err)
	}

	if cells[0].Raw != "1234" {
		t.Fatalf("reopened cell = %+v", cells[0])
	}
}
