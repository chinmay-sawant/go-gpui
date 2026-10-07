package parser

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func recAt(offset int64, text string) entry.RawRecord {
	return entry.RawRecord{
		Path: "test", Generation: 1, Offset: offset,
		Bytes: len(text), Data: []byte(text),
	}
}

func TestGrouperMultiline(t *testing.T) {
	g := NewGrouper(Options{MaxLines: 4, MaxBytes: 4096})

	var out []entry.Entry

	out = append(out, g.Add(recAt(1, "ERROR panic: boom"))...)
	out = append(out, g.Add(recAt(20, "\tat main.run()"))...)
	out = append(out, g.Add(recAt(40, "INFO recovered"))...)

	if len(out) != 1 {
		t.Fatalf("got %d entries, want 1", len(out))
	}

	if !out[0].Multiline {
		t.Fatal("not marked multiline")
	}

	if out[0].Message != "ERROR panic: boom\n\tat main.run()" {
		t.Fatalf("message = %q", out[0].Message)
	}

	gen, pos, ok := g.Safe()
	if !ok || gen != 1 || pos != out[0].Position+int64(out[0].Bytes) {
		t.Fatalf("safe = %d %d %v", gen, pos, ok)
	}

	rest := g.Flush()
	if len(rest) != 1 || rest[0].Message != "INFO recovered" {
		t.Fatalf("flush = %+v", rest)
	}
}

func TestGrouperLineLimit(t *testing.T) {
	g := NewGrouper(Options{MaxLines: 2, MaxBytes: 4096})

	g.Add(recAt(1, "ERROR one"))
	g.Add(recAt(20, "  cont1"))

	out := g.Add(recAt(40, "  cont2"))
	if len(out) != 1 {
		t.Fatalf("got %d entries, want the first closed", len(out))
	}

	if out[0].Message != "ERROR one\n  cont1" {
		t.Fatalf("message = %q", out[0].Message)
	}
}
