package window

import (
	"testing"
	"time"
)

// TestChordWatchDropsRepeats checks the auto-repeat shapes a server can
// report: a release and a press that share one frame, and a press a frame
// or two after the release. A press that comes after the gap is real.
func TestChordWatchDropsRepeats(t *testing.T) {
	t.Parallel()

	w := newChordWatch()
	t0 := time.Unix(0, 0)

	// A real press fires once, and a held key does not fire again.
	if !w.step(0, true, false, t0) {
		t.Fatal("a real press did not fire")
	}

	if w.step(0, false, false, t0.Add(20*time.Millisecond)) {
		t.Fatal("a held key fired again")
	}

	// The release and a press in the same frame are one repeat pulse.
	if w.step(0, true, true, t0.Add(60*time.Millisecond)) {
		t.Fatal("a same-frame release and press fired")
	}

	// A press a frame after the release is another repeat pulse.
	if w.step(0, true, false, t0.Add(70*time.Millisecond)) {
		t.Fatal("a press right after a release fired")
	}

	// A press after the gap is a real press.
	if !w.step(0, true, false, t0.Add(200*time.Millisecond)) {
		t.Fatal("a press after the gap did not fire")
	}
}

// TestChordWatchSwallowsFiredKey checks that a chord key stops typing
// after its chord fired, and that it types again once it is up.
func TestChordWatchSwallowsFiredKey(t *testing.T) {
	t.Parallel()

	w := newChordWatch()
	w.eaten[0] = "c"

	if got := string(w.filter([]rune("cv"))); got != "v" {
		t.Fatalf("filter = %q, want v", got)
	}

	// The key is not pressed in this test, so the filter forgot it.
	if got := string(w.filter([]rune("c"))); got != "c" {
		t.Fatalf("filter after the key went up = %q, want c", got)
	}
}
