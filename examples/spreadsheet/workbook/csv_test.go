package workbook

import (
	"testing"
)

func TestParseCSV(t *testing.T) {
	data := []byte("name,qty,note\r\nTürkçe,2,\"a,b\"\r\n\"multi\nline\",,\"say \"\"hi\"\"\"\r\n")

	tab, err := ParseCSV(data, DefaultCSVOptions())
	if err != nil {
		t.Fatal(err)
	}

	if tab.Rows() != 3 || tab.Cols() != 3 {
		t.Fatalf("shape = %dx%d, want 3x3", tab.Rows(), tab.Cols())
	}

	if got := tab.Cells[1][0].Text; got != "Türkçe" {
		t.Errorf("A2 = %q", got)
	}

	if c := tab.Cells[1][1]; c.Kind != Number || c.Number != 2 {
		t.Errorf("B2 = %+v, want number 2", c)
	}

	if got := tab.Cells[1][2].Text; got != "a,b" {
		t.Errorf("C2 = %q", got)
	}

	if got := tab.Cells[2][0].Text; got != "multi\nline" {
		t.Errorf("A3 = %q", got)
	}

	if c := tab.Cells[2][1]; c.Kind != Blank {
		t.Errorf("B3 = %+v, want blank", c)
	}

	if got := tab.Cells[2][2].Text; got != `say "hi"` {
		t.Errorf("C3 = %q", got)
	}
}

func TestCSVBOM(t *testing.T) {
	tab, err := ParseCSV([]byte("\xEF\xBB\xBFa,b\n"), DefaultCSVOptions())
	if err != nil {
		t.Fatal(err)
	}

	if got := tab.Cells[0][0].Text; got != "a" {
		t.Fatalf("A1 = %q, want the BOM stripped", got)
	}
}

func TestCSVUnequalRows(t *testing.T) {
	tab, err := ParseCSV([]byte("a\nb,c\n"), DefaultCSVOptions())
	if err != nil {
		t.Fatal(err)
	}

	if tab.Rows() != 2 || tab.Cols() != 2 {
		t.Fatalf("shape = %dx%d", tab.Rows(), tab.Cols())
	}

	if len(tab.Cells[0]) != 1 {
		t.Fatalf("row 0 length = %d", len(tab.Cells[0]))
	}
}
