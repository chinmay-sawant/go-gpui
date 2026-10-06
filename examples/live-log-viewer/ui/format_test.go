package ui

import (
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	ok := Entry{TimeOK: true, Time: time.Date(2026, 10, 7, 13, 4, 5, 123000000, time.UTC)}
	if got := formatTime(ok); got != "13:04:05.123" {
		t.Fatalf("time = %q", got)
	}

	raw := Entry{TimeRaw: "not-a-timestamp"}
	if got := formatTime(raw); got != "not-a-times…" {
		t.Fatalf("raw = %q", got)
	}

	if got := formatTime(Entry{}); got != "--:--:--" {
		t.Fatalf("empty = %q", got)
	}
}

func TestSeverity(t *testing.T) {
	if sevClass("WARNING") != "warn" || sevClass("critical") != "fatal" {
		t.Fatal("severity class mapping")
	}

	if sevLabel("warn") != "WAR" || sevLabel("fatal") != "FAT" {
		t.Fatal("severity label")
	}

	if sevLabel("nonsense") != "DBG" {
		t.Fatal("unknown severity label")
	}
}

func TestTruncateAndSplit(t *testing.T) {
	if got := truncate("héllo wörld", 6); got != "héllo…" {
		t.Fatalf("truncate = %q", got)
	}

	if got := truncate("short", 10); got != "short" {
		t.Fatalf("short = %q", got)
	}

	if got := lineCount("a\nb\nc"); got != 3 {
		t.Fatalf("lines = %d", got)
	}

	if got := splitLines("a\r\nb"); len(got) != 2 || got[1] != "b" {
		t.Fatalf("split = %#v", got)
	}
}

func TestBytesText(t *testing.T) {
	if got := bytesText(512); got != "512B" {
		t.Fatalf("bytes = %q", got)
	}

	if got := bytesText(2048); got != "2.0K" {
		t.Fatalf("kilo = %q", got)
	}
}
