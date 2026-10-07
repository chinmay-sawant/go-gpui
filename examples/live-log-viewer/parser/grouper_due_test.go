package parser

import (
	"testing"
	"time"
)

func TestGrouperDue(t *testing.T) {
	g := NewGrouper(Options{FlushAfter: 50 * time.Millisecond})

	g.Add(recAt(1, "INFO alone"))

	if out := g.Due(time.Now()); len(out) != 0 {
		t.Fatalf("flushed early: %+v", out)
	}

	if out := g.Due(time.Now().Add(60 * time.Millisecond)); len(out) != 1 {
		t.Fatalf("hold did not release the record: %+v", out)
	}

	if out := g.Due(time.Now().Add(time.Second)); len(out) != 0 {
		t.Fatalf("flushed twice: %+v", out)
	}
}
