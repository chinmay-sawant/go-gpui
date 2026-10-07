package ui

import "testing"

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		v    float64
		ok   bool
		want string
	}{
		{0, true, "0 B"},
		{512, true, "512 B"},
		{1024, true, "1.0 KiB"},
		{1536, true, "1.5 KiB"},
		{1023.9 * 1024, true, "1023.9 KiB"},
		{1023.96 * 1024, true, "1.0 MiB"},
		{8 * 1024 * 1024 * 1024, true, "8.0 GiB"},
		{-1, true, "0 B"},
		{10, false, "n/a"},
	}

	for _, tt := range tests {
		if got := formatBytes(tt.v, tt.ok); got != tt.want {
			t.Errorf("formatBytes(%v, %v) = %q, want %q", tt.v, tt.ok, got, tt.want)
		}
	}
}

func TestFormatRate(t *testing.T) {
	if got := formatRate(1536, true); got != "1.5 KiB/s" {
		t.Errorf("formatRate = %q", got)
	}

	if got := formatRate(0, false); got != "n/a" {
		t.Errorf("formatRate missing = %q", got)
	}
}

func TestFormatPercent(t *testing.T) {
	if got := formatPercent(0.5, true); got != "50.0%" {
		t.Errorf("formatPercent = %q", got)
	}

	if got := formatPercent(0.123, true); got != "12.3%" {
		t.Errorf("formatPercent rounding = %q", got)
	}

	if got := formatPercent(0, false); got != "n/a" {
		t.Errorf("formatPercent missing = %q", got)
	}
}
