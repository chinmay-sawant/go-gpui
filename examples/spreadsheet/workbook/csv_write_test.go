package workbook

import (
	"bytes"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/formula"
)

func TestWriteCSV(t *testing.T) {
	tab := &Table{Cells: [][]Cell{
		{
			{Kind: Text, Text: `comma, quote " and newline` + "\n"},
			{Kind: Number, Number: 2.5},
			{Kind: Formula, Source: "A1+1"},
		},
	}}

	var buf bytes.Buffer

	if err := tab.WriteCSV(&buf, DefaultCSVOptions()); err != nil {
		t.Fatal(err)
	}

	back, err := ParseCSV(buf.Bytes(), DefaultCSVOptions())
	if err != nil {
		t.Fatal(err)
	}

	if back.Cells[0][0].Text != tab.Cells[0][0].Text {
		t.Errorf("text round trip = %q", back.Cells[0][0].Text)
	}

	if back.Cells[0][1].Number != 2.5 {
		t.Errorf("number round trip = %v", back.Cells[0][1].Number)
	}

	if back.Cells[0][2].Source != "A1+1" {
		t.Errorf("formula round trip = %q", back.Cells[0][2].Source)
	}
}

func TestWriteCSVValues(t *testing.T) {
	tab := &Table{Cells: [][]Cell{{{Kind: Formula, Source: "1+1", Value: formula.Number(2)}}}}

	var buf bytes.Buffer

	opt := DefaultCSVOptions()
	opt.WriteValues = true

	if err := tab.WriteCSV(&buf, opt); err != nil {
		t.Fatal(err)
	}

	if got := buf.String(); got != "2\n" {
		t.Fatalf("values export = %q", got)
	}
}

func TestTableCommand(t *testing.T) {
	tab, err := ParseCSV([]byte("1,2\n=SUM(B6:C6),4\n"), DefaultCSVOptions())
	if err != nil {
		t.Fatal(err)
	}

	w := New(1, "test")
	s := w.AddSheet("S")

	cmd := tab.Command(s.ID(), Pos{Row: 5, Col: 1})
	if _, err := w.Apply(cmd); err != nil {
		t.Fatal(err)
	}

	if got := s.Display(Pos{5, 1}); got != "1" {
		t.Fatalf("B6 = %q", got)
	}

	if got := s.Display(Pos{6, 1}); got != "3" {
		t.Fatalf("B7 = %q, want 3", got)
	}

	if got := s.Display(Pos{6, 2}); got != "4" {
		t.Fatalf("C7 = %q", got)
	}
}
