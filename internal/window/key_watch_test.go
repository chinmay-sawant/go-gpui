package window

import "testing"

// TestKeyWatchDropsRepeatPulses checks the auto-repeat shape a server can
// report: a release and a press with only one frame between them.
func TestKeyWatchDropsRepeatPulses(t *testing.T) {
	t.Parallel()

	w := newKeyWatch()

	if down, _ := w.step(0, true); !down {
		t.Fatal("the first press did not fire")
	}

	for range keyGuardFrames {
		if down, up := w.step(0, false); down || up {
			t.Fatal("a release before the guard fired")
		}

		if down, _ := w.step(0, true); down {
			t.Fatal("a repeat pulse fired a press")
		}
	}
}

func TestKeyWatchFiresRealEvents(t *testing.T) {
	t.Parallel()

	w := newKeyWatch()

	if down, _ := w.step(0, true); !down {
		t.Fatal("the first press did not fire")
	}

	for range keyGuardFrames - 1 {
		if down, up := w.step(0, true); down || up {
			t.Fatal("a held key fired again")
		}
	}

	if _, up := w.step(0, false); !up {
		t.Fatal("the release did not fire")
	}

	for range keyGuardFrames {
		w.step(0, false)
	}

	if down, _ := w.step(0, true); !down {
		t.Fatal("the next real press did not fire")
	}
}
