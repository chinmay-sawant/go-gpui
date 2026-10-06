// Package storage keeps the system monitor's settings, saved process views,
// recording sessions, and downsampled metric history in SQLite through
// database/sql and modernc.org/sqlite. Each example gets its own database
// under DefaultDir, and Open takes an explicit directory for tests and for a
// user override.
//
// The database has one connection. Store methods are synchronous and safe for
// concurrent use; the single connection serializes them. The example's
// recording path runs through one Recorder goroutine over a bounded queue, so
// the UI loop never waits on a write. Live process tables are never written:
// only explicit recordings store metric history.
//
// Durability differs by table. Settings, saved views, and sessions are user
// data: they commit with WAL and synchronous NORMAL. Metric history is
// disposable telemetry: a failed write is counted and reported, never
// silently retried, and cannot fail a session.
package storage
