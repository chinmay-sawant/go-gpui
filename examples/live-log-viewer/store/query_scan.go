package store

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

const entryColumns = `SELECT id, session_id, source_id, source_seq, position,
	generation, ts_ns, ts_raw, ts_ok, severity, message, bytes, multiline,
	truncated, malformed, partial, received_ns FROM entries`

func scanEntry(row sourceScanner) (entry.Entry, error) {
	var (
		e                              entry.Entry
		ts, recv                       int64
		ok, multi, trunc, bad, partial int64
	)

	err := row.Scan(&e.ID, &e.Session, &e.Source, &e.Seq, &e.Position,
		&e.Generation, &ts, &e.TimeRaw, &ok, &e.Severity, &e.Message,
		&e.Bytes, &multi, &trunc, &bad, &partial, &recv)
	if err != nil {
		return entry.Entry{}, err
	}

	if ts != 0 {
		e.Time = time.Unix(0, ts)
	}

	if recv != 0 {
		e.Received = time.Unix(0, recv)
	}

	e.TimeOK = ok != 0
	e.Multiline = multi != 0
	e.Truncated = trunc != 0
	e.Malformed = bad != 0
	e.Partial = partial != 0

	return e, nil
}
