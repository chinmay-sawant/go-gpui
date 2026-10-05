package store

// schemaTail creates the teams, channels, and posts tables.
const schemaTail = `
CREATE TABLE IF NOT EXISTS teams (
	id TEXT PRIMARY KEY,
	position INTEGER NOT NULL,
	name TEXT NOT NULL,
	initials TEXT NOT NULL,
	color TEXT NOT NULL,
	expanded INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS channels (
	id TEXT PRIMARY KEY,
	team_id TEXT NOT NULL,
	position INTEGER NOT NULL,
	name TEXT NOT NULL,
	unread INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
	id TEXT PRIMARY KEY,
	channel_id TEXT NOT NULL,
	position INTEGER NOT NULL,
	author TEXT NOT NULL,
	initials TEXT NOT NULL,
	color TEXT NOT NULL,
	time TEXT NOT NULL,
	subject TEXT NOT NULL,
	text TEXT NOT NULL,
	pinned INTEGER NOT NULL,
	likes INTEGER NOT NULL,
	liked INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS replies (
	id TEXT PRIMARY KEY,
	post_id TEXT NOT NULL,
	position INTEGER NOT NULL,
	author TEXT NOT NULL,
	initials TEXT NOT NULL,
	color TEXT NOT NULL,
	time TEXT NOT NULL,
	text TEXT NOT NULL,
	own INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS reactions (
	post_id TEXT NOT NULL,
	position INTEGER NOT NULL,
	emoji TEXT NOT NULL,
	count INTEGER NOT NULL,
	mine INTEGER NOT NULL,
	PRIMARY KEY (post_id, position)
);

CREATE TABLE IF NOT EXISTS channel_files (
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

CREATE TABLE IF NOT EXISTS channel_state (
	id INTEGER PRIMARY KEY CHECK (id = 1),
	team TEXT NOT NULL,
	channel TEXT NOT NULL,
	tab TEXT NOT NULL
);
`
