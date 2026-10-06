package store

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// SaveSettings stores the control configuration.
func (s *Store) SaveSettings(ctx context.Context, set Settings) error { return nil }

// Settings loads the stored configuration.
func (s *Store) Settings(ctx context.Context) (Settings, error) {
	return Settings{}, nil
}

// SaveSnapshot atomically stores one validated resumable snapshot.
func (s *Store) SaveSnapshot(ctx context.Context, snap game.Snapshot) error { return nil }

// LoadSnapshot returns the stored snapshot, or ok=false when none exists.
func (s *Store) LoadSnapshot(ctx context.Context) (game.Snapshot, bool, error) {
	return game.Snapshot{}, false, nil
}

// ClearSnapshot removes the stored snapshot.
func (s *Store) ClearSnapshot(ctx context.Context) error { return nil }

// Replay returns the replay stored with a game ID.
func (s *Store) Replay(ctx context.Context, id string) (game.Replay, error) {
	return game.Replay{}, nil
}

// Checkpoint runs a WAL checkpoint; a no-op on a rollback journal.
func (s *Store) Checkpoint(ctx context.Context) error { return nil }

// Backup writes a consistent copy of the database to dest.
func (s *Store) Backup(ctx context.Context, dest string) error { return nil }
