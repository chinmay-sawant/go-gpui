package collector

import (
	"reflect"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// fixedStamps advances one second per call, so two sources with the same seed
// see the same clock.
func fixedStamps() func() domain.Stamp {
	base := time.Unix(1_700_000_000, 0)
	n := -1

	return func() domain.Stamp {
		n++

		return domain.Stamp{At: base.Add(time.Duration(n) * time.Second), Mono: time.Duration(n) * time.Second}
	}
}

// TestDummyDeterministic checks that one seed and one call sequence produce
// identical samples and process tables.
func TestDummyDeterministic(t *testing.T) {
	run := func() ([]domain.Sample, [][]domain.Process) {
		d := NewDummy(DummyOptions{Seed: 7, Stamp: fixedStamps()})

		var samples []domain.Sample

		var tables [][]domain.Process

		for range 5 {
			s, err := d.Sample(t.Context())
			if err != nil {
				t.Fatal(err)
			}
			samples = append(samples, s)

			rows, err := d.Processes(t.Context())
			if err != nil {
				t.Fatal(err)
			}
			tables = append(tables, rows)
		}

		return samples, tables
	}

	a, at := run()
	b, bt := run()

	if !reflect.DeepEqual(a, b) {
		t.Fatal("samples differ between runs")
	}
	if !reflect.DeepEqual(at, bt) {
		t.Fatal("process tables differ between runs")
	}
}
