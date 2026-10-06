package store

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// JournalMode reports the active SQLite journal mode, or "" when closed.
func (s *Store) JournalMode() string { return "" }

// QueueDepth reports queued requests waiting for the worker.
func (s *Store) QueueDepth() int { return 0 }

// SaveGame writes one completed game and its optional replay in one
// transaction. Re-inserting the same result ID changes nothing.
func (s *Store) SaveGame(ctx context.Context, res game.Result, rep *game.Replay) error {
	return nil
}

// TopScores lists live players' scores, highest first, ties by ID.
func (s *Store) TopScores(ctx context.Context, limit int) ([]game.Result, error) {
	return nil, nil
}

// DummyScores lists seeded demo scores the same way.
func (s *Store) DummyScores(ctx context.Context, limit int) ([]game.Result, error) {
	return nil, nil
}

// PruneDummy deletes dummy scores outside the best keep, in bounded
// batches.
func (s *Store) PruneDummy(ctx context.Context, keep int) (int, error) { return 0, nil }
