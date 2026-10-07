package transfer

import "testing"

// TestParseContentRange covers the accepted and refused shapes.
func TestParseContentRange(t *testing.T) {
	good := []struct {
		in                string
		start, end, total int64
	}{
		{"bytes 10-19/100", 10, 19, 100},
		{"bytes */100", 0, 0, 100},
		{"bytes 0-9/*", 0, 9, Unknown},
		{"bytes 0-0/1", 0, 0, 1},
		{" bytes 5-5/6 ", 5, 5, 6},
	}

	for _, c := range good {
		start, end, total, err := parseContentRange(c.in)
		if err != nil {
			t.Errorf("parseContentRange(%q): %v", c.in, err)
			continue
		}

		if start != c.start || end != c.end || total != c.total {
			t.Errorf("parseContentRange(%q) = %d,%d,%d, want %d,%d,%d",
				c.in, start, end, total, c.start, c.end, c.total)
		}
	}

	bad := []string{
		"",
		"bytes",
		"bytes 5-/100",
		"bytes 5-4/100",
		"bytes x-9/100",
		"bytes 0-9/-1",
		"bytes 0-9/abc",
	}

	for _, in := range bad {
		if _, _, _, err := parseContentRange(in); err == nil {
			t.Errorf("parseContentRange(%q) accepted", in)
		}
	}
}
