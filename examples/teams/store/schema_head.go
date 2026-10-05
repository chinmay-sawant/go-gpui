package store

// schemaHead creates the meta and profile tables and the menu tables that
// hold the list state.
const schemaHead = `
CREATE TABLE IF NOT EXISTS meta (
	key TEXT PRIMARY KEY,
	value TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS profile (
	id INTEGER PRIMARY KEY CHECK (id = 1),
	presence TEXT NOT NULL,
	dark INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS activity_items (
	id TEXT PRIMARY KEY,
	position INTEGER NOT NULL,
	actor TEXT NOT NULL,
	initials TEXT NOT NULL,
	color TEXT NOT NULL,
	kind TEXT NOT NULL,
	place TEXT NOT NULL,
	text TEXT NOT NULL,
	preview TEXT NOT NULL,
	time TEXT NOT NULL,
	unread INTEGER NOT NULL,
	replies INTEGER NOT NULL,
	replied INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS chats (
	id TEXT PRIMARY KEY,
	position INTEGER NOT NULL,
	name TEXT NOT NULL,
	initials TEXT NOT NULL,
	color TEXT NOT NULL,
	preview TEXT NOT NULL,
	time TEXT NOT NULL,
	unread INTEGER NOT NULL,
	pinned INTEGER NOT NULL,
	muted INTEGER NOT NULL,
	is_group INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS chat_messages (
	id TEXT PRIMARY KEY,
	chat_id TEXT NOT NULL,
	position INTEGER NOT NULL,
	author TEXT NOT NULL,
	initials TEXT NOT NULL,
	color TEXT NOT NULL,
	time TEXT NOT NULL,
	text TEXT NOT NULL,
	own INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS chat_state (
	id INTEGER PRIMARY KEY CHECK (id = 1),
	active TEXT NOT NULL,
	filter TEXT NOT NULL,
	query TEXT NOT NULL
);
`
