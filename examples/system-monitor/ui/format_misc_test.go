package ui

import (
	"testing"
	"time"
)

func TestFormatCount(t *testing.T) {
	tests := map[int]string{0: "0", 12: "12", 999: "999", 1000: "1,000",
		12345: "12,345", 1234567: "1,234,567"}

	for n, want := range tests {
		if got := formatCount(n); got != want {
			t.Errorf("formatCount(%d) = %q, want %q", n, got, want)
		}
	}
}

func TestFormatAge(t *testing.T) {
	tests := map[int64]string{42: "42s", 90: "1m 30s", 3600: "1h 0m",
		43200: "12h 0m", 3*86400 + 4*3600: "3d 4h"}

	for sec, want := range tests {
		if got := formatAge(time.Duration(sec) * time.Second); got != want {
			t.Errorf("formatAge(%ds) = %q, want %q", sec, got, want)
		}
	}
}

func TestTruncate(t *testing.T) {
	if got := truncate("hello", 10); got != "hello" {
		t.Errorf("truncate short = %q", got)
	}

	if got := truncate("hello world", 8); got != "hello w\u2026" {
		t.Errorf("truncate long = %q", got)
	}

	if got := truncate("hello", 1); got != "h" {
		t.Errorf("truncate tiny = %q", got)
	}
}
