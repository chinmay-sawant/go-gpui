package store

// Storage policy
//
// Files. The database is jobs.sqlite beside instance.lock in the data
// directory. While WAL is active SQLite adds jobs.sqlite-wal and
// jobs.sqlite-shm; never copy, rename, or delete those side files by hand.
//
// Journal. Open prefers WAL and falls back to the rollback journal when
// the storage refuses it, for example on some network shares. JournalMode
// reports "wal", "delete", or "memory". WAL still allows one writer, and
// every statement here runs on the one worker and one connection.
//
// Durability per table. jobs rows are durable: a committed transition is
// on disk, and a failed write returns an error instead of a silent
// success. meta rows are durable. schema_migrations rows are durable.
// Progress checkpoints are durable writes too, but a checkpoint only
// records observed bytes; after a restart the partial file is re-measured,
// because the row may be ahead of the file.
//
// Checkpoints. The scheduler calls Checkpoint off the UI loop at a bounded
// rate. Backup runs PRAGMA wal_checkpoint(PASSIVE) on the worker and then
// VACUUM INTO, so a backup never copies a live WAL by hand.
//
// Retention. Cleanup deletes at most one batch per call, newest kept.
// Vacuum and other large maintenance jobs never run during an interaction.
//
// Instance lock. Lock writes instance.lock with O_EXCL and heartbeats its
// mtime every five seconds. A file with no heartbeat for twenty seconds is
// a crash leftover and is reclaimed; a fresh one returns ErrLocked.
//
// Temporary mode. OpenMemory keeps everything in RAM, takes no lock, and
// loses the data on Close. Use it when the data directory is unusable and
// label the screen temporary.
