package store

// schemaCore creates the bookkeeping tables: meta for schema and seed
// markers, sessions for source groups, and sources for checkpoints.
const schemaCore = `
CREATE TABLE IF NOT EXISTS meta (
	key   TEXT PRIMARY KEY,
	value TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS sessions (
	id         INTEGER PRIMARY KEY AUTOINCREMENT,
	name       TEXT NOT NULL,
	kind       TEXT NOT NULL CHECK (kind IN ('dummy','file','burst')),
	created_ns INTEGER NOT NULL,
	UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS sources (
	id          INTEGER PRIMARY KEY AUTOINCREMENT,
	session_id  INTEGER NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
	kind        TEXT NOT NULL CHECK (kind IN ('file','dummy','burst')),
	path        TEXT NOT NULL,
	label       TEXT NOT NULL,
	identity    TEXT NOT NULL DEFAULT '',
	head_hash   INTEGER NOT NULL DEFAULT 0,
	head_len    INTEGER NOT NULL DEFAULT 0,
	generation  INTEGER NOT NULL DEFAULT 1 CHECK (generation >= 1),
	position    INTEGER NOT NULL DEFAULT 0 CHECK (position >= 0),
	size        INTEGER NOT NULL DEFAULT 0,
	total       INTEGER NOT NULL DEFAULT 0,
	state       TEXT NOT NULL DEFAULT 'idle',
	seed        INTEGER NOT NULL DEFAULT 0,
	rate        INTEGER NOT NULL DEFAULT 0,
	lost        INTEGER NOT NULL DEFAULT 0,
	created_ns  INTEGER NOT NULL,
	updated_ns  INTEGER NOT NULL,
	UNIQUE (session_id, path)
);
`
