package store

// schemaData creates the ordered entries table and the indexes the viewer
// pages by. The unique (source_id, generation, position) index is what makes
// a replayed or retried batch idempotent.
const schemaData = `
CREATE TABLE IF NOT EXISTS entries (
	id          INTEGER PRIMARY KEY AUTOINCREMENT,
	session_id  INTEGER NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
	source_id   INTEGER NOT NULL REFERENCES sources(id) ON DELETE CASCADE,
	source_seq  INTEGER NOT NULL,
	position    INTEGER NOT NULL,
	generation  INTEGER NOT NULL,
	ts_ns       INTEGER NOT NULL DEFAULT 0,
	ts_raw      TEXT NOT NULL DEFAULT '',
	ts_ok       INTEGER NOT NULL DEFAULT 0,
	severity    INTEGER NOT NULL DEFAULT 0,
	message     TEXT NOT NULL,
	bytes       INTEGER NOT NULL DEFAULT 0,
	multiline   INTEGER NOT NULL DEFAULT 0,
	truncated   INTEGER NOT NULL DEFAULT 0,
	malformed   INTEGER NOT NULL DEFAULT 0,
	partial     INTEGER NOT NULL DEFAULT 0,
	received_ns INTEGER NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS entries_ident
	ON entries(source_id, generation, position);

CREATE UNIQUE INDEX IF NOT EXISTS entries_seq
	ON entries(source_id, source_seq);

CREATE INDEX IF NOT EXISTS entries_session ON entries(session_id, id);

CREATE INDEX IF NOT EXISTS entries_source ON entries(source_id, id);

CREATE INDEX IF NOT EXISTS entries_sev ON entries(session_id, severity, id);

CREATE INDEX IF NOT EXISTS entries_recv ON entries(received_ns);
`
