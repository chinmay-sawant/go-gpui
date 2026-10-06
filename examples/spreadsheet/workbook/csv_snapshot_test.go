package workbook

import (
	"bytes"
	"testing"
)

func TestWriteCSVBOM(t *testing.T) {
	tab := &Table{Cells: [][]Cell{{{Kind: Text, Text: "a"}}}}

	opt := DefaultCSVOptions()
	opt.WriteBOM = true

	var buf bytes.Buffer

	if err := tab.WriteCSV(&buf, opt); err != nil {
		t.Fatal(err)
	}

	if !bytes.HasPrefix(buf.Bytes(), []byte{0xEF, 0xBB, 0xBF}) {
		t.Fatalf("missing BOM: % x", buf.Bytes())
	}
}

func TestSnapshot(t *testing.T) {
	w := New(1, "test")
	s := w.AddSheet("S")

	mustApply(t, w, s, Pos{1, 1}, "7")
	mustApply(t, w, s, Pos{2, 2}, "=A1+1")

	tab := Snapshot(s, Rect{Min: Pos{1, 1}, Max: Pos{2, 2}})

	if tab.Start != (Pos{1, 1}) || tab.Rows() != 2 || tab.Cols() != 2 {
		t.Fatalf("snapshot = %+v %dx%d", tab.Start, tab.Rows(), tab.Cols())
	}

	if tab.Cells[0][0].Number != 7 {
		t.Fatalf("cell = %+v", tab.Cells[0][0])
	}

	if tab.Cells[1][1].Source != "A1+1" {
		t.Fatalf("formula = %+v", tab.Cells[1][1])
	}
}
