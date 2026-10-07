package ui

import "testing"

func TestColName(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in   int
		want string
	}{
		{0, "A"}, {1, "B"}, {25, "Z"}, {26, "AA"}, {27, "AB"},
		{51, "AZ"}, {52, "BA"}, {701, "ZZ"}, {702, "AAA"}, {-1, ""},
	}
	for _, tc := range cases {
		if got := colName(tc.in); got != tc.want {
			t.Fatalf("colName(%d) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestRefRoundTrip(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct{ r, c int }{{0, 0}, {2, 1}, {199, 19}, {99999, 19}} {
		r, c, ok := parseRef(ref(tc.r, tc.c))
		if !ok || r != tc.r || c != tc.c {
			t.Fatalf("parseRef(ref(%d,%d)) = %d,%d,%v", tc.r, tc.c, r, c, ok)
		}
	}

	for _, bad := range []string{"", "A", "1", "A0", "A-1", "A1B", "  "} {
		if _, _, ok := parseRef(bad); ok {
			t.Fatalf("parseRef(%q) accepted", bad)
		}
	}
}

func TestTSVRoundTrip(t *testing.T) {
	t.Parallel()

	cells := []Cell{
		{Raw: "a"}, {Raw: "with\ttab"},
		{Raw: "line\nbreak"}, {Raw: `quote " here`},
	}
	text := encodeTSV(cells, 2, 2)
	rows := decodeTSV(text)
	if len(rows) != 2 || len(rows[0]) != 2 {
		t.Fatalf("decoded %v", rows)
	}

	if rows[0][1] != "with\ttab" || rows[1][0] != "line\nbreak" || rows[1][1] != `quote " here` {
		t.Fatalf("round trip lost a field: %q", rows)
	}
}

func TestDecodeTSVUnequalRows(t *testing.T) {
	t.Parallel()

	rows := decodeTSV("a\tb\tc\nd\n")
	if len(rows) != 2 || len(rows[0]) != 3 || len(rows[1]) != 1 {
		t.Fatalf("rows = %v", rows)
	}
}
