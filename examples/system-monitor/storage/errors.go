package storage

import (
	"errors"
	"time"
)

// SchemaVersion is the schema this build writes. A database with a higher
// version belongs to a newer build and Open refuses it without modifying it.
const SchemaVersion = 1

// Retention and bounding policy for metric history. Raw rows live 24 hours,
// then Retain folds them into five-minute aggregates; aggregates live 30
// days. The caps bound the tables even when the clock jumps or a recorder
// runs for weeks.
const (
	RawRetention       = 24 * time.Hour
	AggregateRetention = 30 * 24 * time.Hour
	AggregateBucket    = 5 * time.Minute
	MaxRawRows         = 1_000_000
	MaxAggregateRows   = 500_000
	BatchRows          = 5000
)

// ErrSchemaNewer reports a database written by a newer build. Open leaves the
// file untouched.
var ErrSchemaNewer = errors.New("system-monitor: database schema is newer than this build")

// ErrClosed reports work on a closed store.
var ErrClosed = errors.New("system-monitor: store is closed")

// errNotImplemented marks the seam stubs. It disappears as the package fills
// in.
var errNotImplemented = errors.New("system-monitor: not implemented")
