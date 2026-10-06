package storage

// schemaTables creates the durable tables. Cells hold value columns for
// numbers, text, and formula source; calculated caches are never stored, so
// a reopen can never trust a stale result.
const schemaTables = `
CREATE TABLE IF NOT EXISTS meta (
	key TEXT PRIMARY KEY,
	value TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS workbooks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL CHECK (length(name) > 0),
	rev INTEGER NOT NULL DEFAULT 0 CHECK (rev >= 0),
	created_at TEXT NOT NULL,
	updated_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS sheets (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	workbook_id INTEGER NOT NULL REFERENCES workbooks(id) ON DELETE CASCADE,
	name TEXT NOT NULL CHECK (length(name) > 0),
	position INTEGER NOT NULL CHECK (position >= 0),
	UNIQUE (workbook_id, name)
);

CREATE TABLE IF NOT EXISTS cells (
	sheet_id INTEGER NOT NULL REFERENCES sheets(id) ON DELETE CASCADE,
	row_ix INTEGER NOT NULL CHECK (row_ix >= 0 AND row_ix < 1048576),
	col_ix INTEGER NOT NULL CHECK (col_ix >= 0 AND col_ix < 18278),
	kind INTEGER NOT NULL CHECK (kind IN (1, 2, 3)),
	number REAL,
	text TEXT,
	source TEXT,
	PRIMARY KEY (sheet_id, row_ix, col_ix)
) WITHOUT ROWID;
`

// schemaRest adds the revision log, indexes, and preferences.
const schemaRest = `
CREATE INDEX IF NOT EXISTS sheets_book ON sheets (workbook_id, position);

CREATE TABLE IF NOT EXISTS revisions (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	workbook_id INTEGER NOT NULL REFERENCES workbooks(id) ON DELETE CASCADE,
	rev INTEGER NOT NULL CHECK (rev >= 0),
	origin TEXT NOT NULL,
	changed INTEGER NOT NULL CHECK (changed >= 0),
	created_at TEXT NOT NULL,
	UNIQUE (workbook_id, rev)
);

CREATE INDEX IF NOT EXISTS revisions_book ON revisions (workbook_id, rev);

CREATE TABLE IF NOT EXISTS prefs (
	key TEXT PRIMARY KEY,
	value TEXT NOT NULL
);
`

// schemaV1 is the initial schema.
const schemaV1 = schemaTables + schemaRest
