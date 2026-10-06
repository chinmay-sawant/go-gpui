package store

import (
	"strings"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func TestEnsureDummyFixtureShapes(t *testing.T) {
	st := newStore(t)

	setup, err := st.EnsureDummy(bg(), DummyOptions{})
	if err != nil {
		t.Fatal(err)
	}

	entries := allEntries(t, st, Query{Session: &setup.Session.ID})
	if len(entries) < 10000 {
		t.Fatalf("fixture has %d entries", len(entries))
	}

	sevs := map[entry.Severity]bool{}
	seen := map[string]int{}

	var badTime, multiline, unicode, long, invalid int

	for _, e := range entries {
		sevs[e.Severity] = true
		seen[e.Message]++

		if !e.TimeOK {
			badTime++
		}

		if e.Multiline {
			multiline++
		}

		if strings.Contains(e.Message, "✓") {
			unicode++
		}

		if len(e.Message) > 2000 {
			long++
		}

		if e.Malformed {
			invalid++
		}
	}

	repeated := 0

	for _, n := range seen {
		if n > 1 {
			repeated++
		}
	}

	if badTime == 0 || multiline == 0 || unicode == 0 || long == 0 || invalid == 0 || repeated == 0 {
		t.Fatalf("shapes badTime=%d multiline=%d unicode=%d long=%d invalid=%d repeated=%d",
			badTime, multiline, unicode, long, invalid, repeated)
	}

	if len(sevs) < 4 {
		t.Fatalf("severity mix = %v", sevs)
	}
}
