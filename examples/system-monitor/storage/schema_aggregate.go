package storage

// schemaAggregate creates the downsample table.
const schemaAggregate = `
CREATE TABLE IF NOT EXISTS metric_aggregate (
	session_id INTEGER NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
	metric TEXT NOT NULL,
	device TEXT NOT NULL DEFAULT '',
	bucket_ns INTEGER NOT NULL,
	min_value REAL NOT NULL,
	max_value REAL NOT NULL,
	avg_value REAL NOT NULL,
	count INTEGER NOT NULL,
	PRIMARY KEY (session_id, metric, device, bucket_ns)
);

CREATE INDEX IF NOT EXISTS aggregate_age ON metric_aggregate(bucket_ns);
`

// schemaSQL is the version 1 schema.
const schemaSQL = schemaHead + schemaTail + schemaAggregate
