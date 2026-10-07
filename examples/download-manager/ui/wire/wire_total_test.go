package wire

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/store"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

func TestTotalFor(t *testing.T) {
	counts := store.Counts{ByState: map[domain.State]int{
		domain.StateCompleted: 3,
		domain.StateFailed:    2,
		domain.StateCancelled: 1,
	}}

	cases := []struct {
		filter ui.Filter
		want   int
	}{
		{ui.FilterAll, 6},
		{ui.FilterCompleted, 3},
		{ui.FilterFailed, 2},
		{ui.FilterCancelled, 1},
	}

	for _, c := range cases {
		if got := totalFor(counts, c.filter); got != c.want {
			t.Errorf("totalFor(%v) = %d, want %d", c.filter, got, c.want)
		}
	}
}
