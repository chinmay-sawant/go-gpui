package store

import (
	"context"
	"database/sql"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// DummyOptions tunes the demo fixture.
type DummyOptions struct {
	Count  int
	Rate   int
	Seed   int64
	Policy entry.Policy
}

// DummySetup is what EnsureDummy returns: the session, its sources with
// their committed checkpoints, and whether this call seeded the fixture.
type DummySetup struct {
	Session entry.Session
	Sources []entry.Source
	Seeded  bool
	Version int
}

const (
	dummySession = "Dummy log stream"
	dummySeedVer = 1
)

var dummySourceKeys = []struct{ key, label string }{
	{"api", "api"},
	{"worker", "worker"},
	{"db", "database"},
}

// EnsureDummy creates the demo session and its generated sources, seeds the
// initial entries once, and reports the committed checkpoints. Reopening
// never duplicates or overwrites the fixture.
func (s *Store) EnsureDummy(ctx context.Context, o DummyOptions) (DummySetup, error) {
	if o.Count <= 0 {
		o.Count = 10000
	}

	if o.Rate <= 0 {
		o.Rate = 30
	}

	if o.Seed == 0 {
		o.Seed = 0xD00D5EED
	}

	if o.Policy == (entry.Policy{}) {
		o.Policy = entry.DefaultPolicy()
	}

	return runJob(s, ctx, func(ctx context.Context, db *sql.DB) (DummySetup, error) {
		return ensureDummyTx(ctx, db, o)
	})
}
