package parser

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func TestFindSeverity(t *testing.T) {
	cases := []struct {
		line string
		want entry.Severity
	}{
		{"2026-10-07T09:00:00Z INFO hello", entry.Info},
		{"[warn] disk almost full", entry.Warn},
		{"2026-10-07 09:00:00 ERROR boom", entry.Error},
		{`{"level":"fatal","msg":"down"}`, entry.Fatal},
		{`level=debug retry`, entry.Debug},
		{`severity=notice started`, entry.Info},
		{"WARNING: slow", entry.Warn},
		{"critical failure", entry.Fatal},
		{"plain message with error inside", entry.Unknown},
		{"", entry.Unknown},
	}

	for _, c := range cases {
		if got := findSeverity(c.line); got != c.want {
			t.Errorf("findSeverity(%q) = %v, want %v", c.line, got, c.want)
		}
	}
}
