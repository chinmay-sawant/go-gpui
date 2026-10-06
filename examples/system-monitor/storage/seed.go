package storage

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// Seed writes a labeled fixture recording once per fixture version. It
// returns true when it wrote and false when the stored version already covers
// f. Reopening does not duplicate the rows, and a later version replaces the
// previous fixture session without touching settings, views, or recorded
// sessions.
func (s *Store) Seed(ctx context.Context, f domain.Fixture) (bool, error) {
	return false, errNotImplemented
}

// FixtureVersion returns the stored fixture version, or zero when none was
// seeded.
func (s *Store) FixtureVersion(ctx context.Context) (int, error) {
	return 0, errNotImplemented
}
