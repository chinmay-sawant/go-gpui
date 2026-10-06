package parser

import "testing"

func TestFindTime(t *testing.T) {
	cases := []struct {
		text string
		ok   bool
	}{
		{"2026-10-07T09:00:00Z x", true},
		{"2026-10-07T09:00:00.123456789+02:00 x", true},
		{"2026-10-07 09:00:00.123 x", true},
		{"[2026-10-07 09:00:00] x", true},
		{"2026/10/07 09:00:00 x", true},
		{"10/07/2026 09:00:00 x", true},
		{"1759827600 started", true},
		{"1759827600.5 started", true},
		{"09:00:00 x", true},
		{"Oct  7 09:00:00 host x", true},
		{"2026-13-45 99:61:61 x", false},
		{"not-a-time x", false},
		{"", false},
	}

	for _, c := range cases {
		_, raw, ok := findTime(c.text)
		if ok != c.ok {
			t.Errorf("findTime(%q) ok = %v, want %v", c.text, ok, c.ok)
		}

		if ok && raw == "" {
			t.Errorf("findTime(%q) matched with empty raw text", c.text)
		}
	}
}

func TestFindTimeOffsetPreserved(t *testing.T) {
	ts, _, ok := findTime("2026-10-07T09:00:00+02:00 x")
	if !ok {
		t.Fatal("not parsed")
	}

	if _, offset := ts.Zone(); offset != 7200 {
		t.Fatalf("offset = %d, want 7200", offset)
	}
}
