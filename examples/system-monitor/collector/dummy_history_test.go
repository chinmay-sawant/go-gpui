package collector

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// TestDummyHistory checks the shape of the seeded recording and that it ends
// at the requested time with rates already computed.
func TestDummyHistory(t *testing.T) {
	d := NewDummy(DummyOptions{Seed: 9})
	now := time.Unix(1_700_000_000, 0)

	hist := d.Historical(now, 120, 15*time.Second)
	if len(hist) != 120 {
		t.Fatalf("samples = %d", len(hist))
	}
	if !hist[119].Stamp.At.Equal(now) {
		t.Fatalf("last stamp = %v", hist[119].Stamp.At)
	}
	if want := now.Add(-119 * 15 * time.Second); !hist[0].Stamp.At.Equal(want) {
		t.Fatalf("first stamp = %v, want %v", hist[0].Stamp.At, want)
	}
	if !hist[60].CPUPercent.Valid || !hist[60].Mem.UsedPercent.Valid {
		t.Fatal("history sample has no rates")
	}
	if hist[60].Gap {
		t.Fatal("history sample marked as a gap")
	}
}

// TestDummyHistoryContinues checks that the first live sample after a prefill
// computes its rate from the history's last counters.
func TestDummyHistoryContinues(t *testing.T) {
	base := time.Unix(1_700_000_000, 0)
	d := NewDummy(DummyOptions{Seed: 4, Stamp: func() domain.Stamp {
		return domain.Stamp{At: base.Add(2 * time.Second), Mono: 2 * time.Second}
	}})

	hist := d.Historical(base, 3, time.Second)
	live, err := d.Sample(t.Context())
	if err != nil {
		t.Fatal(err)
	}

	got := domain.Compute(hist[len(hist)-1], live)
	if !got.CPUPercent.Valid {
		t.Fatal("first live sample has no rate after prefill")
	}
}
