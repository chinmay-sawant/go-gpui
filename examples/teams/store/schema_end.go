package store

// schemaEnd creates the calendar, calls, and files tables.
const schemaEnd = `
CREATE TABLE IF NOT EXISTS calendar_events (
	id TEXT PRIMARY KEY,
	week INTEGER NOT NULL,
	position INTEGER NOT NULL,
	title TEXT NOT NULL,
	time TEXT NOT NULL,
	dur TEXT NOT NULL,
	color TEXT NOT NULL,
	location TEXT NOT NULL,
	day INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS calendar_state (
	id INTEGER PRIMARY KEY CHECK (id = 1),
	week INTEGER NOT NULL,
	selected TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS calls (
	id TEXT PRIMARY KEY,
	position INTEGER NOT NULL,
	name TEXT NOT NULL,
	initials TEXT NOT NULL,
	color TEXT NOT NULL,
	kind TEXT NOT NULL,
	time TEXT NOT NULL,
	duration TEXT NOT NULL,
	missed INTEGER NOT NULL,
	incoming INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS voicemails (
	id TEXT PRIMARY KEY,
	position INTEGER NOT NULL,
	name TEXT NOT NULL,
	initials TEXT NOT NULL,
	color TEXT NOT NULL,
	time TEXT NOT NULL,
	duration TEXT NOT NULL,
	transcript TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS files (
	id TEXT PRIMARY KEY,
	position INTEGER NOT NULL,
	name TEXT NOT NULL,
	badge TEXT NOT NULL,
	badge_text TEXT NOT NULL,
	kind TEXT NOT NULL,
	modified TEXT NOT NULL,
	modified_by TEXT NOT NULL,
	size TEXT NOT NULL,
	location TEXT NOT NULL,
	team INTEGER NOT NULL,
	shared INTEGER NOT NULL,
	starred INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS files_state (
	id INTEGER PRIMARY KEY CHECK (id = 1),
	filter TEXT NOT NULL,
	active TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS calls_state (
	id INTEGER PRIMARY KEY CHECK (id = 1),
	tab TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS activity_state (
	id INTEGER PRIMARY KEY CHECK (id = 1),
	filter TEXT NOT NULL,
	active TEXT NOT NULL
);
`

// schema is the whole database layout.
const schema = schemaHead + schemaTail + schemaEnd
