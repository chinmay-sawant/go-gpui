package collector

import "testing"

// TestDummySpike checks that spike ticks raise total busy time well above the
// quiet baseline.
func TestDummySpike(t *testing.T) {
	d := NewDummy(DummyOptions{Seed: 5, Stamp: fixedStamps()})

	prev, err := d.Sample(t.Context())
	if err != nil {
		t.Fatal(err)
	}

	max, sum, n := 0.0, 0.0, 0

	for range 40 {
		next, err := d.Sample(t.Context())
		if err != nil {
			t.Fatal(err)
		}

		total := float64(next.CPUTotal.Total - prev.CPUTotal.Total)
		busy := float64(next.CPUTotal.Busy - prev.CPUTotal.Busy)
		frac := busy / total

		if frac > max {
			max = frac
		}
		sum += frac
		n++
		prev = next
	}

	if max < 0.4 {
		t.Fatalf("max busy fraction = %.3f, want a spike", max)
	}
	if sum/float64(n) > 0.35 {
		t.Fatalf("average busy fraction = %.3f, baseline too high", sum/float64(n))
	}
}
