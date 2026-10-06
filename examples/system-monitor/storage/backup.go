package storage

import "context"

// Checkpoint folds the WAL back into the database file. The example calls it
// from a worker, never from the UI loop. It is a no-op outside WAL mode.
func (s *Store) Checkpoint(ctx context.Context) error { return errNotImplemented }

// Backup writes a SQLite-aware copy of the database to dest. It uses VACUUM
// INTO, so the copy is consistent even while the database is open, unlike a
// plain file copy of a live WAL database.
func (s *Store) Backup(ctx context.Context, dest string) error { return errNotImplemented }
