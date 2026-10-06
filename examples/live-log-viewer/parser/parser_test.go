package parser

import (
	"strings"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func TestParseLine(t *testing.T) {
	got := ParseLine([]byte("2026-10-07T09:00:00.123Z ERROR boom"))
	if !got.TimeOK {
		t.Fatalf("time not parsed: %+v", got)
	}

	if got.Severity != entry.Error {
		t.Fatalf("severity = %v, want ERROR", got.Severity)
	}

	if got.TimeRaw != "2026-10-07T09:00:00.123Z" {
		t.Fatalf("raw time = %q", got.TimeRaw)
	}
}

func TestParseLineMalformedEncoding(t *testing.T) {
	got := ParseLine([]byte("INFO bad \xff\xfe byte\r"))
	if !got.Malformed {
		t.Fatal("invalid bytes not reported")
	}

	if !strings.ContainsRune(got.Text, '\uFFFD') {
		t.Fatalf("no replacement rune in %q", got.Text)
	}

	if strings.HasSuffix(got.Text, "\r") {
		t.Fatal("trailing CR kept")
	}

	if got.Severity != entry.Info {
		t.Fatalf("severity = %v", got.Severity)
	}
}

func TestParseLineNoTimestamp(t *testing.T) {
	got := ParseLine([]byte("just a message"))
	if got.TimeOK {
		t.Fatalf("unexpected time %v", got.Time)
	}

	if got.TimeRaw != "" {
		t.Fatalf("raw = %q", got.TimeRaw)
	}

	if got.Severity != entry.Unknown {
		t.Fatalf("severity = %v", got.Severity)
	}
}
