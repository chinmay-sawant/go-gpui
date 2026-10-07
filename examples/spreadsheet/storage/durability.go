package storage

// Durability and failure policy
//
// Every table in the schema holds durable user data: workbooks, sheets,
// cells, revisions, and prefs. A save is acknowledged only after COMMIT,
// and no table is disposable telemetry.
//
// The shipped setting is WAL with synchronous=NORMAL, which keeps committed
// transactions across a process crash but may lose the newest commits on a
// power loss. JournalMode reports what actually activated. When WAL cannot
// activate, for example on a network share, the store keeps the rollback
// journal SQLite chose and keeps working; it never rewrites or resets data
// to force a mode.
//
// Lock contention waits inside SQLite's busy_timeout on every connection.
// This package never replays a non-idempotent write by itself; a caller
// that wants a retry repeats the whole call with a fresh base revision.
//
// Backup runs VACUUM INTO, which is safe while the database is open.
// Copying the file alone can miss committed WAL pages. Before a rename or
// delete on Windows, close the database first; never delete an active -wal
// or -shm file by hand.
