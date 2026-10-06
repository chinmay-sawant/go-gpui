package store

// schemaV1 creates the two tables and the indexes. The partial unique
// index on destination is the storage-level ownership rule: no two active
// jobs may claim the same final path.
const schemaV1 = `
CREATE TABLE IF NOT EXISTS meta (
	key TEXT PRIMARY KEY,
	value TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS jobs (
	id TEXT PRIMARY KEY CHECK (length(id) > 0),
	url TEXT NOT NULL,
	destination TEXT NOT NULL CHECK (length(destination) > 0),
	name TEXT NOT NULL CHECK (length(name) > 0),
	state TEXT NOT NULL CHECK (state IN
		('queued','running','paused','completed','failed','cancelled')),
	done INTEGER NOT NULL DEFAULT 0 CHECK (done >= 0),
	total INTEGER NOT NULL DEFAULT -1 CHECK (total >= -1),
	expected INTEGER NOT NULL DEFAULT -1 CHECK (expected >= -1),
	etag TEXT NOT NULL DEFAULT '',
	last_modified TEXT NOT NULL DEFAULT '',
	checksum TEXT NOT NULL DEFAULT '',
	attempts INTEGER NOT NULL DEFAULT 0 CHECK (attempts >= 0),
	error TEXT NOT NULL DEFAULT '',
	created_ms INTEGER NOT NULL CHECK (created_ms >= 0),
	updated_ms INTEGER NOT NULL CHECK (updated_ms >= 0)
);

CREATE INDEX IF NOT EXISTS jobs_history
	ON jobs (state, updated_ms DESC, id DESC);

CREATE INDEX IF NOT EXISTS jobs_active
	ON jobs (state, created_ms, id);

CREATE UNIQUE INDEX IF NOT EXISTS jobs_active_destination
	ON jobs (destination)
	WHERE state IN ('queued','running','paused');
`

// migrationsV1 is the ordered migration list. Append, never edit.
var migrations = []migration{
	{version: 1, script: schemaV1},
}
