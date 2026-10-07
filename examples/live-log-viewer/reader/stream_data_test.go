package reader

import (
	"bytes"
	"strings"
	"testing"
)

func TestDummyDeterministic(t *testing.T) {
	a := DummyRecords(7, "api", 1, 50, 0)
	b := DummyRecords(7, "api", 1, 50, 0)

	for i := range a {
		if !bytes.Equal(a[i].Data, b[i].Data) {
			t.Fatalf("record %d differs", i)
		}
	}

	c := DummyRecords(8, "api", 1, 50, 0)

	same := 0

	for i := range a {
		if bytes.Equal(a[i].Data, c[i].Data) {
			same++
		}
	}

	if same == len(a) {
		t.Fatal("different seeds produced identical records")
	}
}

func TestDummyShapes(t *testing.T) {
	recs := DummyRecords(1, "api", 1, 1000, 0)

	var malformed, unicode, long, invalid, multi bool

	for _, r := range recs {
		s := string(r.Data)

		malformed = malformed || strings.HasPrefix(s, "2026-13-45")
		unicode = unicode || strings.Contains(s, "✓")
		long = long || len(s) > 2000
		invalid = invalid || strings.Contains(s, "\x80")
		multi = multi || strings.Contains(s, "\n")
	}

	if !malformed || !unicode || !long || !invalid || !multi {
		t.Fatalf("shapes malformed=%v unicode=%v long=%v invalid=%v multi=%v",
			malformed, unicode, long, invalid, multi)
	}
}

func TestDummyMaxRecord(t *testing.T) {
	recs := DummyRecords(1, "api", 1, 1000, 64)

	for _, r := range recs {
		if len(r.Data) > 64 {
			t.Fatalf("record %d is %d bytes", r.Offset, len(r.Data))
		}
	}
}
