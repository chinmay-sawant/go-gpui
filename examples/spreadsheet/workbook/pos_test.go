package workbook

import "testing"

func TestParsePos(t *testing.T) {
	cases := []struct {
		in  string
		pos Pos
		ok  bool
	}{
		{"A1", Pos{0, 0}, true},
		{"$B$2", Pos{1, 1}, true},
		{"ZZZ1048576", Pos{MaxRows - 1, MaxCols - 1}, true},
		{"A0", Pos{}, false},
		{"A1048577", Pos{}, false},
		{"AAAA1", Pos{}, false},
		{"", Pos{}, false},
	}

	for _, c := range cases {
		got, err := ParsePos(c.in)
		if (err == nil) != c.ok {
			t.Errorf("ParsePos(%q) err = %v, want ok=%v", c.in, err, c.ok)
			continue
		}

		if c.ok && got != c.pos {
			t.Errorf("ParsePos(%q) = %v, want %v", c.in, got, c.pos)
		}
	}
}

func TestColName(t *testing.T) {
	cases := []struct {
		col  int
		want string
	}{
		{0, "A"},
		{25, "Z"},
		{26, "AA"},
		{701, "ZZ"},
		{702, "AAA"},
		{18277, "ZZZ"},
		{-1, ""},
		{MaxCols, ""},
	}

	for _, c := range cases {
		if got := ColName(c.col); got != c.want {
			t.Errorf("ColName(%d) = %q, want %q", c.col, got, c.want)
		}
	}
}

func TestRect(t *testing.T) {
	r := Rect{Min: Pos{1, 1}, Max: Pos{3, 2}}

	if r.Count() != 6 {
		t.Errorf("Count = %d, want 6", r.Count())
	}

	if !r.Contains(Pos{2, 2}) || r.Contains(Pos{0, 0}) || r.Contains(Pos{4, 1}) {
		t.Error("Contains is wrong")
	}

	empty := Rect{Min: Pos{3, 3}, Max: Pos{1, 1}}
	if empty.Count() != 0 {
		t.Errorf("empty Count = %d", empty.Count())
	}
}
