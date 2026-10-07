package parser

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func TestGrouperGenerationChange(t *testing.T) {
	g := NewGrouper(Options{})

	g.Add(recAt(1, "old pending"))

	out := g.Add(entry.RawRecord{
		Path: "test", Generation: 2, Offset: 0, Bytes: 3, Data: []byte("new"),
	})
	if len(out) != 1 || out[0].Generation != 1 {
		t.Fatalf("generation flush = %+v", out)
	}

	g.Flush()

	if _, pos, ok := g.Safe(); !ok || pos != 3 {
		t.Fatalf("safe after gen 2 = %d %v", pos, ok)
	}
}

func TestGrouperPartial(t *testing.T) {
	g := NewGrouper(Options{})

	rec := recAt(5, "INFO cut off")
	rec.Partial = true
	rec.Bytes = 12

	out := g.Add(rec)
	if len(out) != 1 || !out[0].Partial {
		t.Fatalf("partial = %+v", out)
	}
}

func TestGrouperMalformedPropagates(t *testing.T) {
	g := NewGrouper(Options{})

	g.Add(recAt(1, "ERROR bad \xff"))

	out := g.Flush()
	if len(out) != 1 || !out[0].Malformed {
		t.Fatalf("malformed not propagated: %+v", out)
	}
}
