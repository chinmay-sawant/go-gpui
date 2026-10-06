package storage

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// Row is one metric value written to history. It matches the collector's
// record shape, so a sample's Records convert directly.
type Row = domain.Record

// Query filters a history read. SessionID is required. Since and Until bound
// At; AfterID continues a keyset read; Limit caps the rows returned. A zero
// Limit takes a default of 1000 rows.
type Query struct {
	SessionID int64
	Metric    domain.Metric
	Device    string
	Since     time.Time
	Until     time.Time
	AfterID   int64
	Limit     int
}

// Point is one stored value with its stable row ID. IDs increase within a
// session, so a reader can page with AfterID while retention removes old rows.
type Point struct {
	ID     int64
	Metric domain.Metric
	Device string
	At     time.Time
	Mono   time.Duration
	Value  float64
	Valid  bool
}

// Aggregate is one downsample bucket: the minimum, maximum, and average of
// the raw values in a bucket, plus how many raw rows it replaces.
type Aggregate struct {
	Metric domain.Metric
	Device string
	Bucket time.Time
	Min    float64
	Max    float64
	Avg    float64
	Count  int64
}

// RetainResult reports what one retention pass moved or removed.
type RetainResult struct {
	Folded     int64
	Deleted    int64
	Aggregates int64
	Trimmed    int64
}
