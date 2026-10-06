package wire

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/store"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

func TestCursorRoundTrip(t *testing.T) {
	at := time.Date(2026, 10, 7, 1, 2, 3, 456_000_000, time.UTC)
	in := store.Cursor{UpdatedAt: at, ID: "j-7"}

	got := decodeCursor(encodeCursor(in))
	if !got.UpdatedAt.Equal(at) || got.ID != "j-7" {
		t.Fatalf("round trip = %+v", got)
	}

	if encodeCursor(store.Cursor{}) != "" {
		t.Fatal("zero cursor did not encode empty")
	}

	if decodeCursor("garbage") != (store.Cursor{}) {
		t.Fatal("garbage cursor did not decode to zero")
	}
}

func TestMapState(t *testing.T) {
	cases := []struct {
		in   domain.State
		want ui.State
	}{
		{domain.StateQueued, ui.StateQueued},
		{domain.StateRunning, ui.StateRunning},
		{domain.StatePaused, ui.StatePaused},
		{domain.StateCompleted, ui.StateCompleted},
		{domain.StateFailed, ui.StateFailed},
		{domain.StateCancelled, ui.StateCancelled},
		{"bogus", ui.StateQueued},
	}

	for _, c := range cases {
		if got := mapState(c.in); got != c.want {
			t.Errorf("mapState(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestTotalOf(t *testing.T) {
	if got := totalOf(domain.Job{Expected: 100, Total: 40}); got != 100 {
		t.Fatalf("expected length not preferred: %d", got)
	}

	if got := totalOf(domain.Job{Expected: -1, Total: 40}); got != 40 {
		t.Fatalf("observed length not used: %d", got)
	}
}
