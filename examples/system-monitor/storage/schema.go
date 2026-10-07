package storage

// schemaHead creates the meta, settings, and saved view tables.
const schemaHead = `
CREATE TABLE IF NOT EXISTS meta (
	key TEXT PRIMARY KEY,
	value TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS settings (
	key TEXT PRIMARY KEY CHECK (key <> ''),
	value TEXT NOT NULL,
	updated_ns INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS saved_views (
	id TEXT PRIMARY KEY CHECK (id <> ''),
	name TEXT NOT NULL CHECK (name <> ''),
	query TEXT NOT NULL,
	sort TEXT NOT NULL,
	descending INTEGER NOT NULL CHECK (descending IN (0, 1)),
	updated_ns INTEGER NOT NULL
);
`

// schemaTail creates the session and raw history tables plus the indexes every
// read uses.
const schemaTail = `
CREATE TABLE IF NOT EXISTS sessions (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	source TEXT NOT NULL,
	mode TEXT NOT NULL,
	state TEXT NOT NULL,
	note TEXT NOT NULL DEFAULT '',
	kind TEXT NOT NULL DEFAULT 'user' CHECK (kind IN ('user', 'fixture')),
	started_ns INTEGER NOT NULL,
	ended_ns INTEGER,
	sampling_ms INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS metric_history (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	session_id INTEGER NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
	metric TEXT NOT NULL,
	device TEXT NOT NULL DEFAULT '',
	ts_ns INTEGER NOT NULL,
	mono_ns INTEGER NOT NULL,
	value REAL NOT NULL,
	valid INTEGER NOT NULL CHECK (valid IN (0, 1))
);

CREATE UNIQUE INDEX IF NOT EXISTS history_unique
	ON metric_history(session_id, metric, device, ts_ns);

CREATE INDEX IF NOT EXISTS history_query
	ON metric_history(session_id, metric, device, ts_ns, id);
`
