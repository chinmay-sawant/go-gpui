package entry

import "time"

// RawRecord is one record as a reader produced it. Offset is a byte offset
// in a file or a source-defined position for a stream. Bytes counts every
// consumed byte, including the line terminator and any bytes skipped from a
// truncated record. Data is the record body without its line terminator.
type RawRecord struct {
	Path       string
	Generation int64
	Offset     int64
	Bytes      int
	Data       []byte
	Partial    bool
	Truncated  bool
	Skipped    int
}

// End is the position just past the record.
func (r RawRecord) End() int64 { return r.Offset + int64(r.Bytes) }

// Entry is one stored log record. ID orders entries across the database.
// Position is the reader position (byte offset or record number) the entry
// started at; Bytes counts the raw bytes it covers.
type Entry struct {
	ID         EntryID
	Session    SessionID
	Source     SourceID
	Seq        int64
	Position   int64
	Generation int64
	Time       time.Time
	TimeRaw    string
	TimeOK     bool
	Severity   Severity
	Message    string
	Bytes      int
	Multiline  bool
	Truncated  bool
	Malformed  bool
	Partial    bool
	Received   time.Time
}
