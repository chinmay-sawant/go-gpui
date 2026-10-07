package workbook

import "testing"

func TestParseInput(t *testing.T) {
	cases := []struct {
		in   string
		kind Kind
		num  float64
		text string
	}{
		{"", Blank, 0, ""},
		{"0", Number, 0, ""},
		{"-5", Number, -5, ""},
		{"2.5", Number, 2.5, ""},
		{"1e12", Number, 1e12, ""},
		{" 12", Text, 0, " 12"},
		{"abc", Text, 0, "abc"},
		{"0x10", Text, 0, "0x10"},
		{"Türkçe", Text, 0, "Türkçe"},
		{"=A1+1", Formula, 0, ""},
		{"Inf", Text, 0, "Inf"},
		{"NaN", Text, 0, "NaN"},
	}

	for _, c := range cases {
		got := ParseInput(c.in)
		if got.Kind != c.kind || got.Number != c.num || got.Text != c.text {
			t.Errorf("ParseInput(%q) = %+v, want kind %d num %v text %q",
				c.in, got, c.kind, c.num, c.text)
		}
	}
}

func TestCellDisplayAndRaw(t *testing.T) {
	if got := (Cell{Kind: Number, Number: 0}).Display(); got != "0" {
		t.Errorf("zero display = %q", got)
	}

	if got := (Cell{Kind: Number, Number: 1e12}).Display(); got != "1000000000000" {
		t.Errorf("large display = %q", got)
	}

	if got := (Cell{Kind: Text, Text: "0"}).Display(); got != "0" {
		t.Errorf("text zero display = %q", got)
	}

	if got := (Cell{Kind: Formula, Source: "A1+1"}).Raw(); got != "=A1+1" {
		t.Errorf("formula raw = %q", got)
	}

	if got := (Cell{Kind: Number, Number: -2.5}).Raw(); got != "-2.5" {
		t.Errorf("number raw = %q", got)
	}

	if got := (Cell{}).Raw(); got != "" {
		t.Errorf("blank raw = %q", got)
	}
}

func TestBlankVersusZero(t *testing.T) {
	w := New(1, "test")
	s := w.AddSheet("S")

	if _, err := w.Apply(Command{Sheet: s.ID(), Edits: []CellEdit{
		{Pos: Pos{0, 0}, Cell: ParseInput("0")},
		{Pos: Pos{0, 1}, Cell: ParseInput("")},
	}}); err != nil {
		t.Fatal(err)
	}

	if c, ok := s.Cell(Pos{0, 0}); !ok || c.Kind != Number {
		t.Fatalf("A1 = %+v ok=%v, want a stored zero", c, ok)
	}

	if _, ok := s.Cell(Pos{0, 1}); ok {
		t.Fatal("B1 is stored, want blank")
	}

	if v := s.Value(Pos{0, 1}); !v.IsBlank() {
		t.Fatalf("B1 value = %v, want blank", v)
	}
}
