package store

// schemaV1 is the first migration: meta, scores, replays, settings, and
// the single resume snapshot. Scores are durable; the dummy flag keeps
// seeded demo entries out of live rankings.
const schemaV1 = `
CREATE TABLE IF NOT EXISTS meta (
	key TEXT PRIMARY KEY,
	value TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS scores (
	id TEXT PRIMARY KEY,
	dummy INTEGER NOT NULL CHECK (dummy IN (0, 1)),
	score INTEGER NOT NULL CHECK (score BETWEEN 0 AND 999999999),
	lines INTEGER NOT NULL CHECK (lines BETWEEN 0 AND 9999),
	level INTEGER NOT NULL CHECK (level BETWEEN 1 AND 20),
	pieces INTEGER NOT NULL CHECK (pieces BETWEEN 0 AND 100000),
	duration_ms INTEGER NOT NULL CHECK (duration_ms BETWEEN 0 AND 86400000),
	seed INTEGER NOT NULL,
	ruleset TEXT NOT NULL,
	fixture_version INTEGER NOT NULL,
	created_at TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_scores_rank ON scores(dummy, score DESC, id);

CREATE TABLE IF NOT EXISTS replays (
	game_id TEXT PRIMARY KEY REFERENCES scores(id) ON DELETE CASCADE,
	seed INTEGER NOT NULL,
	ruleset TEXT NOT NULL,
	fixture_version INTEGER NOT NULL,
	fixture TEXT NOT NULL DEFAULT '',
	events TEXT NOT NULL,
	created_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS settings (
	id INTEGER PRIMARY KEY CHECK (id = 1),
	json TEXT NOT NULL,
	updated_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS snapshot (
	id INTEGER PRIMARY KEY CHECK (id = 1),
	game_id TEXT NOT NULL,
	payload TEXT NOT NULL,
	updated_at TEXT NOT NULL
);
`
