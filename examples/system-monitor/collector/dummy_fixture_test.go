package collector

import (
	"testing"
	"time"
)

// TestDummyFixture checks the fixture metadata storage seeds from.
func TestDummyFixture(t *testing.T) {
	now := time.Unix(1_700_000_000, 0)
	f := DummyFixture(2, now, 120, 15*time.Second)

	if f.Version != SeedVersion || f.Source != "dummy" {
		t.Fatalf("fixture = %+v", f)
	}
	if len(f.Samples) != 120 {
		t.Fatalf("samples = %d", len(f.Samples))
	}
	if f.Started.After(now) || f.Started.Equal(time.Time{}) {
		t.Fatalf("started = %v", f.Started)
	}
	if len(f.Samples[0].Records()) == 0 {
		t.Fatal("fixture sample has no records")
	}
}
